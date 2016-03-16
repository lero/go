package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"strings"
)

var xinga = map[string]string{
	"0":  "chupa essa manga",
	"1":  "ja sabe neh?",
	"2":  "cabeca de nos todos",
	"3":  "burrao",
	"5":  "maragogipe",
	"6":  "vc usa mac os",
	"7":  "macfag",
	"8":  "chori",
	"9":  "ah falo",
	"10": "fodasse",
}

func random() string {
	r, _ := rand.Int(rand.Reader, big.NewInt(10))
	x := r.String()
	return xinga[x]
}

func scanner(conn net.Conn) {
	for {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			bla := scanner.Text()
			if strings.Contains(bla, "!xinga") {
				fmt.Fprintf(conn, "PRIVMSG #trutas :"+random()+"\n")
			}
			if strings.HasPrefix(bla, "PING :") {
				fmt.Fprintf(conn, "PONG"+bla[4:]+"\n")
			}
			fmt.Println(bla)
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

func connect() {
	conn, err := net.Dial("tcp", "irc.freenode.net:6667")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Fprintf(conn, "NICK golero\n")
	fmt.Fprintf(conn, "USER golero 8 * : golero\n")
	fmt.Fprintf(conn, "PRIVMSG NICKSERV :IDENTIFY xxxxx\n")
	fmt.Fprintf(conn, "JOIN #trutas\n")
	scanner(conn)
}

func main() {
	connect()
}
