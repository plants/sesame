package sesame

import (
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
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
