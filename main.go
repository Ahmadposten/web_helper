package main

import (
	"flag"
	"io"
	"fmt"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_REMOTE_FILE = "https://s3.amazonaws.com/syseng-challenge/public_access.log.txt"

)
var ErrorStream *log.Logger

func init(){
	ErrorStream = log.New(os.Stderr, "", 0)
}

func logError(errMessage string){
	ErrorStream.Printf("Error: %s ", errMessage)
}

func getRemoteFile(uri string) (io.Reader, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func getLocalFile(uri string) (io.Reader, error) {
	reader, err := os.Open(uri)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func getFile(uri string, from string) (io.Reader, error) {
	if from == "remote" {
		return getRemoteFile(uri)
	}
	return getLocalFile(uri)
}

func main() {
	ip := flag.String("ip", "", "Ip to fetch logs for")
	remoteFile := flag.String("remote-file", DEFAULT_REMOTE_FILE, "Remote log file")

	localFile := flag.String("local-file", "", "Local log file")
	destination := flag.String("d", "", "optional file to write to")

	flag.Parse()

	if len(*ip) == 0 {
		logError("Ip needs to be specified")
		flag.Usage()
		os.Exit(1)
	}
	var from string
	var uri string

	if len(*localFile) > 0 {
		from = "local"
		uri = *localFile
	} else {
		from = "remote"
		uri = *remoteFile
	}
	reader, err := getFile(uri, from)

	if err != nil {
		logError(fmt.Sprintf("Could not fetch file: %s ", err.Error()))
		os.Exit(1)
	}

	writer := os.Stdout // default is stdout!

	if len(*destination) > 0 {
		writer, err = os.Create(*destination)
		if err != nil {
			logError(fmt.Sprintf("Could not create output file ! %s", err.Error()))
			os.Exit(1)
		}
		defer writer.Close()
	}

	if err := filterIps(*ip, reader, writer); err != nil {
		logError(fmt.Sprintf("Error running your query : %s", err.Error()))
		os.Exit(1)
	}
}
