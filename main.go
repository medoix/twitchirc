package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"github.com/gempir/go-twitch-irc/v2"
)

func abort(message string) {
	fmt.Println("Error: ", message)
	os.Exit(1)
}

func main() {
	channel := ""
	cmdArgs := os.Args[1:]
	if len(cmdArgs) > 1 {
		abort("Only specifiy one channel name")
	} else if len(cmdArgs) < 1 {
		abort("Please specify a channel to join")
	} else {
		// Convert []string from arguments to normal string
		channel = strings.Join(os.Args[1:], " ")
	}

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
		fmt.Println("Connected to the Twitch IRC Server and Joined #", channel)
	})

	client.Join(channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
