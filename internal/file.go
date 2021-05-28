package internal

import (
	"fmt"
	"os"
)

func GenGoFile(lines, tableName string) error {
	file, err := os.Create(fmt.Sprintf("./%s.go", tableName))
	if err != nil {
		return err
	}
	defer func() { _ = file.Close() }()
	_, err = file.WriteString(lines)
	if err != nil {
		return err
	}
	return nil
}
