package main

import (
    "embed"

    "github.com/section14/evenflow/internal/api"
)

//go:embed templates/*
var templates embed.FS

func main() {
    api.Serve(templates)
}
