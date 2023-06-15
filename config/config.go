package config

import (
	"time"

	"github.com/spf13/viper"
)

// Provider defines a set of read-only methods for accessing the application
// configuration params as defined in one of the config files.
type Provider interface {
	ConfigFileUsed() string
	Get(key string) interface{}
	GetBool(key string) bool
	GetDuration(key string) time.Duration
	GetFloat64(key string) float64
	GetInt(key string) int
	GetInt64(key string) int64
	GetSizeInBytes(key string) uint
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringMapStringSlice(key string) map[string][]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	InConfig(key string) bool
	IsSet(key string) bool
}

var defaultConfig *viper.Viper

// Config returns a default config providers
func Config() Provider {
	return defaultConfig
}

// LoadConfigProvider returns a configured viper instance
func LoadConfigProvider(appName string) Provider {
	return readViperConfig(appName)
}

func init() {
	defaultConfig = readViperConfig("NET_TO_UART")
}

func ReadViperConfigFromFile(configPath string) error {
	defaultConfig.SetConfigFile(configPath)
	// set defaultConfig
	return defaultConfig.ReadInConfig()

}

func readViperConfig(appName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()

	// global defaults

	v.SetDefault("json_logs", false)
	v.SetDefault("loglevel", "debug")
	v.SetDefault("uart.port", "/dev/serial/by-id/usb-FTDI_FT232R_USB_UART_AB0PFGMV-if00-port0")
	v.SetDefault("uart.baudrate", 9600)

	v.SetDefault("mqtt.broker", "tcp://192.168.12.199:1883")
	v.SetDefault("mqtt.clientid", "3722600978")
	v.SetDefault("mqtt.username", "LOpbUOpEdK")
	v.SetDefault("mqtt.password", "updOMeRVja")
	v.SetDefault("mqtt.pub_topic", "/pub/model/3416275346/685104052/3722600978")
	v.SetDefault("mqtt.sub_topic", "/sub/model/3416275346/685104052/3722600978")

	return v
}
