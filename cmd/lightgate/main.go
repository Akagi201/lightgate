package main

import (
	"net/http"
	"runtime"

	"github.com/Akagi201/utilgo/conflag"
	log "github.com/Sirupsen/logrus"
	flags "github.com/jessevdk/go-flags"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

var opts struct {
	Conf            string `long:"conf" default:"" description:"config file"`
	HTTPListenAddr  string `long:"http_listen" default:"0.0.0.0:8000" description:"HTTP listen address"`
	WsListenAddr    string `long:"ws_listen" default:"0.0.0.0:8001" description:"WebSocket listen address"`
	AdminListenAddr string `long:"admin_http_listen" default:"0.0.0.0:8010" description:"Admin HTTP listen address"`
	LogLevel        string `long:"log_level" default:"info" description:"log level"`
}

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-0¬2 15:04:05",
	})
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	parser := flags.NewParser(&opts, flags.Default)

	parser.Parse()
	log.Printf("opts: %+v", opts)

	if opts.Conf != "" {
		conflag.LongHyphen = true
		conflag.BoolValue = false
		log.Printf("opts.Conf: %v", opts.Conf)
		args, err := conflag.ArgsFrom(opts.Conf)
		if err != nil {
			log.Fatalf("Parse configs from file failed, err: %v", err)
		}
		log.Printf("args: %v", args)
		parser.ParseArgs(args)
	}

	parser.Parse()

	log.Printf("opts: %+v", opts)

	if level, err := log.ParseLevel(opts.LogLevel); err != nil {
		log.SetLevel(level)
	}

	// reversed proxy logic
	//log.Fatalln(run())
}

func run() error {
	fwd, err := forward.New()
	if err != nil {
		return err
	}

	root := http.NewServeMux()

	root.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL = testutils.ParseURI("http://localhost:8088")
		fwd.ServeHTTP(w, r)
	})

	http.ListenAndServe(":8080", root)

	return nil
}
