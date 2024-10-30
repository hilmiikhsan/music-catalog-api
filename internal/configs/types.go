package configs

type (
	Config struct {
		Service       Service       `mapstructure:"service"`
		Database      Database      `mapstructure:"database"`
		SpotifyConfig SpotifyConfig `mapstructure:"spotify_config"`
	}

	Service struct {
		Port      string `mapstructure:"port"`
		SecretKey string `mapstructure:"secret_key"`
	}

	Database struct {
		DataSourceName string `mapstructure:"data_source_name"`
	}

	SpotifyConfig struct {
		ClientID     string `mapstructure:"client_id"`
		ClientSecret string `mapstructure:"client_secret"`
	}
)
