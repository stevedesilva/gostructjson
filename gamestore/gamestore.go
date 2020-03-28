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

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 30
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

// ErrSaveGame error
var ErrSaveGame = errors.New("Save game error")

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
	Sc    *bufio.Scanner
	Games map[int]game
}

type jsonGame struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Genre string `json:"genre"`
	Price int    `json:"price,omitempty"`
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

	result := make([]jsonGame, 0, len(g.Games))

	for _, gm := range g.Games {
		result = append(result,
			jsonGame{ID: gm.id, Name: gm.name, Genre: gm.genre, Price: gm.price})
	}

	sort.Slice(result, func(a, b int) bool {
		return result[a].ID < result[b].ID
	})

	out, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		return "", ErrSaveGame
	}
	return string(out), nil

}

// Run listens for user command ( List Search or Quit) and response with appropriate response
func (g *Games) Run() (result []string) {

	var decoded []jsonGame
	// Umarshal needs users as a pointer in order to update
	if err := json.Unmarshal([]byte(data), &decoded); err != nil {
		fmt.Println(err)
		return
	}
	// init map
	for _, v := range decoded {
		g.Add(v.ID, v.Price, v.Name, v.Genre)
	}

GamesLoop:
	for {
		fmt.Printf(`
> list  : prints all games
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
			p := strings.Join(args[1:], " ")

			found := g.Search(string(p))
			result = append(result, "Found ", strconv.FormatBool(found))
			fmt.Println("Search for ", in, " was ", found)

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
