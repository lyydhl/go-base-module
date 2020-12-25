package utils_test

import (
	"github.com/lyydhl/go-base-module/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRemoveDuplicateElement(t *testing.T) {
	sli := []string{"a", "b", "c", "d", "d", "e", "e"}
	sli = sli[:]

	element := utils.RemoveDuplicateElement(sli)
	require.NotEmpty(t, element)
	require.Equal(t, len(element), 5)
}
