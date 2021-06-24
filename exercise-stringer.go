package main

import (
	"fmt"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	sIp := ""
	lenIp := len(ip)
	for idx, ip_part := range ip {
		sIp += fmt.Sprintf("%d", ip_part)
		if idx < lenIp-1 {
			sIp += "."
		}
	}

	return fmt.Sprintf("%v", sIp)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
