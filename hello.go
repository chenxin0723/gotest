package main

// import "github.com/astaxie/beego"

import (
  "fmt"
  // "runtime"
  "net/http"
  "reflect"
  "io/ioutil"
)

type Human struct {
  name string
  age int
  phone string
}

type skill []string

type Student struct {
  *Human
  school string
  loan float32
  skill
  int
}

type Employee struct {
  Human
  company string
  money float32
}


func (h *Human) myname() {
  h.name = "wdd"
  fmt.Println(h.name)
}

// func (h Student) myname() {
//   // h.name = "wdd1"
//   fmt.Println(h.name)
// }

func (s Student) myschool() string{
  // fmt.Println(s.school)
  return s.school
}


type Name interface{
  myname()
}

type School interface{
  myschool() string
}



func sing(s School) string {
 return s.myschool()
}

func main() {
  var s School
  student := Student{Human: &Human{name: "chenxin", age: 32, phone: "18357173671"}, school: "beego", skill: []string{"aa", "cc"}}
  s = student
  // student.myschool()
  fmt.Println(s.myschool())
  student.myname()
  fmt.Println(student.Human.name)

resp, _ := http.Get("http://example.com/")


fmt.Println(reflect.ValueOf(resp.Body))

body, _ := ioutil.ReadAll(resp.Body)


fmt.Println(body)
fmt.Println(http.ErrHeaderTooLong)


ss := sing(student)

fmt.Println(ss)
}
// func main(){
//     beego.Run()
// }
