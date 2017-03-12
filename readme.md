## Protobuf Union

Sometime you need to send or receive messages that could be several different
types. With protobufs, this presents a challenge because you need to know the
type before you can unmarshal it.
[This](http://stackoverflow.com/questions/9121612/protocol-buffers-detect-type-from-raw-message)
[is](https://groups.google.com/forum/#!topic/protobuf/t-j-GIPPbEk)
[a very](http://stackoverflow.com/questions/3230841/api-to-know-the-kind-of-protobuf-message-being-sent-in)
[common](http://stackoverflow.com/questions/30564404/how-to-determine-message-type-in-protobuf-so-that-i-can-use-that-type-parsefrom)
[question](http://stackoverflow.com/questions/32639905/protobuf-determining-message-type-to-deserialize).

The most common solution is the
[union technique](https://developers.google.com/protocol-buffers/docs/techniques#union)
proposed on the Protobuf page. The other two solutions are to add a fixed length
ID to the serialized buffer that can be inspected separately. Or create a
protobuf wrapper.

Part of what makes this problem tricky is that the Protobuf standard doesn't
encode the type, so there is no way to get at it. But this, along with the
protection for adding and removing fields and cross-language compatibility
creates another technique.

Two protobuf message types can be used interchangeably so long as the types
agree where their tags intersect. This means we can define a sort of abstract
type message, and other messages can implement it by having those fields match
up.

In this example, TypeHeader is our abstract message. Both Foo and Bar implement
TypeHeader. So a protobuf generated from either Foo or Bar can be unmarshalled
to TypeHeader, the Type can be inspected and then the correct message can be
used to unmarshal the entire message.

I'm not entirely sure this is a good idea, but because it's not just a fluke of
the implementation. The same choices that make it safe to add fields to protobuf
message protect this technique.