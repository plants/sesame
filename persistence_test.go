package sesame

import (
	"github.com/stretchr/testify/assert"

	"os"
	"testing"
)

func newTestUserStore() (*UserStore, error) {
	os.Clearenv()
	os.Setenv("DB_DRIVER", "postgres")
	os.Setenv("DB_ARGS", "asdf")

	return NewUserStore()
}

func TestNewUserStore(t *testing.T) {
	_, err := newTestUserStore()
	assert.Nil(t, err)
}
