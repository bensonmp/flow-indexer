package store

import (
	"reflect"
	"testing"
)

var ids = DocumentList{1, 2, 3, 5, 8, 13, 21}
var expectIDs = DocumentList{0, 1, 2, 3, 5, 8, 13, 21}

func runCodecTest(t *testing.T, codecFactory func() Codec) {
	c := codecFactory()
	c.AddID(0)
	for _, id := range ids {
		b, _ := c.Bytes()
		c.FromBytes(b)
		c.AddID(id)
	}
	resultDocs := c.Documents()
	if !reflect.DeepEqual(expectIDs, resultDocs) {
		t.Errorf("codec(%s)=> %v, want %v", c, resultDocs, expectIDs)
	}

}

func runCodecBench(b *testing.B, codecFactory func() Codec) {
	for n := 0; n < b.N; n++ {
		c := codecFactory()
		c.AddID(0)
		for _, id := range ids {
			b, _ := c.Bytes()
			c.FromBytes(b)
			c.AddID(id)
		}
		c.Bytes()
	}
}

func TestCodec(t *testing.T) {
	//runCodecTest(t, NewBitsetCodec)
	//runCodecTest(t, NewMsgpackCodec)
	runCodecTest(t, func() Codec { return NewBitsetCodec() })
	runCodecTest(t, func() Codec { return NewMsgpackCodec() })
	runCodecTest(t, func() Codec { return NewMsgpackDeltasCodec() })
}

func BenchmarkBitsetCodec(b *testing.B) {
	runCodecBench(b, func() Codec { return NewBitsetCodec() })
}
func BenchmarkMsgpackCodec(b *testing.B) {
	runCodecBench(b, func() Codec { return NewMsgpackCodec() })
}

func BenchmarkMsgpackDeltaCodec(b *testing.B) {
	runCodecBench(b, func() Codec { return NewMsgpackDeltasCodec() })
}