package lib

type Config struct {
	Version string
}

func GetConfig() Config {
	return Config{
		Version: "0.0.1",
	}
}
