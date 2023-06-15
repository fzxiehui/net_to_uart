package uart

import (
	"fmt"
	"testing"
)

func TestUart(t *testing.T) {
	uart := NewUart("/dev/serial/by-id/usb-FTDI_FT232R_USB_UART_AB0PFGMV-if00-port0", 9600)
	uart.Start()
	defer uart.Close()

	for {
		select {
		case data := <-uart.RecvChan:
			fmt.Println(data)
		}
	}
}
