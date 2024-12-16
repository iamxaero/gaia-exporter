package controller

import (
	"fmt"
	"net/http"
)

func (c *Controller) Health(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Gaia Exporter is OK")
}
