package main

import (
	"blockchain/blockchain/internal/handler"
	"blockchain/blockchain/internal/repository"
	"blockchain/blockchain/internal/service"
	"path/filepath"
	"runtime"
)

func main() {
	repo := repository.NewBlockchainRepository()
	svc := service.NewBlockchainService(repo)
	h := handler.NewBlockchainHandler(svc)

	_, file, _, _ := runtime.Caller(0)
	webDir := filepath.Join(filepath.Dir(file), "..", "..", "web")

	r := handler.NewRouter(h, webDir)
	r.Run(":8083")
}
