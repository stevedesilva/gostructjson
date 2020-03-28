package main

import (
	"os"

	gs "github.com/stevedesilva/gostructjson/gamestore"
)

func main() {

	g := gs.New(os.Stdin, 3)

	g.Run()

}
