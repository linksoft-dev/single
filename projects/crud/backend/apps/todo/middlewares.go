package todo

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Infof("Method %s url: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
