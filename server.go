package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"log"
	"log/syslog"
)


func main() {

	logWriter, err := os.OpenFile("Server.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)

if err != nil {
    log.Fatalf("error opening file: %v", err)
}
log.SetOutput(logWriter)
defer logWriter.Close()

	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		log.Fatalf("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		log.Fatalf(err.Error())
		return
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		log.Fatalf(err.Error())
		return
	}

	fmt.Println("Server started...")
	log.Println("Server Started...")

	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			log.Fatalln(err)
			return
		}
		if strings.TrimSpace(string(netData)) == "STOP" {
			fmt.Println("Exiting TCP server!")
			log.Println("Exiting TCP server!")
			return
		}

		fmt.Print("-> ", string(netData))
		log.Println("-> ", string(netData))
		t := time.Now()
		myTime := "received message at " + t.Format(time.RFC3339) + "\n"
		c.Write([]byte(myTime))
	}
}
