// package zerotierAPI
package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	zt := newClient("https://api.zerotier.com/api/v1", "m8urPJlwpysavmTbKwSQaw2eVNjEQufn")

	net, _ := zt.createNetwork()
	fmt.Println(net.Config.Name)
}
