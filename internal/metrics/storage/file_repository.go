package storage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/mrLandyrev/golang-metrics/internal/metrics"
)

type StorageRecord struct {
	Name  string `json:"name"`
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

type FileMetricsRepository struct {
	MemoryMetricsRepository
	filename string
	factory  metrics.MetricsFactory
	isSync   bool
}

func (storage *FileMetricsRepository) Flush() error {
	file, err := os.Create(storage.filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	for _, m := range storage.data {
		record := &StorageRecord{
			Name:  m.Name(),
			Kind:  m.Kind(),
			Value: m.Value(),
		}

		b, err := json.Marshal(record)

		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(file, string(b))

		if err != nil {
			return err
		}
	}

	return nil
}

func (storage *FileMetricsRepository) Restore() error {
	file, err := os.Open(storage.filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var record StorageRecord

	for scanner.Scan() {
		err = json.Unmarshal([]byte(scanner.Text()), &record)
		if err != nil {
			continue
		}

		instance, err := storage.factory.GetInstance(record.Kind, record.Name)

		if err != nil {
			continue
		}

		err = instance.AddValue(record.Value)

		if err != nil {
			continue
		}

		storage.data = append(storage.data, instance)
	}

	return nil
}

func (storage *FileMetricsRepository) Persist(item metrics.Metric) error {
	err := storage.MemoryMetricsRepository.Persist(item)

	if err != nil {
		return err
	}

	if storage.isSync {
		return storage.Flush()
	}

	return nil
}

func NewFileMetricsRepository(filename string, storeInterval time.Duration, NeedRestore bool) (*FileMetricsRepository, error) {
	repo := &FileMetricsRepository{
		MemoryMetricsRepository: *NewMemoryMetricsRepository(),
		filename:                filename,
		isSync:                  storeInterval == 0,
	}

	fmt.Println(storeInterval)

	if NeedRestore {
		repo.Restore()
	}

	if !repo.isSync {
		go func() {
			for {
				repo.Flush()
				time.Sleep(storeInterval)
			}
		}()
	}

	return repo, nil
}
