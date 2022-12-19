package server_test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"rest_app/server"
	"testing"
)

func TestNewServer_OK(test *testing.T) {
	ginActual := server.NewServer()
	assert.NotNil(test, ginActual)
	assert.IsType(test, &gin.Engine{}, ginActual)
	assert.Equal(test, gin.Mode(), gin.DebugMode)
}

