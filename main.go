package main

import (
	"flag"
	"net/http"
	"io"
	"os"
	"log"
)

const (
	DEFAULT_REMOTE_FILE = "https://s3.amazonaws.com/syseng-challenge/public_access.log.txt"
)

func getRemoteFile(uri string) (io.Reader, error){
	resp, err := http.Get(uri)
	if(err != nil){
		return nil, err
	}

	return resp.Body, nil
}

func getLocalFile(uri string) (io.Reader, error){
	reader, err := os.Open(uri)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func getFile(uri string, from string) (io.Reader, error){
	if(from == "remote"){
		return getRemoteFile(uri)
	}
	return getLocalFile(uri)
}

func main(){
	ip := flag.String("ip", "", "Ip to fetch logs for")
	remoteFile := flag.String("remote-file", DEFAULT_REMOTE_FILE, "Remote log file")

	localFile := flag.String("local-file", "", "Local log file")
	destination := flag.String("d", "", "optional file to write to")

	flag.Parse()

	log.Printf("ip is %s ", *ip)
	if len(*ip) == 0 {
		flag.Usage()
		panic("Ip needs to be specified")
	}
	var from string
	var uri string

	if(len(*localFile) > 0 ){
		from = "local"
		uri  = *localFile
	}else{
		from = "remote"
		uri = *remoteFile
	}
	reader, err := getFile(uri, from)

	if err != nil {
		panic(err)
	}

	writer := os.Stdout // default is stdout!

	if len(*destination) > 0 {
		writer, err = os.Create(*destination)
		if err != nil {
			log.Fatalf("Could not create output file ! %v", err)
			return
		}
		defer writer.Close()
	}

	if err := filterIps(*ip, reader, writer); err != nil {
		panic(err)
	}


	if err != nil {
		panic(err)
	}
}
