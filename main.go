package main

import (
	"github.com/dakamarko/demo4/impl"
	"sigs.k8s.io/yaml"
)

type Player struct {
	impl.DataHolder
	Name string
}

func main() {
	var p Player

	input := []byte("Name: Marko")
	err := yaml.Unmarshal(input, &p)
	if err != nil {
		panic("oh ooh")
	}

	p.Set(input)

}
