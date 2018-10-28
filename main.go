package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/kyawmyintthein/runtime-log-level-demo/config"
	"github.com/sirupsen/logrus"

	//	_ "github.com/kyawmyintthein/runtime-log-level-demo/internal/etcdconfig"

	"github.com/kyawmyintthein/runtime-log-level-demo/internal/etcdconfig"
	"github.com/kyawmyintthein/runtime-log-level-demo/internal/logging"
	"github.com/kyawmyintthein/runtime-log-level-demo/router"
)

var (
	dialTimeout    = 2 * time.Second
	requestTimeout = 10 * time.Second
)

func main() {
	var configFilePath string
	var serverPort string
	flag.StringVar(&configFilePath, "config", "config.yml", "absolute path to the configuration file")
	flag.StringVar(&serverPort, "server_port", "3000", "port on which server runs")
	flag.Parse()

	generalConfig := config.LoadConfig(configFilePath)
	etcdconfig.Init(generalConfig)
	logging.Init(generalConfig.LogLevel)

	r := router.GetRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("This is Info log.")
		logrus.Warnf("This is Warn log.")
		logrus.Errorf("This is Error log")
		logrus.Debugf("This is Debug log")
	})

	logrus.Infoln("############################## Server Started ##############################")
	http.ListenAndServe(":"+serverPort, r)
}
