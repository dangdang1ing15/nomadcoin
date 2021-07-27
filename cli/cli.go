package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/dangdang1ing15/nomadcoin/explorer"
	"github.com/dangdang1ing15/nomadcoin/rest"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port:		Set the PORT of the server\n")
	fmt.Printf("-mode:		Choose between 'html' and 'rest'\n\n")
	runtime.Goexit()
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	hport := flag.Int("hport", 3000, "Set port of the html server")
	rport := flag.Int("rport", 4000, "Set port of the rest server")
	mode := flag.String("mode", "run", "Choose between 'html' and 'rest' or 'run' for both")

	flag.Parse()

	switch *mode {
	case "run":
		go rest.Start(*rport)
		explorer.Start(*hport)
	case "rest":
		rest.Start(*rport)
	case "html":
		explorer.Start(*hport)
	default:
		usage()
	}
}
