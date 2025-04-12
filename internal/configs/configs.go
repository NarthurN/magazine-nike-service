package configs

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type BDConfigs struct {
	DBuser      string
	DBpass      string
	DBhost      string
	DBport      string
	DBtablename string
}

type HTTPServerConfigs struct {
	SRVaddress           string
	SRVport              string
	SRVreadTimeout       time.Duration
	SRVreadHeaderTimeout time.Duration
	SRVwriteTimeout      time.Duration
	SRVidleTimeout       time.Duration
}

func LoadDB() (*BDConfigs, error) {
	BDConfigs := &BDConfigs{}
	if err := godotenv.Load(`.env.database`); err != nil {
		return BDConfigs, fmt.Errorf("error loading .env file: %w", err)
	}
	// DB configs
	BDConfigs.DBuser = os.Getenv("DB_USER")
	BDConfigs.DBpass = os.Getenv("DB_PASS")
	BDConfigs.DBhost = os.Getenv("DB_HOST")
	BDConfigs.DBport = os.Getenv("DB_PORT")
	BDConfigs.DBtablename = os.Getenv("DB_TABLENAME")

	return BDConfigs, nil
}

func LoadHTTP() (*HTTPServerConfigs, error) {
	HTTPServerConfigs := &HTTPServerConfigs{}
	if err := godotenv.Load(`.env.http`); err != nil {
		return HTTPServerConfigs, fmt.Errorf("error loading .env file: %w", err)
	}
	// HTTP configs
	HTTPServerConfigs.SRVaddress = os.Getenv("SRV_ADDRESS")
	HTTPServerConfigs.SRVport = os.Getenv("SRV_PORT")

	// Парсим строки в time.Duration
	var err error
	HTTPServerConfigs.SRVreadTimeout, err = time.ParseDuration(os.Getenv("SRV_ReadTimeout"))
	if err != nil {
		return HTTPServerConfigs, fmt.Errorf("invalid SRV_ReadTimeout format: %w", err)
	}

	HTTPServerConfigs.SRVreadHeaderTimeout, err = time.ParseDuration(os.Getenv("SRV_ReadHeaderTimeout"))
	if err != nil {
		return HTTPServerConfigs, fmt.Errorf("invalid SRV_ReadHeaderTimeout format: %w", err)
	}

	HTTPServerConfigs.SRVwriteTimeout, err = time.ParseDuration(os.Getenv("SRV_WriteTimeout"))
	if err != nil {
		return HTTPServerConfigs, fmt.Errorf("invalid SRV_WriteTimeout format: %w", err)
	}

	HTTPServerConfigs.SRVidleTimeout, err = time.ParseDuration(os.Getenv("SRV_IdleTimeout"))
	if err != nil {
		return HTTPServerConfigs, fmt.Errorf("invalid SRV_IdleTimeout format: %w", err)
	}

	return HTTPServerConfigs, nil
}
