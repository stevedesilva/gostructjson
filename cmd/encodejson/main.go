package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

/*
 - A struct type combines different types of fields in a single type. You can use a struct type to represent a concept.
 - each field should have a unique name. Also, each field should have a type, but every field can have a different type.
 - fields must be unique
 - Go initializes a struct's fields to zero-values depending on their type.
 - Field names and types are part of a struct's type.
 - When creating a struct value, it doesn't matter whether you use the field names or not. So, they are equal.
 - when you omit some of the fields, Go assigns zero values to them. Here, "clone" struct value's "rating" and "released" fields are: 0, and false, respectively.
 - Types with different names cannot be compared. However, you can convert one of them to the other because they have the same set of fields. movie{} == movie(performance{}) is ok, or vice versa.
 - `m.title` returns "avengers: end game" because the outer type always takes priority. However, `m.item.title` returns "midnight in paris" because you explicitly get it from the inner type: item.
		type item struct{ title string }

		type movie struct {
			item
			title string
		}

		m := movie{
			title: "avengers: end game",
			item:  item{"midnight in paris"},
		}

		fmt.Println(m.title, "&", m.item.title)

 - tag field
	the json package can read and encode/decode depending on the associated metadata. It's just a string value.
	It's only meaningful when other code reads it. For example, the json package can read it and encode/decode depending on the field tag's value.
	The json package can only encode exported fields.
 - Why do you need to pass a pointer to the Unmarshal function?
 Otherwise, it would not be able to update the given value. It's because, every value in Go is passed by value. So a function can only change the copy, not the original value. However, through a pointer, a function can change the original value.
*/

// ield tags metadata encode and decode field
// normally a raw string literal since it includes ""
// key value pair, key is package = json, tag is read by package
type user struct {
	Name        string      `json:"username"`        // change json display name
	Password    string      `json:"-"`               // don't encode Password
	Permissions permissions `json:"perms,omitempty"` // change encoded field to perms, does not encode empty values
}

type permissions map[string]bool

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
