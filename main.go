package main

import (
	"bytes"
	"encoding/json"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	args := os.Args
	if len(args) == 1 {
		panic("insert filename")
	}

	filename := args[1]
	b, err := os.ReadFile(filename)
	check(err)
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	check(err)
	file, err := os.Create(filename)
	defer file.Close()
	_, err = out.WriteTo(file)
	check(err)
}
