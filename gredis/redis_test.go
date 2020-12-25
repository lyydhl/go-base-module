package gredis_test

import (
	"github.com/lyydhl/go-base-module/gredis"
	"github.com/lyydhl/go-base-module/serialize"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func setup() error {
	return gredis.Setup(30, 30, time.Duration(time.Second*200), "", "")
}

func TestRedisSetToByte(t *testing.T) {
	err := setup()
	require.NoError(t, err)

	city := struct {
		ID         int     `json:"id"`
		CreatedOn  int     `json:"created_on"`
		ModifiedOn int     `json:"modified_on"`
		DeletedOn  int     `json:"deleted_on"`
		Code       string  `json:"code"`
		Title      string  `json:"title"`
		Lat        float64 `json:"lat"`
		Lng        float64 `json:"lng"`
		IsComplete bool    `json:"is_complete"`
		ModifiedBy string  `json:"modified_by"`
		DeletedBy  string  `json:"deleted_by"`
	}{
		ID:         108,
		CreatedOn:  1596368086,
		ModifiedOn: 1596445521,
		DeletedOn:  0,
		Code:       "350200",
		Title:      "厦门市",
		Lat:        24.485406605176305,
		Lng:        118.09643549976651,
		IsComplete: true,
		ModifiedBy: "",
		DeletedBy:  "",
	}

	bytes, err := serialize.GetBytes(city)
	require.NoError(t, err)
	require.NotEmpty(t, bytes)

	err = gredis.SetToBytes("temp", bytes, 20)
	require.NoError(t, err)

	get, err := gredis.Get("temp")
	require.NoError(t, err)
	require.NotEmpty(t, get)
}
