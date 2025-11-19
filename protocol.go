package sio

import (
	eioparser "github.com/kaiterra/socket.io-go/engine.io/parser"
	"github.com/kaiterra/socket.io-go/parser"
)

const (
	SocketIOProtocolVersion = parser.ProtocolVersion
	EngineIOProtocolVersion = eioparser.ProtocolVersion
)
