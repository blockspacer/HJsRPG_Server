// automatically generated by the FlatBuffers compiler, do not modify

package 

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ClientGameServer struct {
	_tab flatbuffers.Struct
}

func (rcv *ClientGameServer) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func CreateClientGameServer(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	builder.Prep(1, 0)
	return builder.Offset()
}