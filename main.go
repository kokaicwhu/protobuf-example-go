package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf-example-go/hello"
	"protobuf-example-go/proto/enumpb"
	simplepb "protobuf-example-go/proto/simple"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {

	hello.Hello()

	simple := doSimple()
	writeToFile("simple.bin", simple)

	simple2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", simple2)

	simple2Text := toJSON(simple2)
	fmt.Println(simple2Text)

	simple3 := &simplepb.SimpleMessage{}
	fromJSON(simple2Text, simple3)
	fmt.Println("Successfully create the simple3:", simple3)

	doEnum()
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Cannot serialize to file")
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Cannot write to file", err)
		return err
	}

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Failed to read the file", err)
		return err
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Cannot put bytes into the buffer strcut", err)
		return err
	}

	return nil
}

func doSimple() *simplepb.SimpleMessage {

	simple := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Simple Name",
		SimpleList: []int32{1, 4, 7, 11},
	}

	return &simple
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Cannot convert ot JSON", err)
	}

	return out
}

func fromJSON(in string, pb proto.Message) {
	if err := jsonpb.UnmarshalString(in, pb); err != nil {
		log.Fatalln("Cannot unmarshal the JSON into the pb strcut", err)
	}
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:  42,
		Day: enumpb.DayOfTheWeek_FRIDAY,
	}

	fmt.Println(em)
}
