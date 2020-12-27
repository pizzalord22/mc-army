package main

import (
	"github.com/Tnze/go-mc/bot/path"
	"github.com/mattn/go-colorable"
	"log"
	"math"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	_ "github.com/Tnze/go-mc/data/lang/en-us"
)

//const timeout = 45

// wood types
const (
	treeStart = 109
	treeEnd   = 123
)

// const block id
const (
	air   = 0
	water = 34
)

var (
	c           *bot.Client
	watch       chan time.Time
	destination bot.Position
)

// todo make a struct that contains bot data
// make the different actions queryable with a channel
func main() {
	w := ParseWorld(`
.....~......
.....MM.....
.F........T.
....MMM.....
............`,
	)
	Astar(w)
	return
	log.SetOutput(colorable.NewColorableStdout())
	c = bot.NewClient()
	c.Name = "bot1"
	//Login
	err := c.JoinServer("localhost", 25565)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//Register event handlers
	c.Events.GameStart = onGameStart
	c.Events.ChatMsg = onChatMsg
	c.Events.Disconnect = onDisconnect
	c.Events.SoundPlay = onSound
	c.Events.Die = onDeath
	c.Events.PositionChange = onMove

	//JoinGame
	err = c.HandleGame()
	if err != nil {
		log.Fatal(err)
	}
}

func onDeath() error {
	log.Println("Died and Respawned")
	time.Sleep(1 * time.Second)
	c.Respawn() // If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
	return nil
}

func onGameStart() error {
	log.Println("Game start")

	watch = make(chan time.Time)
	go watchDog()

	return c.UseItem(0)
}

func onSound(name string, category int, x, y, z float64, volume, pitch float32) error {
	if name == "entity.fishing_bobber.splash" {
		if err := c.UseItem(0); err != nil { //retrieve
			return err
		}
		log.Println("gra~")
		time.Sleep(time.Millisecond * 300)
		if err := c.UseItem(0); err != nil { //throw
			return err
		}
		watch <- time.Now()
	}
	return nil
}

func onChatMsg(m chat.Message, pos byte, uuid uuid.UUID) error {
	msg := m.String()
	log.Println("Chat:", msg)
	if msg == "<prof_pizza_v> run" {
		c.Inputs = path.Inputs{
			Yaw:       0,
			Pitch:     0,
			ThrottleX: 1,
			ThrottleZ: 0,
			Jump:      false,
		}
		log.Println("starting to run")
	}

	if msg == "<prof_pizza_v> stop" {
		c.Inputs = path.Inputs{
			Yaw:       0,
			Pitch:     0,
			ThrottleX: 0,
			ThrottleZ: 0,
			Jump:      false,
		}
		log.Println("stopping all actions")
	}

	if msg == "<prof_pizza_v> tree" {
		findTree()
		chopTree()
	}

	return nil
}

func onDisconnect(c chat.Message) error {
	log.Println("Disconnect:", c)
	return nil
}

func watchDog() {
	//to := time.NewTimer(time.Second * timeout)
	//for {
	//	select {
	//	case <-watch:
	//	case <-to.C:
	//		log.Println("rethrow")
	//		if err := c.UseItem(0); err != nil {
	//			panic(err)
	//		}
	//	}
	//	to.Reset(time.Second * timeout)
	//}
}

func euclidean(x1, y1, z1, x2, y2, z2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2) + math.Pow(z2-z1, 2))
}

// todo should chop down the whole tree
// needs the trees position and height
func chopTree() {}

// get players current position and walk to the given coords
func walk(x, y, z float64) {

}
