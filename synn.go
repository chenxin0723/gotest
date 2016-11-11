package synn

import (
  "sync"
)

type Student struct {

 name string
 man  *Man

}

type Man struct {
  lock sync.Mutex

}

func (man *Man) syn() {
  man.lock.Lock()
  println("syn")
  println(man)
}


// func main() {


// println("111")

// man := Man{}
// println(&Man{})
// // println(man)
// man.syn()

// student := Student{"xinchen", &man}

// println(student.man)
// // &Man{}
// // man.syn()
// // defer sync.Mutex.unlock()
// println("222")

// }
