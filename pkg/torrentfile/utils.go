package torrentfile

import (
	"log"
	"os"

	"github.com/jackpal/bencode-go"
)

// Open parses a torrent file
func Open(path string) (TorrentFile, error) {
	file, err := os.Open(path)
	if err != nil {
		return TorrentFile{}, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatalf("An error occurred closing the file")
			return
		}
	}(file)

	bto := bencodeTorrent{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		return TorrentFile{}, err
	}

	return bto.toTorrentFile()
}
