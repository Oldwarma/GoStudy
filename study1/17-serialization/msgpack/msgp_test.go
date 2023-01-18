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

	msg2000.Body = "wcaonima\"devDateTime\":\"2022-12-28 16:49:37.597\",\n\"softVersion\":\"4.11.3\",\n\"phy\":\"3844\",\n\"distro\":\"Linaro\",\n\"virt\":\"0\",\n\"cpus\":\"6\",\n\"kernel\":\"4.4.194\",\n\"devStDateTime\":\"2022-12-2714:39:00.000\",\n\"temLow\":\"55\",\n\"version\":\"4.4.194\",\n\"frequency\":\"1.42\",\n\"temHigh\":\"66\",\n\"edgeId\":\"00010002000300040005011\",\n\"disk\":\"14047\",\n\"diskLmt\":\"5\",\n\"devRunTime\":\"94237\",\n\"tempValue\":\"30.00\",\n\"cpuRate\":\"15.55\",\n\"diskUsed\":" +
		"\"82.02\",\n\"arch\":\"arm\",\n\"cpuLmt\":\"5\",\n\"memUsed\":\"0.33\",\n\"memLmt\":\"5\",\n\"status\":\"0\""
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
