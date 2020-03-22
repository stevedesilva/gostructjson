package gamestore

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

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

// Games struct exposed to client
type Games struct {
	Sc *bufio.Scanner
	// Games []game
	Games map[int]game
}

// New Constructor
func New(reader io.Reader, size int) *Games {
	return &Games{
		Sc:    bufio.NewScanner(reader),
		Games: make(map[int]game, size),
	}
}

// Add game
func (g *Games) Add(id, price int, name, genre string) {
	g.Games[id] = game{item{id, name, price}, genre}
}

// List games
func (g *Games) List() []string {
	result := make([]string, 0, len(g.Games))
	for _, gm := range g.Games {
		res := fmt.Sprintf("#%d: %-15q %-20s $%d\n", gm.id, gm.name, "("+gm.genre+")", gm.price)
		result = append(result, res)
	}
	sort.Strings(result)
	return result
}

// GetByID EXERCISE: Query By Id
//
//  Add a new command: "id". So the users can query the games
//  by id.
//
//  1. Before the loop, index the games by id (use a map).
//
//  2. Add the "id" command.
//     When a user types: id 2
//     It should print only the game with id: 2.
//
//  3. Handle the errors:
//
//     id
//     wrong id
//
//     id HEY
//     wrong id
//
//     id 10
//     sorry. i don't have the game
//
//     id 1
//     #1: "god of war" (action adventure) $50
//
//     id 2
//     #2: "x-com 2" (strategy) $40
//
//
// EXPECTED OUTPUT
//  Please also run the solution and try the program with
//  list, quit, and id commands to see it in action.
// ---------------------------------------------------------
func (g *Games) GetByID(id int) string {
	gm := g.Games[id]
	res := fmt.Sprintf("#%d: %-15q %-20s $%d\n", gm.id, gm.name, "("+gm.genre+")", gm.price)
	return res
}

// Search for item
func (g *Games) Search(in string) (found bool) {
	for _, v := range g.Games {
		if v.name == in {
			return true
		}
	}
	return
}

// Run listens for user command ( List Search or Quit) and response with appropriate response
func (g *Games) Run() (result []string) {

GamesLoop:
	for {
		fmt.Printf(`
> List  : prints all games
> quit  : quit 
> search: games list for item
> id [num]: return the game for this id

		`)

		if !g.Sc.Scan() {
			break GamesLoop
		}

		args := strings.Fields(g.Sc.Text())
		fmt.Println(">>> Args = ", args)
		if len(args) == 0 {
			msg := "Continue: invalid args length"
			result = append(result, msg)
			fmt.Println(msg, len(args))
			continue
		}

		switch in := args[0]; in {
		case "quit":
			msg := "Bye!"
			fmt.Println(msg)
			result = append(result, msg)
			return

		case "list":
			for _, m := range g.List() {
				result = append(result, m)
				fmt.Println(m)
			}
		case "search":
			found := g.Search(in)
			result = append(result, "Found ", strconv.FormatBool(found))
			fmt.Println("Found ", found)

		case "id":
			if len(args) != 2 {
				fmt.Println("Wrong id")
				result = append(result, "Missing args")
				continue GamesLoop
			}
			v, err := strconv.Atoi(args[1])
			if err != nil {
				result = append(result, "Missing id")
				continue GamesLoop
			}

			gm := g.GetByID(v)
			fmt.Println("Found ", gm)
			result = append(result, gm)

		default:
			fmt.Println("Default " + in)
		}

	}

	return
}
