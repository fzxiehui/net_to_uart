package mqtt

import (
	"testing"
	"time"

	"github.com/fzxiehui/net_to_uart/config"
)

func TestMQTT(t *testing.T) {

	cfg := config.Config()
	mqtt := NewMqtt(cfg.GetString("mqtt.broker"),
		cfg.GetString("mqtt.clientid"),
		cfg.GetString("mqtt.username"),
		cfg.GetString("mqtt.password"))

	mqtt.Connect()
	mqtt.Subscribe(cfg.GetString("mqtt.sub_topic"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	mqtt.Publish(cfg.GetString("mqtt.pub_topic"), []byte("hello"), 0)
	time.Sleep(13 * time.Second)

}
