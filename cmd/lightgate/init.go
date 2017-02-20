package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/Akagi201/utilgo/conflag"
	log "github.com/Sirupsen/logrus"
	"github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/jessevdk/go-flags"
)

// all global shared variables
var (
	logger *log.Logger
	opts   struct {
		Conf            string `long:"conf" default:"" description:"config file"`
		AppName         string `long:"app_name" default:"lightgate" description:"application name"`
		HTTPListenAddr  string `long:"http_listen" default:"0.0.0.0:8000" description:"HTTP listen address"`
		WsListenAddr    string `long:"ws_listen" default:"0.0.0.0:8001" description:"WebSocket listen address"`
		AdminListenAddr string `long:"admin_http_listen" default:"0.0.0.0:8010" description:"Admin HTTP listen address"`
		LogLevel        string `long:"log_level" default:"info" description:"log level"`
	}
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func init() {
	parser := flags.NewParser(&opts, flags.Default)

	_, err := parser.Parse()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	if opts.Conf != "" {
		conflag.LongHyphen = true
		conflag.BoolValue = false
		args, err := conflag.ArgsFrom(opts.Conf)
		if err != nil {
			fmt.Printf("Parse configs from file failed, err: %v", err)
			os.Exit(-1)
		}
		fmt.Printf("args: %v", args)
		_, err = parser.ParseArgs(args)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(-1)
		}
	}

	//_, err = parser.Parse()
	//if err != nil {
	//	fmt.Printf("%v", err)
	//	os.Exit(-1)
	//}
}

func init() {
	logger = log.New()

	if level, err := log.ParseLevel(strings.ToLower(opts.LogLevel)); err != nil {
		logger.Level = level
	}

	log.SetFormatter(&logrus_logstash.LogstashFormatter{
		Type:            opts.AppName,
		TimestampFormat: time.RFC3339Nano,
	})

	logger.Formatter = &logrus_logstash.LogstashFormatter{
		Type:            opts.AppName,
		TimestampFormat: time.RFC3339Nano,
	}
}
