package db

import (
	"fmt"
	"os"
)

const DBPath = "/home/ositadinma/src/web services/basichttprequest/server-client2/db.log"

func ReadFromDB() ([]byte, error) {
	file, err := os.OpenFile(DBPath, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("Database Read Err: %w", err)
	}

	fileInfo, err1 := file.Stat()
	if err1 != nil {
		return nil, fmt.Errorf("Database Stat Error: %w", err1)
	}

	dbBuffer := make([]byte, fileInfo.Size())

	_, err3 := file.Read(dbBuffer)
	if err3 != nil {
		return nil, fmt.Errorf("Database Read Err: %w", err3)
	}

	return dbBuffer, nil
}

func WriteToDB(data []byte) error {
	file, err := os.OpenFile(DBPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("Database Write Error: %w", err)
	}

	_, err2 := file.Write(data)
	if err2 != nil {
		return fmt.Errorf("Database Write Error: %w", err2)
	}
	return nil
}
