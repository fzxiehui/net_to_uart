package mqtt

import (
	EMQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/fzxiehui/net_to_uart/log"
)

type MQTT struct {
	Client      EMQTT.Client
	Opts        *EMQTT.ClientOptions
	RecvHandler func(topic string, payload []byte) error
}

func NewMqtt(broker string,
	clientid string,
	username string,
	password string) *MQTT {
	opts := EMQTT.NewClientOptions().AddBroker(broker).SetClientID(clientid)
	opts.SetUsername(username)
	opts.SetPassword(password)
	return &MQTT{
		Opts:        opts,
		RecvHandler: nil,
	}
}

func (m *MQTT) Connect() error {
	c := EMQTT.NewClient(m.Opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	m.Client = c
	return nil
}

func (m *MQTT) Disconnect() {
	m.Client.Disconnect(250)
}

func (m *MQTT) Publish(topic string, payload []byte, qos int) error {

	if token := m.Client.Publish(topic, byte(qos), false, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	// m.Client.Publish(topic, byte(qos), false, payload)
	return nil
}

func (m *MQTT) Subscribe(topic string, qos int) error {
	if token := m.Client.Subscribe(topic, byte(qos), nil); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	m.Opts.SetDefaultPublishHandler(m.SubscribeHandler)
	return nil
}

func (m *MQTT) Unsubscribe(topic string) error {
	if token := m.Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (m *MQTT) SetRecvHandler(handler func(topic string, payload []byte) error) {
	m.RecvHandler = handler
}

func (m *MQTT) SubscribeHandler(client EMQTT.Client, msg EMQTT.Message) {

	if m.RecvHandler != nil {
		m.RecvHandler(msg.Topic(), msg.Payload())
	}
	log.Debugf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}
