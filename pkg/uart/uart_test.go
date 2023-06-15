package uart

import (
	"testing"
	"time"

	"github.com/fzxiehui/net_to_uart/config"
)

func TestUart(t *testing.T) {
	cfg := config.Config()
	uart := NewUart(cfg.GetString("uart.port"), cfg.GetInt("uart.baudrate"))
	uart.Start()
	uart.SendChan <- []byte("hello")
	uart.SendChan <- []byte("world")
	uart.SendChan <- []byte("!")
	uart.SendChan <- []byte("!")
	uart.SendChan <- []byte("!")
	time.Sleep(2 * time.Second)
}
