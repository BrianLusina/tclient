package torrentfile

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	torrent, err := Open("testdata/archlinux-2019.12.01-x86_64.iso.torrent")
	require.Nil(t, err)

	goldenPath := "testdata/archlinux-2019.12.01-x86_64.iso.torrent.golden.json"

	if *update {
		serialized, err := json.MarshalIndent(torrent, "", "  ")
		require.Nil(t, err)
		_ = ioutil.WriteFile(goldenPath, serialized, 0644)
	}

	expected := TorrentFile{}
	golden, err := ioutil.ReadFile(goldenPath)
	require.Nil(t, err)

	err = json.Unmarshal(golden, &expected)
	require.Nil(t, err)

	assert.Equal(t, expected, torrent)
}
