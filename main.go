package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dup2X/crsdk/api/server"
	"github.com/dup2X/gopkg/config"
	"github.com/dup2X/gopkg/logger"
)

var confPath string

func main() {
	flag.StringVar(&confPath, "config", "./configs/main.conf", "-config=./configs/main.conf")
	flag.Parse()
	cfg, err := config.New(confPath)
	if err != nil {
		fmt.Printf("load config failed, err:%v\n", err)
		os.Exit(-2)
	}
	err = logger.NewLoggerWithConfig(cfg)
	if err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		os.Exit(-2)
	}
	if err := server.Serve(cfg); err != nil {
		fmt.Printf("start server failed, err:%v\n", err)
		os.Exit(-2)
	}
}
