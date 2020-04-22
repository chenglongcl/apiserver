package main

import (
	"apiserver/config"
	"apiserver/mq"
	v "apiserver/pkg/version"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info")
)

type BindingConfig struct {
	RouteKey string `mapstructure:"routeKey"`
	Queue    string `mapstructure:"queue"`
}

func main() {
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
	//create rabbitMQ consumer
	m, err := mq.New(viper.GetString("rabbitMQUrl")).Open()
	if err != nil {
		log.Error("[ERROR]", err)
	}
	defer m.Close()
	c, err := m.Consumer("apiserver-consume")
	if err != nil {
		log.Error("[ERROR] Create consumer failed", err)
		return
	}
	defer c.Close()

	defaultExChange := viper.GetString("rabbitMQDefaultExchange")
	var bindingsConf []BindingConfig
	viper.UnmarshalKey("rabbitMQBindings", &bindingsConf)
	bindings := make([]*mq.Binding, len(bindingsConf))
	for k, b := range bindingsConf {
		bindings[k] = &mq.Binding{
			RouteKey: b.RouteKey,
			Queues: []*mq.Queue{
				mq.DefaultQueue(b.Queue),
			},
		}
	}
	exb := []*mq.ExchangeBinds{
		{
			Exch:     mq.DefaultExchange(defaultExChange, mq.ExchangeDirect),
			Bindings: bindings,
		},
	}

	msgC := make(chan mq.Delivery, 1)
	defer close(msgC)
	c.SetExchangeBinds(exb)
	c.SetMsgCallback(msgC)
	c.SetQos(10)
	if err = c.Open(); err != nil {
		log.Error("[ERROR] Open failed", err)
		return
	}
	for msg := range msgC {
		log.Infof("Tag(%d) Body: %s\n", msg.DeliveryTag, string(msg.Body))
		msg.Ack(true)
		time.Sleep(time.Millisecond * 10)
	}
}
