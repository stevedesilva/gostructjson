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

	sc := bufio.NewScanner(os.Stdin)

GamesLoop:
	for {
		fmt.Printf(`
> List  : prints all games
> Quit  : quit 
> Search: games list for item

		`)

		if !sc.Scan() {
			break GamesLoop
		}

		switch in := sc.Text(); in {
		case "Quit":
			fmt.Println("Exiting " + in)
			return

		case "List":
			for _, m := range g.List() {
				fmt.Println(m)
			}

		default:
			if found := g.Search(in); found {
				fmt.Println("Found " + in)
			} else {
				fmt.Println("Not Found " + in)
			}

		}

	}

}
