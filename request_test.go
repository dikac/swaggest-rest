package rest_test

import (
	"testing"

	"github.com/dikac/swaggest-rest"
	"github.com/stretchr/testify/assert"
)

func TestRequestErrors_Error(t *testing.T) {
	err := rest.RequestErrors{
		"foo": []string{"bar"},
	}

	assert.EqualError(t, err, "bad request")
	assert.Equal(t, map[string]interface{}{"foo": []string{"bar"}}, err.Fields())
}
