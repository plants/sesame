package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"os"
	"testing"
)

type SetConfigSuite struct {
	suite.Suite
}

func (suite *SetConfigSuite) SetupTest() {
	os.Clearenv()
	config = SesamedConfig{}
}

func (suite *SetConfigSuite) TestDebug() {
	// default
	parseConfig()
	assert.Equal(suite.T(), config.Debug, false)

	// non-default
	os.Setenv("SESAMED_DEBUG", "true")
	parseConfig()
	assert.Equal(suite.T(), config.Debug, true)
}

func (suite *SetConfigSuite) TestPort() {
	// default
	parseConfig()
	assert.Equal(suite.T(), config.Port, 2884)

	// non-default
	os.Setenv("SESAMED_PORT", "8080")
	parseConfig()
	assert.Equal(suite.T(), config.Port, 8080)
}

func (suite *SetConfigSuite) TestStorageType() {
	// default
	parseConfig()
	assert.Equal(suite.T(), config.StorageType, "memory")

	// non-default
	os.Setenv("SESAMED_STORAGETYPE", "rethinkdb")
	parseConfig()
	assert.Equal(suite.T(), config.StorageType, "rethinkdb")
}

func TestSetConfigSuite(t *testing.T) {
	suite.Run(t, new(SetConfigSuite))
}
