package main

import (
	"fmt"
	"other/test/singleflight"
)

func main() {
	var g singleflight.Group
	for i := 0; i < 10; i++ {
		v, err := g.Do("key", func() (interface{}, error) {
			fmt.Println("call bar")
			return "bar", nil
		})

		if v != "bar" || err != nil {
			fmt.Errorf("Do v = %v, error = %v", v, err)
		}
	}
}
