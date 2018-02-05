package main

import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

type Student struct {
	id   int
	name string
	age  int
}

func (s *Student)showStu() {
	fmt.Println("id ==> ", s.id)
	fmt.Println("name ==> ", s.name)
	fmt.Println("age ==> ", s.age)
}

func (p *Person)showPerson() {
	fmt.Println("id ==> ", p.Id)
	fmt.Println("name ==> ", p.Name)
	fmt.Println("age ==> ", p.Age)
}

// toJSON  json.Marshal(st)
// toBytes json.Unmarshal([]byte, &st)
func main() {
	// p1 := &Person{3, "John", 12}
	// p2 := &Person{id: 1, name: "John2"}
	// fmt.Println(*p2)
	// showPersonJson(*p2)
	// var p3Bytes []byte = toJson(p2)
	// fmt.Println("p3Bytes", p3Bytes)
	// p4 := &Person{}
	// err := json.Unmarshal(p3Bytes, &p4)
	// if err != nil { panic(err) }
	// showPersonJson(p4)
	p1 := &Person{1, "Kevin", 12}
	p1bytes, err := json.Marshal(p1)
	if err != nil { panic(err) }

	fmt.Println(string(p1bytes))


	s1 := &Student{1, "Tom", 44}
	s1bytes, err := json.Marshal(s1)
	if err != nil { panic(err) }

	s1.showStu()

	fmt.Println(string(s1bytes))

	p2 := &Person{}
	err = json.Unmarshal(p1bytes, &p2)
	if err != nil { panic(err) }

	fmt.Println("showPerson:")
	p2.showPerson()

}

func showPersonJson(p Person) {
	pJson, err := json.Marshal(p)
	if err != nil { panic(err) }
	fmt.Println(string(pJson))
}

func toJson(p *Person) []byte {
	pJson, err := json.Marshal(*p)
	if err != nil { panic(err) }
	return pJson
}