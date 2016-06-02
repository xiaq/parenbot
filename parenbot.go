package main

import (
	"bytes"
	"container/list"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/xiaq/tg"
)

var logall = flag.Bool("logall", false, "log all runes in all messages")

func main() {
	flag.Parse()

	buf, err := ioutil.ReadFile("token.txt")
	if err != nil {
		log.Fatalln("read token file:", err)
	}
	token := strings.TrimSpace(string(buf))

	bot := tg.NewBot(token)
	bot.OnUpdate(onUpdate)
	bot.Main()
}

func onUpdate(bot *tg.Bot, u *tg.Update) {
	if u.Message == nil || u.Message.Text == nil {
		return
	}
	text := *u.Message.Text
	patch := makePatch(text)
	if patch == "" {
		return
	}

	bot.Get("/sendMessage", tg.Query{
		"chat_id": u.Message.Chat.ID,
		"text":    patch + " ○(￣^￣○)",
	}, nil)
}

type verboseRune rune

func (r verboseRune) String() string {
	return fmt.Sprintf("%s U+%X", string(r), rune(r))
}

func makePatch(s string) string {
	openers := list.New()
	for _, r := range s {
		if _, isOpener := closerOf[r]; isOpener {
			// Push an opener to the list.
			openers.PushBack(r)
			if *logall {
				log.Printf("pushed opener %s", verboseRune(r))
			}
		} else if opener, isCloser := openerOf[r]; isCloser {
			if *logall {
				log.Printf("found closer %s, popping opener %s", verboseRune(r), verboseRune(opener))
			}
			// Remove the last matching opener from the list, if any
			for p := openers.Back(); p != nil; p = p.Prev() {
				if p.Value.(rune) == opener {
					openers.Remove(p)
					break
				}
			}
		} else {
			if *logall {
				log.Printf("other char %s", verboseRune(r))
			}
		}
	}
	var buf bytes.Buffer
	for p := openers.Back(); p != nil; p = p.Prev() {
		buf.WriteRune(closerOf[p.Value.(rune)])
	}
	return buf.String()
}
