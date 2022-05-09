package config

func Default() Config {
	return Config{
		Listen: ":1373",
	}
}
