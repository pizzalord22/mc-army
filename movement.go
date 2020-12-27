package main

import (
	"fmt"
	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/world/entity/player"
	"github.com/beefsack/go-astar"
)

// called every time the bot moves
func onMove(p player.Pos) error {
	return nil
}

// find the best route to the destination
func mapRoute(start, destination bot.Position) []bot.Position {
	var routes [][]bot.Position
	// todo find all possible routes to the destination

	// todo select the best route
	_ = routes
	return []bot.Position{}
}

// todo translate the resulting path into mc instructions
//  this should probably only take care of straight lines (diagonals are straight lines)

// get the most efficient path to a given location
func Astar(w World) {
	p, dist, found := astar.Path(w.From(), w.To())
	if !found {
		fmt.Println("Could not find a path")
		return
	}

	fmt.Printf("Resulting path\n%s\nDistance: %f\n", w.RenderPath(p), dist)
	for _,v := range p{
		x := v.(*Tile).X
		y := v.(*Tile).Z
		fmt.Printf("x: %d, y: %d\n",x,y)
	}
}
