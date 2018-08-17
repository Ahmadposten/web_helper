package main
import (
	"strings"
	"fmt"
	"net"
)

func parseRecord(rawRecord string) (Record, error) {
	splitString := strings.SplitN(rawRecord, " ", 2)
	rec := Record{}
	if(len(splitString) < 2){
		// This is a malformed string

		err := fmt.Errorf("log %s is a malformed string! ", rawRecord)
		return rec, err
	}

	rec.Ip = net.ParseIP(splitString[0])
	rec.Log = splitString[1]
	return rec, nil
}
