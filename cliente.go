package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {

	var message = flag.String("me", "", "Mensage que se envia")
	flag.Parse()

	conn, err := net.Dial("tcp", "109.116.135.41:9000")

	if err != nil {
		log.Fatalf("Error try to connect: %v", err)

	}

	defer conn.Close()

	conn.Write([]byte(*message + "\n"))

	res, _ := bufio.NewReader(conn).ReadString('\n')

	fmt.Println(res)
}
