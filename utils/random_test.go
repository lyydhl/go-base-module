package utils_test

import (
	"github.com/lyydhl/go-base-module/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomString(t *testing.T) {
	s1 := utils.RandomString(6)
	s2 := utils.RandomString(6)
	require.NotEmpty(t, s1)
	require.NotEmpty(t, s2)
	require.NotEqual(t, s1, s2)
	require.Equal(t, len(s1), 6)
	require.Equal(t, len(s2), 6)
}

func TestRandomInt(t *testing.T) {
	n1 := utils.RandomInt(8)
	n2 := utils.RandomInt(8)
	require.NotEmpty(t, n1)
	require.NotEmpty(t, n2)
	require.NotEqual(t, n1, n2)
	require.Equal(t, len(n1), 8)
	require.Equal(t, len(n2), 8)
}
