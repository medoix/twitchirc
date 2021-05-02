package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"github.com/gempir/go-twitch-irc/v2"
)

func main() {

	// Get the streamer name argument to set as the channel
	if len(os.Args) < 2 {
		fmt.Println("Usage: twitchirc <streamer-name>")
		os.Exit(1)
	}
	channel := "#" + strings.ToLower(os.Args[1])

	// Create the client as either an anonymous or authenticated
	client := twitch.NewAnonymousClient()
	//client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		time := time.Now()
		fmt.Printf("%v:%v [%v] | %v", time.Hour(), time.Minute(), message.User.DisplayName, message.Message)
		fmt.Println()
		//fmt.Println("PrivateMessage:", message.Message)
	})

        client.OnWhisperMessage(func(message twitch.WhisperMessage) {
                fmt.Println("WhisperMessage:", message.Message)
        })

        client.OnClearChatMessage(func(message twitch.ClearChatMessage) {
		fmt.Printf("[%v] %v", message.TargetUsername, message.Message)
        })

        client.OnClearMessage(func(message twitch.ClearMessage) {
                fmt.Printf("[%v] %v", message.Login, message.Message)
        })

	client.OnConnect(func() {
		fmt.Println("Connected to the Twitch IRC Server and Joined", channel)
	})

	client.Join(channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
