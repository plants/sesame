package sesame

import (
	"github.com/eaigner/hood"
	"github.com/kelseyhightower/envconfig"
)

type UserStore struct {
	conn *hood.Hood
}

type UserStoreConfig struct {
	Driver string
	Args   string
}

func NewUserStore() (*UserStore, error) {
	us := &UserStore{}

	var config UserStoreConfig
	err := envconfig.Process("DB", &config)
	if err != nil {
		return us, err
	}

	us.conn, err = hood.Open(config.Driver, config.Args)
	if err != nil {
		return us, err
	}

	return us, nil
}
