// https://go-macaron.com/
// adapted from minimal example

package main

import "gopkg.in/macaron.v1"

func main() {

  m := macaron.Classic()
  m.Get("/", func() string {
      return "Hello world!"
  })
  m.Run()

} // main()
