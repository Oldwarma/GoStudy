package main

import (
	"encoding/json"
	"github.com/vmihailenco/msgpack/v5"
	"gopkg.in/yaml.v3"
	"testing"
)

type message struct {
	Code int64
	Name string
	Body string
}

var msg2000 message
var msg100 message

var messages []message

func init() {
	msg2000 = message{
		Code: 1,
		Name: "Test",
	}

	msg2000.Body = "wcaonima"
	msg100 = message{
		Code: 2,
		Name: "Test",
	}
	msg100.Body = "wcaonima"

	for i := 0; i < 100; i++ {
		body := "wcaonima"
		messages = append(messages, message{
			Code: 1,
			Name: "Test",
			Body: body,
		})
	}
}

func Benchmark_MsgpackMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(msg2000)

	}
}

func Benchmark_JsonMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(msg2000)
	}
}

func Benchmark_YamlMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		yaml.Marshal(msg2000)
	}
}

func Benchmark_MsgpackUnmarshal(b *testing.B) {
	marshal, _ := msgpack.Marshal(msg2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := message{}
		msgpack.Unmarshal(marshal, &m)
	}
}

func Benchmark_JsonUnmarshal(b *testing.B) {
	marshal, _ := json.Marshal(msg2000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m := message{}
		json.Unmarshal(marshal, &m)
	}
}

func Benchmark_MsgpackMarshal100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(msg100)
	}
}

func Benchmark_JsonMarshal100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(msg100)
	}
}

func Benchmark_MsgpackMarshalMessages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		msgpack.Marshal(messages)

	}
}

func Benchmark_JsonMarshalMessages(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(messages)
	}
}
