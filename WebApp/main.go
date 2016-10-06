// https://go-macaron.com/
// adapted from minimal example

package main

import "gopkg.in/macaron.v1"

func main() {

  m := macaron.Classic()
  m.Get("/hello", func() string {
      return "Hello world!"
  })

  m.Get("/reverse", func() string {
      return "Hello From Reverse"
  })
  m.Run()

} // main()
