package uart

import (
	"log"

	"github.com/tarm/serial"
)

type Uart struct {
	serialClient *serial.Port
	SendChan     chan []byte
	RecvChan     chan []byte
}

func NewUart(name string, baud int) *Uart {
	c := serial.Config{Name: name, Baud: baud}
	s, err := serial.OpenPort(&c)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	sendChan := make(chan []byte, 512)
	recvChan := make(chan []byte, 512)

	return &Uart{
		serialClient: s,
		SendChan:     sendChan,
		RecvChan:     recvChan,
	}
}

func (u *Uart) Send() {
	for {
		data := <-u.SendChan
		u.serialClient.Write(data)
	}
}

func (u *Uart) Recv() {
	for {
		data := make([]byte, 512)
		n, err := u.serialClient.Read(data)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		u.RecvChan <- data[:n]
	}
}

func (u *Uart) Close() {
	u.serialClient.Close()
}

func (u *Uart) Start() {
	go u.Send()
	go u.Recv()
}
