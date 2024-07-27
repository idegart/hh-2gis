package main

import (
	"2GIS/internal/apps/server"
	"time"
)

func init() {
	time.Local = time.UTC
}

func main() {
	server.Run()
}
