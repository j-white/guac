//
// Copyright 2023 The GUAC Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var header = "Code generated by ent, DO NOT EDIT."

// Deletes all Ent generated code from a given directory.
func main() {
	if err := cmd(); err != nil {
		log.Fatalln(err)
	}
}

func cmd() error {
	path := os.Args[1]

	if path == "" {
		return errors.New("no path provided")
	}

	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	log.Printf("Starting at %q", path)
	return enterDirectory(path)
}

func enterDirectory(path string) error {
	log.Printf("Entering %s", path)

	dir, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range dir {
		if err := processDirectory(path, entry); err != nil {
			return err
		}
	}

	return nil
}

func processDirectory(path string, entry os.DirEntry) error {
	if entry.IsDir() {
		return enterDirectory(filepath.Join(path, entry.Name()))
	}

	if filepath.Ext(entry.Name()) != ".go" {
		log.Printf("Skip non-go file %q", entry.Name())
		return nil
	}

	file, err := os.Open(filepath.Join(path, entry.Name()))
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)
	for i := 0; i < 5; i++ {
		line, _, err := reader.ReadLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			return err
		}

		// Only delete files containing the "Code generated by ent, DO NOT EDIT." header.
		if strings.Contains(string(line), header) {
			log.Printf("Deleting %q", entry.Name())
			if err := os.Remove(filepath.Join(path, entry.Name())); err != nil {
				return err
			}
			break
		}
	}

	return nil
}
