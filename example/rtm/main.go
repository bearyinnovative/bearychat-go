// Demo bot built with BearyChat RTM
//
//      ./rtm -rtmToken <your-BearyChat-RTM-token-here>
package main

import (
	"flag"
	"log"
	"time"

	"github.com/bcho/bearychat.go"
)

const (
	RTM_API_BASE = "https://rtm.bearychat.com"
)

var rtmToken string

func main() {
	flag.Parse()

	rtmClient, err := bearychat.NewRTMClient(
		rtmToken,
		bearychat.WithRTMAPIBase(RTM_API_BASE),
	)
	if err != nil {
		log.Fatal(err)
	}

	user, wsHost, err := rtmClient.Start()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("rtm connected as %s", user.Name)

	rtmLoop, err := bearychat.NewRTMLoop(wsHost)
	if err != nil {
		log.Fatal(err)
	}

	if err := rtmLoop.Start(); err != nil {
		log.Fatal(err)
	}
	defer rtmLoop.Stop()

	go rtmLoop.Keepalive(time.NewTicker(2 * time.Second))

	errC := rtmLoop.ErrC()
	messageC, err := rtmLoop.ReadC()
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-errC:
			log.Printf("rtm loop error: %+v", err)
			if err := rtmLoop.Stop(); err != nil {
				log.Fatal(err)
			}
			return
		case message := <-messageC:
			if message.Type() != bearychat.RTMMessageTypeP2PMessage &&
				message.Type() != bearychat.RTMMessageTypeChannelMessage {
				continue
			}
			log.Printf(
				"rtm loop received: %s from %s",
				message["text"],
				message["uid"],
			)

			if message["uid"] == user.Id {
				continue
			}

			reply := bearychat.RTMMessage{
				"type":        message.Type(),
				"text":        "Pardon?",
				"refer_key":   message["key"],
				"vchannel_id": message["vchannel_id"],
			}

			if message.Type() == bearychat.RTMMessageTypeChannelMessage {
				reply["channel_id"] = message["channel_id"]
			}

			if message.Type() == bearychat.RTMMessageTypeP2PMessage {
				reply["to_uid"] = message["uid"]
			}

			if err := rtmLoop.Send(reply); err != nil {
				log.Fatal(err)
			}
		}
	}
}

func init() {
	flag.StringVar(&rtmToken, "rtmToken", "", "BearyChat RTM token")
}
