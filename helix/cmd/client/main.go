package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Helix Client (HTTP mode)")
	fmt.Println("Type: SET key value")
	fmt.Println("      GET key")
	fmt.Println("      DEL key")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		var cmd, key, value string
		fmt.Sscanf(input, "%s %s %s", &cmd, &key, &value)

		switch cmd {
		case "SET":
			http.Get(fmt.Sprintf("http://localhost:8080/set?key=%s&value=%s", key, value))
			fmt.Println("OK")
		case "GET":
			res, _ := http.Get(fmt.Sprintf("http://localhost:8080/get?key=%s", key))
			body, _ := io.ReadAll(res.Body)
			fmt.Println(string(body))
		case "DEL":
			http.Get(fmt.Sprintf("http://localhost:8080/del?key=%s", key))
			fmt.Println("OK")
		default:
			fmt.Println("unknown command")
		}
	}
}
