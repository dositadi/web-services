package main

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

// Struct for logging message to the console
type ConsoleLog struct{}

func (cl ConsoleLog) Log(message string) {
	fmt.Println("CONSOLE: ", message)
}

// Struct for logging message to a File
type FileLog struct {
	FilePath string
	File     *os.File
}

// FileLog constructor or generator
func NewFileLog(filePath string) (*FileLog, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("Error opening specified file: %v,%E", filePath, err)
	}
	return &FileLog{FilePath: filePath, File: file}, nil
}

// Function to log the message into the file after creating the file
func (f *FileLog) Log(message string) {
	if f.File == nil {
		log.Fatal("Error: FileLog file is nil")
		return
	}
	_, err := f.File.WriteString("File: " + message + "\n")
	if err != nil {
		fmt.Printf("Error writing into file: %v, %E, \n", f.FilePath, err)
	}
}

// Close the file after logging the message and also clean up resources
func (f FileLog) Close() {
	if f.File != nil {
		f.File.Close()
	}
}

// ApplicationServices uses Logger interface to log events
type ApplicationServices struct {
	logger Logger // This depends on the logger interface
}

func (as ApplicationServices) ProcessData(data string) {
	as.logger.Log("Processing data: " + data)
	// After processing
	as.logger.Log("Processing complete for data: " + data)
}

func main() {
	consoleLog := ConsoleLog{}
	appService1 := ApplicationServices{logger: consoleLog}
	appService1.ProcessData("GoodMorning")

	fmt.Println()
	fileLog, err := NewFileLog("DecoplingInterface/file.log")
	if err != nil {
		log.Fatal("Error: Could not create file!")
		return
	}

	defer fileLog.Close()

	appService2 := ApplicationServices{logger: fileLog}
	appService2.ProcessData("Order 123")
}
