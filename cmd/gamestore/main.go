package main

import (
	"bufio"
	"fmt"
	"os"

	gs "github.com/stevedesilva/gostructjson/gamestore"
)

func main() {

	g := gs.New()
	g.Add(1, 50, "god of war", "action adventure")
	g.Add(2, 30, "x-com 2", "strategy")
	g.Add(3, 20, "minecraft", "sandbox")

	for _, m := range g.List() {
		fmt.Println(m)
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		in := sc.Text()

		switch {
		case "Quit" == in:
			{
				fmt.Println("Exiting " + in)
				return
			}

		default:
			{
				if found := g.Search(in); found {
					fmt.Println("Found " + in)
				} else {
					fmt.Println("Not Found " + in)
				}
			}
		}

	}

}
