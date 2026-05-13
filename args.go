package main

import (
	"encoding/json"
	"flag"
	"os"
	"path/filepath"
)

type Options struct {
	file string
}

func parseArgs() (Options, error) {
	// For now we only have this option, but there may be more in the future.
	// TODO: Separate file initialisation from argument parsing.
	file := flag.String("f", "", "file to read-write notes")

	flag.Parse()

	if *file == "" {
		*file = defaultSaveLocation()
	}

	err := ensureFileExists(*file)
	if err != nil {
		return Options{}, err
	}

	return Options{file: *file}, nil
}

func defaultSaveLocation() string {
	file := "terminoter.json"
	app := "terminoter"

	dir := os.Getenv("XDG_DATA_HOME")
	if dir == "" {
		home := os.Getenv("HOME")
		if home == "" {
			return filepath.Join(".", file)
		}
		dir = filepath.Join(home, ".local", "share")
	}
	return filepath.Join(dir, app, file)
}

func ensureFileExists(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return err
	}
	// If file doesn't exist, create.
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0o600)
	if err != nil {
		return err
	}

	err = f.Close()
	if err != nil {
		return err
	}

	info, err := os.Stat(path)
	if err != nil {
		return err
	}

	// If the file is empty, initialise it.
	if info.Size() == 0 {
		data := Data{}

		b, err := json.Marshal(data)
		if err != nil {
			return err
		}

		err = os.WriteFile(path, b, 0644)
		if err != nil {
			return err
		}
	}

	return nil
}
