// automatically generated by the FlatBuffers compiler, do not modify

package TestProtocol

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type TestPing struct {
	_tab flatbuffers.Table
}

func GetRootAsTestPing(buf []byte, offset flatbuffers.UOffsetT) *TestPing {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &TestPing{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *TestPing) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TestPing) IntValue() int32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *TestPing) MutateIntValue(n int32) bool {
	return rcv._tab.MutateInt32Slot(4, n)
}

func TestPingStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func TestPingAddIntValue(builder *flatbuffers.Builder, intValue int32) {
	builder.PrependInt32Slot(0, intValue, 0)
}
func TestPingEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
