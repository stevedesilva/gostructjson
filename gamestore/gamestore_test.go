package gamestore_test

import (
	"testing"

	gs "github.com/stevedesilva/gostructjson/gamestore"
	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------
// EXERCISE: Warm Up
//
//  Starting with this exercise, you'll build a command-line
//  game store.
//
//  1. Declare the following structs:
//
//     + item: id (int), name (string), price (int)
//
//     + game: embed the item, genre (string)
//
//
//  2. Create a game slice using the following data:
//
//     id  name          price    genre
//
//     1   god of war    50       action adventure
//     2   x-com 2       30       strategy
//     3   minecraft     20       sandbox
//
//
//  3. Print all the games.
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------
func TestGamestore_created(t *testing.T) {
	g := gs.New()
	assert.NotNil(t, g)

	g.Add(1, 50, "god of war", "action adventure")
	g.Add(2, 30, "x-com 2", "strategy")
	g.Add(3, 20, "minecraft", "sandbox")

	got := g.List()
	want := []string{"", "", "", "#1: \"god of war\"    (action adventure)   $50\n", "#2: \"x-com 2\"       (strategy)           $30\n", "#3: \"minecraft\"     (sandbox)            $20\n"}
	assert.Equal(t, want, got)

}

// ---------------------------------------------------------
// EXERCISE: List
//
//  Now, it's time to add an interface to your program using
//  the bufio.Scanner. So the users can list the games, or
//  search for the games by id.
//
//  1. Scan for the input in a loop (use bufio.Scanner)
//
//  2. Print the available commands.
//
//  3. Implement the quit command: Quits from the loop.
//
//  4. Implement the list command: Lists all the games.
//
//
// EXPECTED OUTPUT
//  Please run the solution and try the program with list and
//  quit commands.
// ---------------------------------------------------------
func TestGamestore_should_find_game(t *testing.T) {
	// use your solution from the previous exercise
	g := gs.New()
	assert.NotNil(t, g)

	g.Add(1, 50, "god of war", "action adventure")
	g.Add(2, 30, "x-com 2", "strategy")
	g.Add(3, 20, "minecraft", "sandbox")

	result := g.Search("god of war")
	assert.True(t, result)

	result = g.Search("god of peace")
	assert.False(t, result)
}

func TestGamestore_should_list_games(t *testing.T) {
	// use your solution from the previous exercise
	g := gs.New()
	assert.NotNil(t, g)

	g.Add(1, 50, "god of war", "action adventure")
	g.Add(2, 30, "x-com 2", "strategy")
	g.Add(3, 20, "minecraft", "sandbox")

}
