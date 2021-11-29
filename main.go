package main

import (
	"net/http"

	"github.com/csillagrobert95/gowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
