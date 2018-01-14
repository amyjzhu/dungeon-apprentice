package main

import "path/filepath"

func main() {

	absPath, _ := filepath.Abs("monsters.json")
	loadStringFromFile(absPath)

	// startServer()
}
