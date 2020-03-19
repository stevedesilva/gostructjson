package gamestore

import "fmt"

// Item struct
type item struct {
	id    int
	name  string
	price int
}

// Game struct
type game struct {
	item
	genre string
}

// Games struct
type Games struct {
	games []game
}

// New Constructor
func New() *Games {
	return &Games{}
}

// Add game
func (g *Games) Add(id, price int, name, genre string) {

	game := game{item{id, name, price}, genre}
	g.games = append(g.games, game)
}

// List games
func (g *Games) List() []string {
	result := make([]string, len(g.games))
	for _, gm := range g.games {
		res := fmt.Sprintf("#%d: %-15q %-20s $%d\n", gm.id, gm.name, "("+gm.genre+")", gm.price)
		result = append(result, res)
	}
	return result
}

// Search for item
func (g *Games) Search(in string) (found bool) {
	for _, v := range g.games {
		if v.name == in {
			return true
		}
	}
	return
}
