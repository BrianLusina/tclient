package client

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/brianlusina/tclient/pkg/bitfield"
	"github.com/brianlusina/tclient/pkg/handshake"
	"github.com/brianlusina/tclient/pkg/message"
)

func recvBitfield(conn net.Conn) (bitfield.Bitfield, error) {
	if err := conn.SetDeadline(time.Now().Add(5 * time.Second)); err != nil {
		log.Fatalf("Failed to set deadline on connection %s", err)
	}

	defer func(c net.Conn) {
		// disable the deadline
		if err := c.SetDeadline(time.Time{}); err != nil {
			log.Fatalf("Failed to disable connection deadline %s", err)
		}
	}(conn)

	msg, err := message.Read(conn)
	if err != nil {
		return nil, err
	}

	if msg == nil {
		err := fmt.Errorf("Expected bitfield but got %s", msg)
		return nil, err
	}

	if msg.ID != message.MsgBitfield {
		err := fmt.Errorf("Expected bitfield but got id %d", msg.ID)
		return nil, err
	}

	return msg.Payload, nil
}

func completeHandshake(conn net.Conn, infohash, peerID [20]byte) (*handshake.Handshake, error) {
	if err := conn.SetDeadline(time.Now().Add(3 * time.Second)); err != nil {
		log.Fatalf("Failed to set deadline on connection %s", err)
	}

	defer func(c net.Conn) {
		// disable the deadline
		if err := c.SetDeadline(time.Time{}); err != nil {
			log.Fatalf("Failed to disable connection deadline %s", err)
		}
	}(conn)

	req := handshake.New(infohash, peerID)
	_, err := conn.Write(req.Serialize())
	if err != nil {
		return nil, err
	}

	res, err := handshake.Read(conn)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(res.InfoHash[:], infohash[:]) {
		return nil, fmt.Errorf("Expected infohash %x but gor %x", res.InfoHash, infohash)
	}

	return res, nil
}
