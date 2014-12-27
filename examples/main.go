package main

import "log"
import "reflect"
import . "gopkg.in/nowk/go-calm.v1"

type S struct{}

func main() {
	v := reflect.ValueOf(S{})
	err := Calm(func() {
		v.Len()
	})

	log.Printf("recovered: %s", err)
}
