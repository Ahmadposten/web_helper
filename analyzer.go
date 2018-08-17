package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

// This contains the needed queries
// Since the file is a log file, it might be too large to fit in memory
// Thus we need to read it and filter on the fly keeping

func isCidrNotation(ip string) bool {
	portions := strings.Split(ip, "/")
	return len(portions) == 2 && portions[1] != ""
}

func toCidr(ip string) string {
	// Here it is a single ip address
	return fmt.Sprintf("%s/32", ip)
}

func filterIps(ip string, source io.Reader, dest io.Writer) error {
	scanner := bufio.NewScanner(source)
	buf := make([]byte, 0, 1024*1024)
	scanner.Buffer(buf, 1024*1024)

	if !isCidrNotation(ip) {
		ip = toCidr(ip)
	}
	_, network, err := net.ParseCIDR(ip)

	if err != nil {
		return err
	}

	for scanner.Scan() {
		stringRecord := scanner.Text()
		rec, err := parseRecord(stringRecord)

		if err != nil {
			return err
		}
		if network.Contains(rec.Ip) {
			_, err := io.WriteString(dest, fmt.Sprintf("%s\n", stringRecord))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
