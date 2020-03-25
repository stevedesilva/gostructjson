package gamestore

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

// ErrSaveGame error
var ErrSaveGame = errors.New("Save game error")

// Item struct
type Item struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price int    `json:"price,omitempty"`
}

// Game struct
type Game struct {
	Item
	Genre string `json:"genre,omitempty"`
}

// Games struct exposed to client
type Games struct {
	Sc    *bufio.Scanner
	Games map[int]Game
}

// New Constructor
func New(reader io.Reader, size int) *Games {
	return &Games{
		Sc:    bufio.NewScanner(reader),
		Games: make(map[int]Game, size),
	}
}

// Add game
func (g *Games) Add(id, price int, name, genre string) {
	g.Games[id] = Game{Item{id, name, price}, genre}
}

// List games
func (g *Games) List() []string {
	result := make([]string, 0, len(g.Games))
	for _, gm := range g.Games {
		res := fmt.Sprintf("#%d: %-15q %-20s $%d\n", gm.ID, gm.Name, "("+gm.Genre+")", gm.Price)
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
	res := fmt.Sprintf("#%d: %-15q %-20s $%d\n", gm.ID, gm.Name, "("+gm.Genre+")", gm.Price)
	return res
}

// Search for item
func (g *Games) Search(in string) (found bool) {
	for _, v := range g.Games {
		if v.Name == in {
			return true
		}
	}
	return
}

// Save EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------
func (g *Games) Save() (string, error) {
	result := make([]Game, len(g.Games))
	for _, gm := range g.Games {
		result[gm.ID-1] = gm
	}

	out, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return "", ErrSaveGame
	}
	return string(out), nil

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
> save   : exports the data to json and quits

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

		case "save":
			v, err := g.Save()
			if err != nil {
				continue
			}
			result = append(result, v)
			fmt.Println("Save ", result)
			return

		default:
			fmt.Println("Default " + in)
		}

	}

	return
}
