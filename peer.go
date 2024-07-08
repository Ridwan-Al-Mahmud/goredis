package main

import (
	"net"
	"log/slog"
)

type Peer struct {
	conn  net.Conn
	msgCh chan Message
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.conn.Write(msg)
}

func NewPeer(conn net.Conn, msgCh chan Message) *Peer {
	return &Peer {
		conn: conn,
		msgCh: msgCh,
	}
}

func (p *Peer) readLoop() error {
	buf := make([]byte, 1024)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			slog.Error("Read error", "err", err)
			return err
		}
		msgBuf := buf[:n]
		copy(msgBuf, buf[:n])
		p.msgCh <- Message {
			data: msgBuf,
			peer: p,
		}
	}
}