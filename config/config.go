package config

import "time"

type Appconfig struct {
	Web struct {
		APIHost            string        `yaml:"apiHost" env:"RISK_MANAGEMENT_WEB_API_HOST"`
		DebugHost          string        `yaml:"debugHost" env:"RISK_MANAGEMENT_WEB_DEBUG_HOST"`
		ReadTimeout        time.Duration `yaml:"readTimeout" env:"RISK_MANAGEMENT_WEB_READ_TIMEOUT"`
		WriteTimeout       time.Duration `yaml:"writeTimeout" env:"RISK_MANAGEMENT_WEB_WRITE_TIMEOUT"`
		ChannelDialTimeout time.Duration `yaml:"channelDialTimeout" env:"RISK_MANAGEMENT_WEB_CHANNEL_DIAL_TIMEOUT"`
		IdleTimeout        time.Duration `yaml:"idleTimeout" env:"RISK_MANAGEMENT_WEB_IDLE_TIMEOUT"`
		ShutdownTimeout    time.Duration `yaml:"shutdownTimeout" env:"RISK_MANAGEMENT_WEB_SHUTDOWN_TIMEOUT"`
	} `yaml:"web"`
}