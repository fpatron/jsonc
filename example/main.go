package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fpatron/jsonc"
)

type Metadata struct {
	Version      string   `json:"version"`
	Author       string   `json:"author"`
	Contributors []string `json:"contributors"`
}

type Address struct {
	City    string `json:"city"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type User struct {
	ID      int      `json:"id"`
	Name    string   `json:"name"`
	Address Address  `json:"address"`
	Roles   []string `json:"roles"`
}

type Config struct {
	Debug     bool     `json:"debug"`
	Theme     string   `json:"theme"`
	Languages []string `json:"languages"`
}

type ExampleData struct {
	Metadata Metadata `json:"metadata"`
	Users    []User   `json:"users"`
	Config   Config   `json:"config"`
	Files    []string `json:"files"`
}

func main() {
	// Read example.json
	data, err := os.ReadFile("example.json")
	if err != nil {
		log.Fatalf("Failed to read example.json: %v", err)
	}

	var example ExampleData
	err = jsonc.Unmarshal(data, &example)
	if err != nil {
		log.Fatalf("Failed to parse example.json: %v", err)
	}

	fmt.Printf("Parsed data: %+v\n", example)
}
