package emeBartApi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/parnurzeal/gorequest"
	"github.com/unrolled/render"
)

func Test_HelloWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "http://emebartapi.herokuapp.com/api/route.aspx?cmd=routeinfo&route=6&key=MW9S-E7SL-26DU-VV8V", nil)
	if err != nil {
		t.Fatal(err)
	}
	res := httptest.NewRecorder()
	controller := Controller{render.New(), gorequest.New(), "http://api.bart.gov"}
	controller.Routes(res, req)

	exp := "Hello World"
	act := res.Body.String()
	if exp != act {
		t.Fatalf("Expected %s gog %s", exp, act)
	}
}
