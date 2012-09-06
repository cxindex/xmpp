//primitive example of conference-bot
package main

import (
	"github.com/cxindex/xmpp"
	"fmt"
	"log"
)

func main() {
	//server:port, user, domain, password, resource, config
	Conn, err := xmpp.Dial("xmpp.ru:5222", "hypnotoad", "xmpp.ru", "123", "necrotoad", nil)
	if err != nil {
		log.Fatal(err)
	}

	//pop up in roster
	if err := Conn.SignalPresence(""); err != nil { 
		log.Fatal(err)
	}

	if err := Conn.SendPresenceRoom("ttyh@conference.jabber.ru/GNewToad", ""); err != nil {
		log.Fatal(err)
	}

	for {
		next, err := Conn.Next()
		if err != nil {
			log.Fatal(err)
		}

		switch got := next.Value.(type) {
		case *xmpp.ClientMessage:
			fmt.Println(got)
		case *xmpp.ClientPresence:
			//do something
		default:
			//do something else
		}
	}
	// fmt.Println(Conn.SendPresence("ttyh@conference.jabber.ru/GNewToad", "unavailable", "")) //exit from
}
