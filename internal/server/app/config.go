package app

import "time"

type ServerConfig struct {
	Address            string
	StoreInterval      time.Duration
	FileStoragePath    string
	NeedRestore        bool
	DatabaseConnection string
	SignKey            string
}
