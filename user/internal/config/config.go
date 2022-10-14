package config

type Config struct {
	App   App   `yaml:"app"`
	HTTP  HTTP  `yaml:"http"`
	Twirp Twirp `yaml:"twirp"`
	Log   Log   `yaml:"log"`
	PG    PG    `yaml:"postgres"`
}

type (
	App struct {
		Name string `yaml:"name" env-required:"true"`
		Ver  string `yaml:"version" env-required:"true"`
	}

	HTTP struct {
		Port string `yaml:"port" env-required:"true"`
	}

	Twirp struct {
		Prefix string `yaml:"prefix" env-required:"true"`
	}

	Log struct {
		Level string `yaml:"level" env-required:"true"`
	}

	PG struct {
		MaxConns int    `yaml:"max_connections" env-required:"true"`
		URL      string `env:"PG_URL" env-required:"true"`
	}
)