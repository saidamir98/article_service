package handlers

import (
	"uacademy/article/config"
	"uacademy/article/storage"
)

// Handler ...
type handler struct {
	Stg storage.StorageI
	Cfg config.Config
}

// NewHandler ...
func NewHandler(stg storage.StorageI, cfg config.Config) handler {
	return handler{
		Stg: stg,
		Cfg: cfg,
	}
}
