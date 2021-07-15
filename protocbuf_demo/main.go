package main

import (
	address "./address/proto"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func main() {
	var cb address.ContactBook
	p1 := address.Person{
		Id:     1,
		Name:   "zhuyuchen",
		Gender: address.GenderType_MALE,
		Number: "18902406220",
	}
	fmt.Println("p1", p1)

	cb.Persons = append(cb.Persons, &p1)
	data, err := proto.Marshal(&p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v\n", err)
		return
	}
	ioutil.WriteFile("./proto.dat", data, 0644)
	data2, err := ioutil.ReadFile("./proto.dat")
	if err != nil {
		fmt.Printf("read file failed,err:%v\n", err)
		return
	}
	var p2 address.Person
	proto.Unmarshal(data2, &p2)
	fmt.Printf("p2: %#v",p2)

}
