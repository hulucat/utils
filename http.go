package utils

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(address string) (rt []byte, err error) {
	Debugf("Http get address: %s", address)
	resp, err := http.Get(address)
	if err != nil {
		Errorf("Error do http get to %s: %s", address, err.Error())
		return
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
