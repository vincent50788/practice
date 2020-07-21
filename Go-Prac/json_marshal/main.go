package main

import (
	"encoding/json"
	"log"
)

func main() {
	type Test struct {
		S string "json:s"
		b bool   "json:b"
		i int    "json:i"
	}

	t := Test{
		S: "harvey",
		b: false,
		i: 123,
	}

	out, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
	}

	tmp := &Test{}

	b := []byte(`{"s":"Bob","b":"true"}`)

	json.Unmarshal(b, tmp)

	log.Printf("out: %b", out)
	log.Printf("tmp: %v", tmp)
	log.Println(out)
}
