package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

import (
	"./gtclient"
	hs "./simplehttpserver"
)

var version string

var (
	port         = flag.String("p", "", "port")
	subdomain    = flag.String("sub", "", "request subdomain to serve on")
	remote       = flag.String("r", "localtunnel.net:34000", "the remote gotunnel server host/ip:port")
	fileServer   = flag.Bool("fs", false, "Server files in the current directory. Use -p to specify the port.")
	serveDir     = flag.String("d", "", "The directory to serve. To be used with -fs.")
	showVersion  = flag.Bool("v", false, "Show version and exit")
)

// var version string

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "\nOptions:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *showVersion {
		fmt.Println("Version - ", version)
		return
	}

	if *port == "" || *remote == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *fileServer {
		dir := ""
		// Simple file server.
		if *port == "" {
			fmt.Fprintf(os.Stderr, "-fs needs -p (port) option")
			flag.Usage()
			os.Exit(1)
		}
		if *serveDir == "" {
			dir, _ = os.Getwd()
		} else {
			if path.IsAbs(*serveDir) {
				dir = path.Clean(*serveDir)
			} else {
				wd, _ := os.Getwd()
				dir = path.Clean(path.Join(wd, *serveDir))
			}
		}
		go hs.NewSimpleHTTPServer(*port, dir)
	}

	servInfo := make(chan string)

	go func() {
		serverat := <-servInfo
		fmt.Printf("Your site should be available at: %s\n", serverat)
	}()

	if !gtclient.SetupClient(*port, *remote, *subdomain, servInfo) {
		flag.Usage()
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
