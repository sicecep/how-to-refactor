package main

import (
	"testing"
)

var player1 = &Player{"Sepry", 0}
var player2 = &Player{"Haryandi", 0}
var g = TennisGame{player1, player2}

func TestFirstPlayerWinFirstGame(t *testing.T) {
	g.AddPoint(player1)
	g.AddPoint(player1)

	t.Log(g.printScore())
	t.Log(g.GetScore())

	Teardown()
}

func Teardown() {
	// reset players point for every test
	player1.Point = 0
	player2.Point = 0
}
