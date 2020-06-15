package config

import (
	"log"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

// ConfigList is Configファイルからデータ取得(config.ini)
type ConfigList struct {
	APIKey      string
	APISecret   string
	LogFile     string
	ProductCode string

	TradeDuration time.Duration
	Durations     map[string]time.Duration
	DbName        string
	UserName      string
	Password      string
	SQLDriver     string
	Port          int

	BackTest         bool
	UsePercent       float64
	DataLimit        int
	StopLimitPercent float64
	NumRanking       int
}

var Config ConfigList

func init() {
	cfg, err := ini.Load("config.ini")
	if err != nil {
		log.Printf("Failed to read file: %v", err)
		os.Exit(1)
	}

	durations := map[string]time.Duration{
		// "1s": time.Second,
		"1m": time.Minute,
		"1h": time.Hour,
		"1d": 24 * time.Hour,
	}

	Config = ConfigList{
		APIKey:    cfg.Section("bitflyer").Key("api_key").String(),
		APISecret: cfg.Section("bitflyer").Key("api_secret").String(),

		LogFile:     cfg.Section("gotrading").Key("log_file").String(),
		ProductCode: cfg.Section("gotrading").Key("product_code").String(),

		Durations:     durations,
		TradeDuration: durations[cfg.Section("gotrading").Key("trade_duration").String()],

		DbName:    cfg.Section("db").Key("name").String(),
		UserName:  cfg.Section("db").Key("username").String(),
		Password:  cfg.Section("db").Key("password").String(),
		SQLDriver: cfg.Section("db").Key("driver").String(),

		Port: cfg.Section("web").Key("port").MustInt(),

		BackTest:         cfg.Section("gotrading").Key("back_test").MustBool(),
		UsePercent:       cfg.Section("gotrading").Key("use_percent").MustFloat64(),
		DataLimit:        cfg.Section("gotrading").Key("data_limit").MustInt(),
		StopLimitPercent: cfg.Section("gotrading").Key("stop_limit_percent").MustFloat64(),
		NumRanking:       cfg.Section("gotrading").Key("num_ranking").MustInt(),
	}
}