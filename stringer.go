// +build ignore

package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	var ipStrs []string
	for _, v := range ip {
		ipStrs = append(ipStrs, strconv.FormatUint(uint64(v), 10))
	}
	return strings.Join(ipStrs, ".")
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
