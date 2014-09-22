package main

import "log"
import "reflect"
import . "github.com/nowk/go-calm"

type S struct{}

func main() {
	v := reflect.ValueOf(S{})
	err := Calm(func() {
		v.Len()
	})

	log.Printf("recovered: %s", err)
}
