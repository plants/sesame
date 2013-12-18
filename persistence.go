package sesame

import (
	"errors"
	r "github.com/dancannon/gorethink"
	"github.com/kelseyhightower/envconfig"
	"strings"
	"time"

	"net/url"
)

// UserStore manages connections for persisting Users to RethinkDB
type UserStore struct {
	conn  *r.Session
	table r.RqlTerm
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

	us.table = r.Db(options["database"]).Table("users")

	return us, nil
}

// Get takes an email address and returns a *User, or an error
func (store *UserStore) Get(email string) (*User, error) {
	u := new(User)

	row, err := store.table.GetAllByIndex("email", email).RunRow(store.conn)
	if err != nil {
		return u, err
	}
	if row.IsNil() {
		return u, errors.New("could not find a user with email \"" + email + "\"")
	}

	row.Scan(&u)

	return u, nil
}

// Save takes a *User and saves it to RethinkDB. It updates User.Updated, as well.
func (store *UserStore) Save(user *User) error {
	user.Updated = time.Now()

	response, err := store.table.Insert(user, r.InsertOpts{Upsert: true}).RunWrite(store.conn)
	if err != nil {
		return err
	}

	if response.Inserted == 1 && len(response.GeneratedKeys) > 0 {
		user.Id = response.GeneratedKeys[0]
	}

	return nil
}

// Delete deletes a user by email.
func (store *UserStore) Delete(email string) error {
	response, err := store.table.GetAllByIndex("email", email).Delete().RunWrite(store.conn)
	if err != nil {
		return err
	}
	if response.Deleted <= 0 {
		return errors.New("no user with email \"" + email + "\" to delete")
	}
	return nil
}
