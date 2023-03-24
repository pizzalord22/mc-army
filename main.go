package main

import (
    "context"
    "fmt"
    "github.com/Tnze/go-mc/net"
    "github.com/mattn/go-colorable"
    GMMAuth "github.com/maxsupermanhd/go-mc-ms-auth"
    "log"
    "time"

    "github.com/Tnze/go-mc/bot"
    "github.com/Tnze/go-mc/bot/basic"
    "github.com/Tnze/go-mc/bot/msg"
    "github.com/Tnze/go-mc/bot/playerlist"
    "github.com/Tnze/go-mc/chat"
    _ "github.com/Tnze/go-mc/data/lang/en-us"
    "github.com/Tnze/go-mc/data/packetid"
    pk "github.com/Tnze/go-mc/net/packet"
)

const timeout = 45

var (
    c *bot.Client
    p *basic.Player

    playerList  *playerlist.PlayerList
    chatHandler *msg.Manager

    watch chan time.Time
)

func main() {
    mauth, err := GMMAuth.GetMCcredentials("./auth.cache", "88650e7e-efee-4857-b9a9-cf580a00ef43")
    if err != nil {
        log.Print(err)
        return
    }
    //log.Print("Authenticated as ", mauth.Name, " (", mauth.UUID, ")")
    c = bot.NewClient()
    c.Auth = mauth
    log.SetOutput(colorable.NewColorableStdout()) // optional for colorable output

    p = basic.NewPlayer(c, basic.DefaultSettings, basic.EventsListener{
        GameStart:  onGameStart,
        Disconnect: onDisconnect,
        Death:      onDeath,
    })

    playerList = playerlist.New(c)
    chatHandler = msg.New(c, p, playerList, msg.EventsHandler{
        SystemChat:        onSystemChat,
        PlayerChatMessage: onPlayerChat,
        DisguisedChat:     onDisguisedChat,
    })

    // Register event handlers

    //c.Events.AddListener(soundListener)

    // Login
    //err := c.JoinServer("193.31.31.162:25513")
    for {
        playerList = playerlist.New(c)
        for _, player := range playerList.PlayerInfos {
            fmt.Println(player.Name)
        }
        err = c.JoinServerWithOptions("193.31.31.162:25513", bot.JoinOptions{
            MCDialer:    &net.DefaultDialer,
            Context:     context.Background(),
            NoPublicKey: false,
            KeyPair:     nil,
        })
        if err != nil {
            log.Fatal(err)
        }
        log.Println("Login success")

        // JoinGame
        err = c.HandleGame()
        if err != nil {
            log.Println(err)
        }
    }
}

func onDeath() error {
    log.Println("Died and Respawned")
    // If we exclude Respawn(...) then the player won't press the "Respawn" button upon death
    return p.Respawn()
}

func onGameStart() error {
    log.Println("Game start")

    //watch = make(chan time.Time)
    //go watchDog()

    //return UseItem(0)
    return nil
}

var soundListener = bot.PacketHandler{
    ID:       packetid.ClientboundSound,
    Priority: 0,
    F: func(p pk.Packet) error {
        var (
            SoundID       pk.VarInt
            SoundCategory pk.VarInt
            X, Y, Z       pk.Int
            Volume, Pitch pk.Float
        )
        if err := p.Scan(&SoundID, &SoundCategory, &X, &Y, &Z, &Volume, &Pitch); err != nil {
            return err
        }
        return onSound(int(SoundID))
    },
}

func UseItem(hand int32) error {
    return c.Conn.WritePacket(pk.Marshal(
        packetid.ServerboundUseItem,
        pk.VarInt(hand),
    ))
}

//goland:noinspection SpellCheckingInspection
func onSound(id int) error {
    if id == 369 {
        if err := UseItem(0); err != nil { // retrieve
            return err
        }
        log.Println("gra~")
        time.Sleep(time.Millisecond * 300)
        if err := UseItem(0); err != nil { // throw
            return err
        }
        watch <- time.Now()
    }
    return nil
}

func onSystemChat(c chat.Message, overlay bool) error {
    log.Printf("System Chat: %v, Overlay: %v", c, overlay)
    return nil
}

func onPlayerChat(c chat.Message, b bool) error {
    log.Println("Player Chat:", c, b)
    return nil
}

func onDisguisedChat(c chat.Message) error {
    log.Println("Disguised Chat:", c)
    return nil
}

func onDisconnect(c chat.Message) error {
    log.Println("Disconnect:", c)
    return nil
}

func watchDog() {
    to := time.NewTimer(time.Second * timeout)
    for {
        select {
        case <-watch:
        case <-to.C:
            log.Println("rethrow")
            if err := UseItem(0); err != nil {
                panic(err)
            }
        }
        to.Reset(time.Second * timeout)
    }
}
