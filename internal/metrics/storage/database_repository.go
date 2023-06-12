package storage

import (
	"database/sql"

	retry "github.com/mrLandyrev/golang-metrics/internal"
	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type DatabaseMetricsRepository struct {
	db      *sql.DB
	factory metrics.MetricsFactory
}

func (storage *DatabaseMetricsRepository) GetAll() ([]metrics.Metric, error) {
	var rows *sql.Rows

	err := retry.HandleFunc(func() error {
		var err error

		rows, err = storage.db.Query(`
			SELECT name, kind, value
			FROM metrics
		`)

		if err != nil {
			return err
		}

		if rows.Err() != nil {
			return rows.Err()
		}

		return nil
	}, 4, nil)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	data := make([]metrics.Metric, 0)
	var name string
	var kind string
	var value string

	for rows.Next() {
		rows.Scan(&name, &kind, &value)

		instance, err := storage.factory.GetInstance(kind, name)

		if err != nil {
			continue
		}

		err = instance.AddValue(value)

		if err != nil {
			continue
		}

		data = append(data, instance)
	}

	return data, nil
}

func (storage *DatabaseMetricsRepository) GetByKindAndName(kind string, name string) (metrics.Metric, error) {
	var rows *sql.Rows

	err := retry.HandleFunc(func() error {
		var err error
		rows, err = storage.db.Query(`
			SELECT value
			FROM metrics
			WHERE name = $1
			AND kind = $2
			LIMIT 1
		`,
			name,
			kind,
		)

		if err != nil {
			return err
		}

		if rows.Err() != nil {
			return rows.Err()
		}

		return nil
	}, 4, nil)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	if !rows.Next() {
		return nil, nil
	}

	var value string
	err = rows.Scan(&value)

	if err != nil {
		return nil, err
	}

	instance, err := storage.factory.GetInstance(kind, name)

	if err != nil {
		return nil, err
	}

	err = instance.AddValue(value)

	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (storage *DatabaseMetricsRepository) Persist(item metrics.Metric) error {
	_, err := storage.db.Exec(`
		INSERT INTO metrics (name, kind, value) 
		VALUES($1, $2, $3)
		ON CONFLICT (name, kind) DO UPDATE
		SET value = EXCLUDED.value;
	`, item.Name(), item.Kind(), item.Value())

	if err != nil {
		return err
	}

	return nil
}

func (storage *DatabaseMetricsRepository) Clear() error {
	return nil
}

func NewDatabaseMetricsRepository(db *sql.DB) (*DatabaseMetricsRepository, error) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS metrics (
			name VARCHAR(255),
			kind VARCHAR(255),
			value VARCHAR(255),
			UNIQUE(name, kind)
		)
	`)

	if err != nil {
		return nil, err
	}

	return &DatabaseMetricsRepository{db: db}, nil
}
