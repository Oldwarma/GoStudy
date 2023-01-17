package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/tinylib/msgp/msgp"
	"testing"
)

var doc = Doc{DocId: 123, Position: "搜索工程师", Company: "百度",
	City: "北京"}
var person = Person{DocId: 456, Position: "几年太么", Company: "急急急挖",
	City: "北京"}

type Doc struct {
	DocId       int
	Position    string
	Company     string
	SchoolLevel int
	City        string
}
type Person struct {
	DocId       int
	Position    string
	Company     string
	SchoolLevel int
	City        string
}

func (p Person) EncodeMsg(writer *msgp.Writer) error {
	//TODO implement me
	panic("implement me")
}

func TestJson(t *testing.T) {
	bs, _ := json.Marshal(doc)
	fmt.Printf("json encode byte length %d\n", len(bs))
	var inst Doc
	_ = json.Unmarshal(bs, &inst)
	fmt.Printf("json decode position %s\n", inst.Position)
}
func TestGob(t *testing.T) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	_ = encoder.Encode(doc)
	fmt.Printf("gob encode byte length %d\n", len(buffer.Bytes()))
	var inst Doc
	decoder := gob.NewDecoder(&buffer)
	_ = decoder.Decode(&inst)
	fmt.Printf("gob decode position %s\n", inst.Position)
}

//func TestGogoProtobuf(t *testing.T) {
//	bs, _ := proto.Marshal(&pb.Doc{})
//	fmt.Printf("pb encode byte length %d\n", len(bs))
//	var xx Message
//	_ = proto.Unmarshal(bs, xx)
//	fmt.Printf("pb decode position %s\n")
//}

//func TestMsgp(t *testing.T) {
//	var buf bytes.Buffer
//	_ = msgp.Encode(&buf, &person)
//	fmt.Printf("msgp encode byte length %d\n", len(buf.Bytes()))
//	var inst Person
//	_ = msgp.Decode(&buf, &inst)
//	fmt.Printf("msgp decode position %s\n", inst.Position)
//}

type Message interface {
	Reset()
	String() string
	ProtoMessage()
}

func BenchmarkJsonEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		json.Marshal(doc)
	}
}
func BenchmarkJsonDecode(b *testing.B) {
	bs, _ := json.Marshal(doc)
	var inst Doc
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal(bs, &inst)
	}
}

func BenchmarkGobEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		encoder := gob.NewEncoder(&buffer)
		encoder.Encode(doc)
	}
}
func BenchmarkGobDecode(b *testing.B) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(doc)
	var inst Doc
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer.Reset()
		decoder := gob.NewDecoder(&buffer)
		decoder.Decode(&inst)
	}
}

func BenchmarkMsgpEncode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		msgp.Encode(&buf, &person)
	}
}
