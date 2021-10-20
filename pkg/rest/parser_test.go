package rest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type QueryGetterMock map[string]string

func (q QueryGetterMock) Query(key string, defaultValue ...string) string {
	val, ok := q[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return val
}

func TestParser_StringFromEnum(t *testing.T) {
	testEnum := []string{"value 1", "value 2"}
	tests := map[string]struct {
		source QueryGetterMock
		key    string
		exp    string
		expErr bool
	}{
		"valid": {
			source: make(QueryGetterMock),
			key:    "key",
			exp:    "",
		},
		"empty": {
			source: QueryGetterMock{"key": "value 1"},
			key:    "key",
			exp:    "value 1",
		},
		"wrong": {
			source: QueryGetterMock{"key": "value 3"},
			key:    "key",
			expErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := QueryParam(tt.source).StringFromEnum(tt.key, testEnum)

			assert.Equal(t, tt.expErr, err != nil)
			assert.Equal(t, tt.exp, got)
		})
	}
}

func TestParser_SimplePagination(t *testing.T) {
	tests := map[string]struct {
		source QueryGetterMock
		exp    SimplePaginationParams
		expErr bool
	}{
		"valid": {
			source: QueryGetterMock{"limit": "1", "page": "2"},
			exp:    SimplePaginationParams{Page: 2, Limit: 1},
		},
		"valid_without_limit": {
			source: QueryGetterMock{"page": "3"},
			exp:    SimplePaginationParams{Page: 3},
		},
		"valid_without_page": {
			source: QueryGetterMock{"limit": "2"},
			exp:    SimplePaginationParams{Page: 1, Limit: 2},
		},
		"wrong_page": {
			source: QueryGetterMock{"page": "2f"},
			exp:    SimplePaginationParams{},
			expErr: true,
		},
		"wrong_limit": {
			source: QueryGetterMock{"limit": "2.2"},
			exp:    SimplePaginationParams{Page: 1},
			expErr: true,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := QueryParam(tt.source).SimplePagination()

			assert.Equal(t, tt.expErr, err != nil)
			assert.Equal(t, tt.exp, got)
		})
	}
}
