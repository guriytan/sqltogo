package sqltogo

import (
	"github.com/liangyaopei/sqltogo/internal"
	"io"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	open, err := os.Open("./input.sql")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = open.Close() }()
	bytes, err := io.ReadAll(open)
	if err != nil {
		t.Fatal(err)
	}
	toGo, err := Parse(internal.BytesToString(bytes), "main", false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(toGo)
}
