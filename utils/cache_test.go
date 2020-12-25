package utils_test

import (
	"github.com/lyydhl/go-base-module/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCache(t *testing.T) {
	key := "testing_key"
	utils.SetCache(key, key, 5)

	cache, b := utils.GetCache(key)
	require.Equal(t, b, true)
	require.NotEmpty(t, cache)

	utils.DeleteCache(key)

	cache, b = utils.GetCache(key)
	require.Equal(t, b, false)
	require.Empty(t, cache)
}
