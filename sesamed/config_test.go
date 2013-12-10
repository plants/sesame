package main

import (
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"

	"os"
	"testing"
)

type SetConfigSuite struct {
	suite.Suite
	logs *bytes.Buffer
}

func (suite *SetConfigSuite) SetupTest() {
	os.Clearenv()
	config = SesamedConfig{}

	suite.logs = new(bytes.Buffer)
	writer := bufio.NewWriter(suite.logs)
	log.SetOutput(writer)
}

func (suite *SetConfigSuite) TearDownTest() {
	log.SetOutput(os.Stdout)
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

func TestSetConfigSuite(t *testing.T) {
	suite.Run(t, new(SetConfigSuite))
}
