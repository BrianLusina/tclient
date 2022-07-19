package main

import (
	"flag"
	"log"

	"github.com/brianlusina/tclient/pkg/torrentfile"
)

func main() {
	var torrentFilePath string
	var ouputFilePath string
	flag.StringVar(&torrentFilePath, "input", "", "Input file")
	flag.StringVar(&ouputFilePath, "output", "./", "Where to download the file")
	flag.Parse()

	if len(torrentFilePath) == 0 {
		log.Fatalf("Please provide a valid path")
	}

	tf, err := torrentfile.Open(torrentFilePath)
	if err != nil {
		log.Fatalf("Failed to open file %s, err: %s", torrentFilePath, err)
	}

	err = tf.DownloadToFile(ouputFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
