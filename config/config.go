// Package config @Author Zhan 2024/1/18 10:00:00
package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	SerialPort     string   `mapstructure:"serial_port" yaml:"serial_port"`
	SensorNum      int      `mapstructure:"sensor_num" yaml:"sensor_num"`
	SmtpHost       string   `mapstructure:"smtp_host" yaml:"smtp_host"`
	SmtpPort       int      `mapstructure:"smtp_port" yaml:"smtp_port"`
	SmtpEmail      string   `mapstructure:"smtp_email" yaml:"smtp_email"`
	SmtpUsername   string   `mapstructure:"smtp_username" yaml:"smtp_username"`
	SmtpPassword   string   `mapstructure:"smtp_password" yaml:"smtp_password"`
	MailReceiver   []string `mapstructure:"mail_receiver" yaml:"mail_receiver"`
	MailDelay      int      `mapstructure:"mail_delay" yaml:"mail_delay"`
	MaxTemp        float64  `mapstructure:"max_temp" yaml:"max_temp"`
	MinTemp        float64  `mapstructure:"min_temp" yaml:"min_temp"`
	SampleInterval int      `mapstructure:"sample_interval" yaml:"sample_interval"`
	Position       string   `mapstructure:"position" yaml:"position"`
}

func Viper(path ...string) Config {
	var configPath string
	var config Config
	if len(path) == 0 {
		flag.StringVar(&configPath, "c", "", "Read configuration from the specified file.")
		flag.Parse()
		if configPath == "" {
			panic("Please set the configuration file by -c [conf_path]")
		} else {
			fmt.Printf("Use the configuration file from command flag.\n")
		}
	} else {
		configPath = path[0]
		fmt.Printf("Use the configuration specified by argument.\n")
	}

	v := viper.New()
	fmt.Printf("config path: %s\n", configPath)
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic("Fail to read configuration file.")
	}

	if err := v.Unmarshal(&config); err != nil {
		panic("Fail to unmarshal configuration.")
	}
	return config
}
