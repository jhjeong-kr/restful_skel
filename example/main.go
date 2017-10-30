package main

import (
	"os"

	"restful_skel"
	"restful_skel/config"
)

func init() {
}

func main() {
	config.ParseCommandLine()
	os.Exit(restful_skel.Run())
}
