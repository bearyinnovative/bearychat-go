// Chaos Monkey will talk to you randomly 🙊
//
// Envvars:
//
//      - `CM_RTM_TOKEN`: BearyChat RTM token
//      - `CM_VICTIMS`: user ids who will be talk with,
//                      separates with comma: `=bw52O,=bw52P`
package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/bearyinnovative/bearychat-go"
)

type Config struct {
	rtmToken string
	victims  []string
}

func (c Config) isVictimUID(uid string) bool {
	for _, victimUID := range c.victims {
		if victimUID == uid {
			return true
		}
	}
	return false
}

func (c Config) randomVictim() string {
	return c.victims[rand.Intn(len(c.victims))]
}

func (c Config) insultMessage(user *bearychat.User) bearychat.RTMMessage {
	messages := []string{
		fmt.Sprintf("hi %s", user.Name),
		fmt.Sprintf("hey %s, are you chill?", user.Name),
		fmt.Sprintf("苟利国家生死以，岂因祸福避趋之，%s 识得唔识得啊？", user.Name),
	}

	return bearychat.RTMMessage{
		"type":        bearychat.RTMMessageTypeP2PMessage,
		"vchannel_id": user.VChannelId,
		"to_uid":      user.Id,
		"text":        messages[rand.Intn(len(messages))],
		"call_id":     rand.Int(), // FIXME `call_id` sequence generator
		"refer_key":   nil,
	}
}

func main() {
	config := mustLoadConfigFromEnv()

	rtmClient, err := bearychat.NewRTMClient(config.rtmToken)
	checkErr(err)

	user, wsHost, err := rtmClient.Start()
	checkErr(err)

	rtmLoop, err := bearychat.NewRTMLoop(wsHost)
	checkErr(err)

	checkErr(rtmLoop.Start())
	defer rtmLoop.Stop()

	go rtmLoop.Keepalive(time.NewTicker(2 * time.Second))

	errC := rtmLoop.ErrC()
	messageC, err := rtmLoop.ReadC()
	checkErr(err)

	tickTock := time.NewTicker(15 * time.Second)
	defer tickTock.Stop()

	for {
		select {
		case err := <-errC:
			checkErr(err)
			return
		case message := <-messageC:
			if !message.IsChatMessage() {
				continue
			}
			if !message.IsP2P() {
				continue
			}
			if message.IsFromMe(*user) {
				continue
			}
			uid := message["uid"].(string)
			if !config.isVictimUID(uid) {
				continue
			}

			log.Printf("user %s said: %s", uid, message["text"])

			checkErr(rtmLoop.Send(message.Refer("🙊")))
		case <-tickTock.C:
			user, err := rtmClient.User.Info(config.randomVictim())
			checkErr(err)

			log.Printf("insulting user %s", user.Name)
			checkErr(rtmLoop.Send(config.insultMessage(user)))
		}
	}
}

func mustLoadConfigFromEnv() (config Config) {
	config.rtmToken = os.Getenv("CM_RTM_TOKEN")
	if config.rtmToken == "" {
		log.Fatalf("`CM_RTM_TOKEN` is required!")
		return
	}

	svictims := os.Getenv("CM_VICTIMS")
	if svictims == "" {
		log.Fatalf("`CM_VICTIMS` is required!")
		return
	}

	config.victims = strings.Split(svictims, ",")
	if len(config.victims) == 0 {
		log.Fatalf("`CM_VICTIMS` is required!")
		return
	}

	return
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
