package config

import (
	"flag"
	"os"
)

const (
	EXITNORMAL = 0
	EXITNONROOT = 1
	EXITPORT = 80
	EXITLOG = 3
)

var LogFormat = "text"
var LogLevel = "info"
var ListeningPort = 8088

func init() {
}

func ParseCommandLine() {
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flag.StringVar(&LogFormat, "logformat", LogFormat, "log format = {text, json}")
	flag.StringVar(&LogLevel, "loglevel", LogLevel, "log level = {info, warning, fatal, error, panic, debug}")
	flag.IntVar(&ListeningPort, "port", ListeningPort, "port for RESTful API serving")
	flag.Parse()
}
