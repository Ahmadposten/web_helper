# Web helper

## How it works
The files it should deal with are log files which means that these files can be very large and might not fit in memory.
To tackle this, file is read in a buffered manner and is also outputed in chuncks either to stdout(default) or to a file optionally specified.

## Usage

### Basic usage

For the basic usage you will need to run
`./web_helper --ip <someip>`

It will assume that the file is https://s3.amazonaws.com/syseng-challenge/public_access.log.txt
will download it (chunked) and for each chunk will check if it satisifes the ip.

In case of using this command you will get the answer on stdout


### Another remote file

Run `./web_helper --ip <someip> --remote-file <urlToRemoteFile>``

### Local file
Run `./web_helper --ip <someip> --local-file> <urltoLocalFile>`

It will tell you what is wrong if you specify a bad file

### Installation
I've included a binary `web_builder_osx` which is build for `osx` and another `web_builder_linux` for linux
use the one compatible with your box.

#### Custom build
1. You need to have golang (this is tested on 1.9 but will possible work with other versions)
1. You need to run `go get github.com/ahmadposten/web_helper`
1. run `go build -o web_helper *.go`
1. You will have a binary named web_helper which can be used normally

### Testing
run `go test -v`
