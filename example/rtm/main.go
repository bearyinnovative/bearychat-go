// Demo bot built with BearyChat RTM
//
//      ./rtm -rtmToken <your-BearyChat-RTM-token-here>
package main

import (
	"flag"
	"log"

	bearychat "github.com/bearyinnovative/bearychat-go"
)

var rtmToken string

func init() {
	flag.StringVar(&rtmToken, "rtmToken", "", "BearyChat RTM token")
}

func main() {
	flag.Parse()

	context, err := bearychat.NewRTMContext(rtmToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	err, messageC, errC := context.Run()
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		select {
		case err := <-errC:
			log.Printf("rtm loop error: %+v", err)
			if err := context.Loop.Stop(); err != nil {
				log.Fatal(err)
			}
			return
		case message := <-messageC:
			if !message.IsChatMessage() {
				continue
			}

			// from self
			if message.IsFromUID(context.UID()) {
				continue
			}

			log.Printf(
				"received: %s from %s",
				message["text"],
				message["uid"],
			)

			// only reply mentioned myself
			if mentioned, content := message.ParseMentionUID(context.UID()); mentioned {
				if err := context.Loop.Send(message.Refer(content)); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}
