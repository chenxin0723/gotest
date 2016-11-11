package main

import (
  // "os"
"github.com/chenxin0723/gotest/synn"
  // "fmt"
)


func main() {

  // outfile, _ := os.Create("test.txt")
  // defer outfile.Close()
  // outfile.WriteString("my name is  xinchen \n")
  // outfile.Write([]byte("so should it matter \n"))
  // // os.Remove("test.txt")
  // file, _ := os.Open("test.txt")
  // defer file.Close()
  // bytes := make([]byte, 1024)
  // n, _ := file.Read(bytes)
  // os.Stdout.Write(bytes[:n])


println("111")

man := synn.Man{}
// println(&Man{})
// println(man)
// man.syn()

student := synn.Student{"xinchen", &man}

println(student.man)
// &Man{}
// man.syn()
// defer sync.Mutex.unlock()
println("222")

}
