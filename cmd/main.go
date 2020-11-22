package main

import (
	"flag"
	"fmt"
	"os"
	"renuka.io/dialog-wifi-signal-logger/pkg/wifi/signal"
)

func main() {
	ip := flag.String("ip", "", "Router local IP")
	sId := flag.String("sid", "", "Session ID from cookies")
	flag.Parse()

	client := signal.Client{
		IP:        *ip,
		SessionId: *sId,
	}

	strength, err := client.GetStrength()
	if err != nil {
		os.Exit(1)
	}

	fmt.Println(strength)
}
