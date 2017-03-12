package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	fooBuf, err := proto.Marshal(&Foo{
		Type: Type_FOO,
		Foo:  "I am a foo",
	})
	check(err)
	handler(fooBuf)

	barBuf, err := proto.Marshal(&Bar{
		Type: Type_BAR,
		Bar:  1234,
	})
	check(err)
	handler(barBuf)
}

func handler(buf []byte) {
	// buf can be unmarshalled to TypeHeader just to check the type
	var out TypeHeader
	err := proto.Unmarshal(buf, &out)
	check(err)
	switch out.Type {
	case Type_FOO:
		var foo Foo
		proto.Unmarshal(buf, &foo)
		fmt.Println("Foo:", foo.Foo)
	case Type_BAR:
		var bar Bar
		proto.Unmarshal(buf, &bar)
		fmt.Println("Bar:", bar.Bar)
	}
}

// just to keep err checking simple for a simple demo
func check(err error) {
	if err != nil {
		panic(err)
	}
}
