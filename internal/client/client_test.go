package client

import (
	"net"
	"testing"

	"github.com/brianlusina/tclient/pkg/message"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createClientAndServer(t *testing.T) (clientConn, serverConn net.Conn) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.Nil(t, err)

	// net.Dial does not block, so we need this signalling channel to make sure
	// we don't return before serverConn is ready
	done := make(chan struct{})
	go func() {
		defer ln.Close()
		serverConn, err = ln.Accept()
		require.Nil(t, err)
		done <- struct{}{}
	}()

	clientConn, err = net.Dial("tcp", ln.Addr().String())
	<-done

	return clientConn, serverConn
}

func TestRead(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	client := Client{Conn: clientConn}

	msgBytes := []byte{
		0x00, 0x00, 0x00, 0x05,
		4,
		0x00, 0x00, 0x05, 0x3c,
	}
	expected := &message.Message{
		ID:      message.MsgHave,
		Payload: []byte{0x00, 0x00, 0x05, 0x3c},
	}
	_, err := serverConn.Write(msgBytes)
	require.Nil(t, err)

	msg, _ := client.Read()
	assert.Equal(t, expected, msg)
}

func TestSendRequest(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	client := Client{Conn: clientConn}
	err := client.SendRequest(1, 2, 3)
	assert.Nil(t, err)
	expected := []byte{
		0x00, 0x00, 0x00, 0x0d,
		6,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x03,
	}
	buf := make([]byte, len(expected))
	_, err = serverConn.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, expected, buf)
}

func TestSendInterested(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	client := Client{Conn: clientConn}
	err := client.SendInterested()
	assert.Nil(t, err)
	expected := []byte{
		0x00, 0x00, 0x00, 0x01,
		2,
	}
	buf := make([]byte, len(expected))
	_, err = serverConn.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, expected, buf)
}

func TestSendNotInterested(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	client := Client{Conn: clientConn}
	err := client.SendNotInterested()
	assert.Nil(t, err)
	expected := []byte{
		0x00, 0x00, 0x00, 0x01,
		3,
	}
	buf := make([]byte, len(expected))
	_, err = serverConn.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, expected, buf)
}

func TestSendUnchoke(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	client := Client{Conn: clientConn}
	err := client.SendUnchoke()
	assert.Nil(t, err)
	expected := []byte{
		0x00, 0x00, 0x00, 0x01,
		1,
	}
	buf := make([]byte, len(expected))
	_, err = serverConn.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, expected, buf)
}

func TestSendHave(t *testing.T) {
	clientConn, serverConn := createClientAndServer(t)
	client := Client{Conn: clientConn}
	err := client.SendHave(1340)
	assert.Nil(t, err)
	expected := []byte{
		0x00, 0x00, 0x00, 0x05,
		4,
		0x00, 0x00, 0x05, 0x3c,
	}
	buf := make([]byte, len(expected))
	_, err = serverConn.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, expected, buf)
}
