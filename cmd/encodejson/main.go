package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	name     string
	password string
	permissions
}

func main() {
	users := []user{
		{"steve", "111", nil},
		{"clive", "222", permissions{"admin": true}},
		{"ben", "333", permissions{"write": true}},
	}

	out, err := json.Marshal(users) // only encodes exported fields
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(out))

}
