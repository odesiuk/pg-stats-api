package cmd

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_index(t *testing.T) {
	app := Setup(nil, Config{AppName: "TEST"})
	resp, _ := app.Test(httptest.NewRequest("GET", "/", nil))
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, `{"name":"TEST"}`, string(body))
}
