package sesame

import (
	"errors"
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"os"
	"testing"
	"time"
)

func newTestUserStore() (*UserStore, error) {
	os.Clearenv()
	os.Setenv("DB_URL", "rethinkdb://localhost:28015/test")
	os.Setenv("DB_POOLSIZE", "5")

	return NewUserStore()
}

func TestNewUserStore(t *testing.T) {
	store, err := newTestUserStore()
	assert.Nil(t, err)
	assert.NotNil(t, store)
}

type UserStoreSuite struct {
	suite.Suite
	store *UserStore
	user  *User
}

func (suite *UserStoreSuite) SetupSuite() {
	store, err := newTestUserStore()
	if err != nil {
		panic(err)
	}
	suite.store = store

	// create necessary DB, table, indices
	_, err = r.DbCreate("test").Run(suite.store.conn)
	if err != nil {
		panic(err)
	}
	_, err = r.Db("test").TableCreate("users").Run(suite.store.conn)
	if err != nil {
		panic(err)
	}
	_, err = r.Db("test").Table("users").IndexCreate("email").Run(suite.store.conn)
	if err != nil {
		panic(err)
	}

	suite.user = NewUser("test@example.com", "password")
	_, err = r.Db("test").Table("users").Insert(suite.user).RunWrite(suite.store.conn)
	if err != nil {
		panic(err)
	}
}

func (suite *UserStoreSuite) TearDownSuite() {
	_, err := r.DbDrop("test").Run(suite.store.conn)
	if err != nil {
		fmt.Println("dropping \"test\" failed. Do it by hand before re-running")
	}
}

func (suite *UserStoreSuite) TestGetValidUser() {
	u, err := suite.store.Get(suite.user.Email)

	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), suite.user.Email, u.Email)
	assert.Equal(suite.T(), suite.user.Password, u.Password)
	assert.Equal(suite.T(), suite.user.Salt, u.Salt)
	assert.WithinDuration(suite.T(), suite.user.Updated, u.Updated, 1*time.Second)
	assert.WithinDuration(suite.T(), suite.user.Created, u.Created, 1*time.Second)
}

func (suite *UserStoreSuite) TestGetInvalidUser() {
	_, err := suite.store.Get("bad@example.com")

	assert.Equal(suite.T(), errors.New("could not find a user with email \"bad@example.com\""), err)
}

func TestUserStoreSuite(t *testing.T) {
	suite.Run(t, new(UserStoreSuite))
}
