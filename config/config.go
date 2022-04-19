package config

type Config struct {
    HostAddr string
    Frontend string
}

func Init() *Config {
    return &Config{
        HostAddr: "127.0.0.1:7696",
        Frontend: "./frontend",
    }
}
