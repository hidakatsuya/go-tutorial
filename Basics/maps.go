package main

import "fmt"

type Foo struct {
	Lat, Long float64
}

var m map[string]Foo

func main() {
	m = make(map[string]Foo)
	m["Bell Labs"] = Foo{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
