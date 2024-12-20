package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Controller) Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Gaia Exporter is OK")
}

func (c *Controller) GaiaGetInfo(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Gaiad unavailible: %v", err)
		return nil
	}

	data, _ := ioutil.ReadAll(resp.Body)

	return data
}
