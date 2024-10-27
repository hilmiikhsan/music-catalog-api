package configs

type (
	Config struct {
		Service  Service  `mapstructure:"service"`
		Database Database `mapstructure:"database"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretKey string `mapstructure:"secret_key"`
	}

	Database struct {
		DataSourceName string `mapstructure:"data_source_name"`
	}
)
