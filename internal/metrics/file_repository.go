package metrics

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type StorageRecord struct {
	Name  string `json:"name"`
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

type FileMetricsRepository struct {
	MemoryMetricsRepository
	filename   string
	factory    MetricsFactory
	hasChanges bool
	isSync     bool
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

func (storage *FileMetricsRepository) Persist(item Metric) error {
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
		hasChanges:              false,
		isSync:                  storeInterval == 0,
	}

	if NeedRestore {
		repo.Restore()
	}

	if !repo.isSync {
		go func() {
			for {
				if repo.hasChanges {
					repo.Flush()
				}
				time.Sleep(storeInterval)
			}
		}()

		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		go func() {
			<-gracefulStop
			if repo.hasChanges {
				repo.Flush()
			}
			os.Exit(0)
		}()
	}

	return repo, nil
}
