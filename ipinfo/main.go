package main

import (
	"fmt"
	"log"
	"net"

	"github.com/ipinfo/go-ipinfo/ipinfo"
)

func main() {
	//Get Access token by signing up a free account at ipinfo.io/signup
	authTransport := ipinfo.AuthTransport{Token: ""}
	httpClient := authTransport.Client()
	client := ipinfo.NewClient(httpClient)
	info, err := client.GetInfo(net.ParseIP("8.8.8.8"))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}
