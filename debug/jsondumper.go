package debugTools

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

func writeToJSON(filePath string, data any) {
	gameStateJSON, _ := json.Marshal(data)
	file, _ := os.Create(filePath)
	defer file.Close()
	file.Write(gameStateJSON)
}

type mazeData struct {
	MazeData [][][]int
}

func openJSON(filename string) mazeData {
	file, error := os.Open(filename)
	if error != nil {
		log.Fatalf("Failed to open file: %v", error)
	}
	defer file.Close()

	byteval, error := io.ReadAll(file)
	if error != nil {
		log.Fatalf("Failed to read file: %v", error)
	}
	var result mazeData
	error = json.Unmarshal(byteval, &result)
	if error != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", error)
	}
	fmt.Println("Got JSON data")
	return result
}
