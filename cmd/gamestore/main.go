package main

import (
	"os"

	gs "github.com/stevedesilva/gostructjson/gamestore"
)

func main() {

	g := gs.New(os.Stdin,3)

	g.Add(1, 50, "god of war", "action adventure")
	g.Add(2, 30, "x-com 2", "strategy")
	g.Add(3, 20, "minecraft", "sandbox")

	g.Run()

}

