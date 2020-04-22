package producermq

import (
	"apiserver/mq"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type BindingConfig struct {
	RouteKey string `mapstructure:"routeKey"`
	Queue    string `mapstructure:"queue"`
}

var p *mq.Producer

func Init() {
	m, err := mq.New(viper.GetString("rabbitMQUrl")).Open()
	if err != nil {
		log.Error("[ERROR]", err)
		return
	}

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

	p, err = m.Producer("apiserver-producer")
	if err != nil {
		log.Error("[ERROR] Create producer failed", err)
		return
	}
	if err := p.SetExchangeBinds(exb).Confirm(true).Open(); err != nil {
		log.Error("[ERROR] Open failed", err)
		return
	}
}

func GetProducer() *mq.Producer {
	return p
}
