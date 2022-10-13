package main

import (
	"navi/kstgo/httplib"
	"testing"
	"time"
)

func TestSvc(t *testing.T) {
	resp, _ := httplib.Post("http://172.23.117.152:9999/svc").DoRequest()

	_ = resp.Body.Close()

	time.Sleep(time.Second * 10)
}
