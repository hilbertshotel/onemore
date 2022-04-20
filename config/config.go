package config

import "time"

type Mongo struct {
	Uri      string
	Timeout  time.Duration
	Database string
	Coll     string
}

type Config struct {
	HostAddr string
	Frontend string
	Mongo    Mongo
}

func Init() *Config {
	return &Config{
		HostAddr: "127.0.0.1:7696",
		Frontend: "./frontend",
		Mongo: Mongo{
			Uri:      "mongodb://localhost:27017",
			Timeout:  time.Second * 10,
			Database: "onemore",
			Coll:     "habits",
		},
	}
}
