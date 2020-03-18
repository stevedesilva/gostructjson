package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type permissions map[string]bool

// ield tags metadata encode and decode field
// normally a raw string literal since it includes ""
// key value pair, key is package = json, tag is read by package
type user struct {
	Name        string      `json:"username"`        // change json display name
	Password    string      `json:"-"`               // don't encode Password
	Permissions permissions `json:"perms,omitempty"` // change encoded field to perms, does not encode empty values
}

func main() {

	// encoder()
	decoder()
}

/*
		[
	        {
	                "username": "steve"
	        },
	        {
	                "username": "clive",
	                "perms": {
	                        "admin": true
	                }
	        },
	        {
	                "username": "ben",
	                "perms": {
	                        "write": true
	                }
	        }
	]
*/
func encoder() {
	users := []user{
		{"steve", "111", nil},
		{"clive", "222", permissions{"admin": true}},
		{"ben", "333", permissions{"write": true}},
	}

	// only encodes exported fields, pretty prints

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(string(out))
}

func decoder() {
	var input []byte

	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		input = append(input, in.Bytes()...)
	}

	var users []user
	// Umarshal needs users as a pointer in order to update
	if err := json.Unmarshal(input, &users); err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range users {
		fmt.Print("+ " + user.Name)
		switch p := user.Permissions; {
		case p == nil:
			fmt.Print(" has no power")
		case p["admin"]:
			fmt.Print(" is an admin")
		case p["write"]:
			fmt.Print(" can write")

		}
		fmt.Println()
	}

}
