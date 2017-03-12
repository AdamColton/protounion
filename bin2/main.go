package main

import (
	"fmt"
	"github.com/adamcolton/protounion"
	"github.com/golang/protobuf/proto"
)

const (
	FOOBAR = uint32(iota + 100)
	BAZ
)

func main() {
	foobarBuf, err := proto.Marshal(&Foobar{
		Type: FOOBAR,
		Foo:  "I am a foobar",
		Bar:  1234,
	})
	check(err)
	handler(foobarBuf)

	bazBuf, err := proto.Marshal(&Baz{
		Type: BAZ,
		Baz:  true,
	})
	check(err)
	handler(bazBuf)
}

func handler(buf []byte) {
	// buf can be unmarshalled to TypeHeader just to check the type
	var out protounion.TypeHeader
	err := proto.Unmarshal(buf, &out)
	check(err)
	switch out.Type {
	case FOOBAR:
		var foobar Foobar
		proto.Unmarshal(buf, &foobar)
		fmt.Println("Foo:", foobar.Foo, "Bar:", foobar.Bar)
	case BAZ:
		var baz Baz
		proto.Unmarshal(buf, &baz)
		fmt.Println("Baz:", baz.Baz)
	}
}

// just to keep err checking simple for a simple demo
func check(err error) {
	if err != nil {
		panic(err)
	}
}
