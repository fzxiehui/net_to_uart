package mqtt2uart

import (
	"time"

	"github.com/fzxiehui/net_to_uart/config"
	"github.com/fzxiehui/net_to_uart/log"
	"github.com/fzxiehui/net_to_uart/pkg/mqtt"
	"github.com/fzxiehui/net_to_uart/pkg/uart"
)

type MQTT2UART interface {
	Start()
	Stop()
}

type mqtt2uart struct {
	MQTT2UART
	MqttClient *mqtt.MQTT
	UartClient *uart.UART
}

var m2u *mqtt2uart

func NewMQTT2UART() MQTT2UART {
	cfg := config.Config()
	mqttclient := mqtt.NewMqtt(cfg.GetString("mqtt.broker"),
		cfg.GetString("mqtt.clientid"),
		cfg.GetString("mqtt.username"),
		cfg.GetString("mqtt.password"))

	uartclient := uart.NewUart(cfg.GetString("uart.port"), cfg.GetInt("uart.baudrate"))

	return &mqtt2uart{
		MqttClient: mqttclient,
		UartClient: uartclient,
	}
}

func mqMsg(topic string, payload []byte) error {
	log.Debug("MQMSG")
	log.Debugf("mqtt msg: %s, %s", topic, payload)

	m2u.UartClient.SendChan <- payload

	return nil
}

func (m *mqtt2uart) Start() {

	err := m.MqttClient.Connect()
	if err != nil {
		panic(err)
	}

	m.MqttClient.Subscribe(config.Config().GetString("mqtt.sub_topic"), 0)

	m2u = m
	m.MqttClient.SetRecvHandler(mqMsg)

	m.UartClient.Start()

	for {
		select {
		case msg := <-m.UartClient.RecvChan:
			m.MqttClient.Publish(config.Config().GetString("mqtt.pub_topic"), msg, 0)
		case <-time.After(1 * time.Second):
		}
	}
}

func (m *mqtt2uart) Stop() {
	m.MqttClient.Disconnect()
	m.UartClient.Close()
}
