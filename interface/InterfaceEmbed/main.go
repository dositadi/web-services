package main

import "fmt"

type MyFile struct {
	File []byte
}

type Reader interface {
	Read(file []byte) (int, error)
}

type Writer interface {
	Write(source []byte) (int, error)
}

type ReaderWriter interface {
	Reader
	Writer
}

func (destination *MyFile) Write(source []byte) (int, error) {
	fmt.Println("Writing into destination file...")
	destination.File = append(destination.File, []byte(" ")...)
	destination.File = append(destination.File, source...)
	return len(source), nil
}

func (source MyFile) Read(destination []byte) (int, error) {
	fmt.Println("Reading from source file...")
	copy(destination, source.File)
	return len(destination), nil
}

func main() {
	file := &MyFile{
		File: []byte("Go is awesome!"),
	}

	var rw ReaderWriter

	rw = file
	readBuffer := make([]byte, len(file.File))
	rw.Read(readBuffer)
	fmt.Println(string(readBuffer))

	rw.Write([]byte("I can't stop loving this language"))
	fmt.Println(string(file.File))
}
