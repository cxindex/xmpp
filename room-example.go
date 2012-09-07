//primitive example of conference-bot
package main

import (
	"fmt"
	"github.com/cxindex/xmpp"
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
			Conn.SendRoom("ttyh@conference.jabber.ru", "Oh! I can smell a cave troll.")      //to chat
			Conn.Send("ttyh@conference.jabber.ru/cx", "Oh! I can smell a cave troll.")       //to occupant
			Conn.Send("cx.index@gmail.com", "Do not take me as a conjuror of cheap tricks!") //to jid
		default:
			//do something else
		}
	}
	// fmt.Println(Conn.SendPresence("ttyh@conference.jabber.ru/GNewToad", "unavailable", "")) //exit from
}
