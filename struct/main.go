package main

import (
	"fmt"
	"github.com/fatih/structs"
)

type UserInfo struct {
	Name string `json:"name" structs:"name"`
	Age  int    `json:"age" structs:"age"`
	Profile `json:"profile" structs:"profile"`
}
type Profile struct {
	Hobby string `json:"hobby" structs:"hobby"`
}
func main()  {
	u1 := UserInfo{Name: "zhuyuchen", Age: 18, Profile: Profile{"双色球"}}
	m1:=structs.Map(&u1)
	fmt.Println(m1)
}
