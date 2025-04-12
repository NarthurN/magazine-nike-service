package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type configs struct {
	BDConfigs
	HTTPServerConfigs
}

type BDConfigs struct {
	DBuser string
	DBpass string
	DBhost string
	DBport string
}

type HTTPServerConfigs struct {
	SRVaddress           string
	SRVport              string
	SRVreadTimeout       time.Duration
	SRVreadHeaderTimeout time.Duration
	SRVwriteTimeout      time.Duration
	SRVidleTimeout       time.Duration
}

func Load() (*configs, error) {
	configs := &configs{}
	if err := godotenv.Load(`.env`); err != nil {
		return configs, fmt.Errorf("error loading .env file: %w", err)
	}
	// DB configs
	configs.DBuser = os.Getenv("DB_USER")
	configs.DBpass = os.Getenv("DB_PASS")
	configs.DBhost = os.Getenv("DB_HOST")
	configs.DBport = os.Getenv("DB_PORT")

	// HTTP configs
	configs.SRVaddress = os.Getenv("SRV_ADDRESS")
	configs.SRVport = os.Getenv("SRV_PORT")

	// Парсим строки в time.Duration
	var err error
	configs.SRVreadTimeout, err = time.ParseDuration(os.Getenv("SRV_ReadTimeout"))
	if err != nil {
		return configs, fmt.Errorf("invalid SRV_ReadTimeout format: %w", err)
	}

	configs.SRVreadHeaderTimeout, err = time.ParseDuration(os.Getenv("SRV_ReadHeaderTimeout"))
	if err != nil {
		return configs, fmt.Errorf("invalid SRV_ReadHeaderTimeout format: %w", err)
	}

	configs.SRVwriteTimeout, err = time.ParseDuration(os.Getenv("SRV_WriteTimeout"))
	if err != nil {
		return configs, fmt.Errorf("invalid SRV_WriteTimeout format: %w", err)
	}

	configs.SRVidleTimeout, err = time.ParseDuration(os.Getenv("SRV_IdleTimeout"))
	if err != nil {
		return configs, fmt.Errorf("invalid SRV_IdleTimeout format: %w", err)
	}

	return configs, nil
}
