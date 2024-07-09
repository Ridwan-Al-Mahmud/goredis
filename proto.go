package main

import (
	"fmt"
	"bytes"
	"github.com/tidwall/resp"
)

const (
	CommandSET = "SET"
	CommandGET = "GET"
	CommandHELLO = "hello"
	CommandClient = "client"
)

type Command interface {
	
}

type HelloCommand struct {
	value string
}

type ClientCommand struct {
	value string
}

type SetCommand struct {
	key, val []byte
}

type GetCommand struct {
	key []byte
}

func respWriteMap(m map[string]string) []byte {
	buf := &bytes.Buffer{}
	buf.WriteString("%" + fmt.Sprintf("%d\r\n", len(m)))
	rw := resp.NewWriter(buf)
	for k, v := range m {
		rw.WriteString(k)
		rw.WriteString(":"+ v)
	}
	return buf.Bytes()
}