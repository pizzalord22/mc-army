package main

import (
    "fmt"
    "github.com/Tnze/go-mc/bot"
)

// get tge closest bock
func getClosestBlock(x, y, z float64, blocks [][3]float64) bot.Position {
    var closestCoord bot.Position
    var closest = 1e9
    fmt.Println(blocks)
    for _, v := range blocks {
        current := euclidean(x, y, z, v[0], v[1], v[2])
        if current < closest {
            closest = current
            closestCoord = bot.Position{
                X: int(v[0]),
                Y: int(v[1]),
                Z: int(v[2]),
            }
        }
    }
    return closestCoord
}

// find the closest tree
func findTree() {
    var woodBlocks [][3]float64
    for z := c.Pos.Z - 20; z < c.Pos.Z+20; z++ {
        for y := c.Pos.Y - 5; y < c.Pos.Y+5; y++ {
            for x := c.Pos.X - 20; x < c.Pos.X+20; x++ {
                status := c.Wd.GetBlockStatus(int(x), int(y), int(z))
                if status > treeStart && status < treeEnd {
                    woodBlocks = append(woodBlocks, [3]float64{x, y, z})
                }
            }
        }
    }
    closest := getClosestBlock(c.Pos.X, c.Pos.Y, c.Pos.Z, woodBlocks)
    // write function that will walk to the correct location
    destination = closest
    fmt.Println(closest)
}


