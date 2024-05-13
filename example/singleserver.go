package main

import (
	"context"
	"log"
	"time"

	"github.com/setavenger/go-electrum/electrum"
)

func main() {
	client, err := electrum.NewClientTCP(context.Background(), "electrum.example.com:50001", "127.0.0.1:9050")

	if err != nil {
		log.Fatal(err)
	}

	serverVer, protocolVer, err := client.ServerVersion(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server version: %s [Protocol %s]", serverVer, protocolVer)

	go func() {
		for {
			if err := client.Ping(context.Background()); err != nil {
				log.Fatal(err)
			}
			time.Sleep(60 * time.Second)
		}
	}()
}
