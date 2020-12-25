package serialize_test

import (
	"encoding/json"
	"github.com/lyydhl/go-base-module/serialize"
	"testing"

	"github.com/stretchr/testify/require"
)

type City struct {
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
}

func TestToByte(t *testing.T) {

	city := City{
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

	bytes, err := serialize.ToBytes(city)
	require.NoError(t, err)
	require.NotEmpty(t, bytes)

	toStrings, err := serialize.ToString(bytes)
	require.NoError(t, err)
	require.NotEmpty(t, toStrings)

	var cc City
	err = json.Unmarshal(bytes, &cc)
	require.NoError(t, err)
	require.Equal(t, cc.ID, city.ID)
}
