package main

import "bytes"
import "encoding/csv"
import "errors"
import "fmt"
import "io"
import "path/filepath"
import "strings"

import "github.com/codegangsta/cli"

func ListWordlistNames(c *cli.Context) {
	fmt.Println("The following word lists are available:")
	for _, name := range AssetNames() {
		cleanName, err := sanitizeWordlistName(name)
		if err != nil {
			continue
		}

		fmt.Printf("  %s\n", cleanName)
	}
}

func getWordlistByName(name string) (words []string, err error) {
	raw, err := Asset("data/" + name + ".asc")
	if err == nil {
		words, err = parseWordlist(bytes.NewReader(raw), 2)
		return
	}

	raw, err = Asset("data/" + name + ".txt")
	if err == nil {
		words, err = parseWordlist(bytes.NewReader(raw), 1)
		return
	}

	err = errors.New("Wordlist '" + name + "' was not found")
	return
}

func parseWordlist(r io.Reader, cols int) (words []string, err error) {
	words = make([]string, 0, 1000)

	// Tab-separated tuples
	valueReader := csv.NewReader(r)
	valueReader.Comma = '\t'
	valueReader.Comment = '-'
	valueReader.FieldsPerRecord = cols
	valueReader.LazyQuotes = true

	record, err := valueReader.Read()
	for ; err == nil; record, err = valueReader.Read() {
		words = append(words, record[cols - 1]) // discard anything but the last column
	}

	// For the purposes of parseWordlist, EOF is a success
	if err == io.EOF {
		err = nil
	}

	return
}

func sanitizeWordlistName(name string) (cleanName string, err error) {
	ext := filepath.Ext(name)
	if ext != ".txt" && ext != ".asc" {
		err = errors.New("Wordlist name '" + name + "' must end in '.txt' or '.asc'")
		return
	}

	cleanName = strings.TrimSuffix(filepath.Base(name), ext)
	return
}
