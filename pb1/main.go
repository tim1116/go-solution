package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"pb1/service"
)

func main() {
	test := &service.Student{
		Name:   "Bob",
		Male:   true,
		Scores: []int32{98, 85, 88},
	}
	// 编码
	data, err := proto.Marshal(test)
	if err != nil {
		log.Panicf("err: %s", err)
	}

	//fmt.Println(data)

	newTest := &service.Student{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}

	fmt.Println(newTest.GetName())
	fmt.Println(newTest.GetMale())
	fmt.Println(newTest.GetScores())
}
