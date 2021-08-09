package main

import (
    "log"
    "os"
    "time"

    "github.com/grunndev/hello_assert_bot"
    "github.com/lrstanley/girc"
)

func main() {
    hello_assert_bot.Pledge("stdio inet dns proc rpath")

    client := girc.New(girc.Config{
        Server: "irc.dal.net",
        Port:   6697,
        Nick:   "hello_assert_bot",
        User:   "hello_assert_bot",
        Debug:  os.Stdout,
        SSL:    true,
    })

    client.Handlers.Add(girc.CONNECTED, func(c *girc.Client, e girc.Event) {
        c.Cmd.Join("#grunndev")
    })

    client.Handlers.Add(girc.JOIN, func(c *girc.Client, e girc.Event) {
        if e.Source.Name == "assert" {
            c.Cmd.Message(e.Params[0], "Hello assert!")
        }
    })

    for {
        if err := client.Connect(); err != nil {
            log.Printf("error: %s", err)
            log.Println("reconnecting in 30 seconds...")
            time.Sleep(30 * time.Second)
        } else {
            return
        }
    }
}
