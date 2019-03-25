package main

import (
	"flag"
	"go-sensitive/internal/pkg/server"
	"os"
	"path/filepath"
)

var port int
var configPath string

func main() {

	flagParse()
	server.Start(port, configPath)
}
func flagParse() {
	defaultConfigPath, _ := os.Getwd()
	defaultConfigPath = filepath.Dir(defaultConfigPath)
	flag.IntVar(&port, "port", 8081, "listen port")
	flag.StringVar(&configPath, "config", defaultConfigPath, "config dir path")
	flag.Parse()
}
