package sesame

import (
	r "github.com/dancannon/gorethink"
	"github.com/kelseyhightower/envconfig"
	"strings"

	"net/url"
)

// UserStore manages connections for persisting Users to RethinkDB
type UserStore struct {
	conn *r.Session
}

// UserStoreConfig stores configuration from envconfig.
//
// - `URL`: something like rethinkdb://authkey@host:port/dbname
// - `PoolSize`: the maximum number of connections to hold, idles at half of
//   this value.
type UserStoreConfig struct {
	URL      string
	PoolSize int
}

// NewUserStore creates a new `UserStore` from the environment
func NewUserStore() (*UserStore, error) {
	us := &UserStore{}

	config := UserStoreConfig{}
	err := envconfig.Process("DB", &config)
	if err != nil {
		return us, err
	}

	// parse config from a string that looks like
	// `rethinkdb://accesskey@localhost:28015/dbname`
	url, err := url.Parse(config.URL)
	if err != nil {
		return us, err
	}

	options := map[string]interface{}{
		"address":  url.Host,
		"database": strings.Trim(url.Path, "/"),
	}
	if url.User != nil && url.User.Username() != "" {
		options["authkey"] = url.User.Username()
	}

	if config.PoolSize > 0 {
		options["maxActive"] = config.PoolSize
		options["maxIdle"] = config.PoolSize / 2
	}

	us.conn, err = r.Connect(options)
	if err != nil {
		return us, nil
	}

	return us, nil
}
