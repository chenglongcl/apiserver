package main

import (
	"apiserver/config"
	"apiserver/mgomodel"
	"apiserver/pkg/gormx"
	"apiserver/pkg/json"
	"apiserver/pkg/oss"
	"apiserver/pkg/producermq"
	"apiserver/pkg/redisgo"
	v "apiserver/pkg/version"
	"apiserver/router"
	"apiserver/router/middleware"
	"apiserver/ws"
	"errors"
	"fmt"
	"github.com/chenglongcl/log"
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"github.com/json-iterator/go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info")
)

func main() {
	overseer.Run(overseer.Config{
		Program: program,
		Address: ":8080",
		Fetcher: &fetcher.File{Path: "public/update/apiserver"},
		Debug:   false,
	})
}

func program(state overseer.State) {
	pflag.Parse()
	if *version {
		info := v.Get()
		marshalled, err := jsoniter.MarshalIndent(&info, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalled))
		return
	}
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init mysql db
	gormx.InitMySQL()
	defer gormx.Close()
	// int mongo
	mgomodel.Init()
	defer mgomodel.Close()
	// init redis
	_ = redisgo.Init()
	json.Init()
	//init oss
	oss.Init()
	//init inject
	//inject.Init()
	producermq.Init()
	//Set gin mode
	gin.SetMode(viper.GetString("runmode"))
	//Create the gin engine
	g := gin.New()
	//webSockets
	m := melody.New()
	ws.Init()
	router.Load(
		g,
		m,
		//Middlewares
		middleware.RequestId(),
		middleware.Logging(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	//log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
	log.Infof(http.Serve(state.Listener, g).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
