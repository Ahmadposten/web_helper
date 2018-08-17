package main

import "net"

type Record struct {
	Ip net.IP `json:"ip"`
	Log string `json:"log"`
}
