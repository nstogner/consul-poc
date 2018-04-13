package main

import (
	"fmt"
	"net"
)

func main() {
	infs, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for i, inf := range infs {
		fmt.Printf("[%d] Name: %s\n", i, inf.Name)
		fmt.Printf("[%d] HardwareAddr: %s\n", i, inf.HardwareAddr)

		addrs, err := inf.Addrs()
		if err != nil {
			panic(err)
		}
		for _, a := range addrs {
			fmt.Printf("[%d] Addr: %s\n", i, a.String())
		}
	}
}
