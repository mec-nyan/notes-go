package main

import "flag"

type Options struct {
	file string
}

func parseArgs() Options {
	file := flag.String("f", "", "file to read-write notes")

	flag.Parse()

	return Options{
		file: *file,
	}
}
