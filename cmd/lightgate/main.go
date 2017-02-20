package main

import (
	"net/http"

	"github.com/Akagi201/light"
	jwtmiddleware "github.com/Akagi201/middleware/jwt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pressly/lg"
	"github.com/vulcand/oxy/forward"
	"github.com/vulcand/oxy/testutils"
)

func main() {
	fwd, err := forward.New()
	if err != nil {
		logger.Fatalf("HTTP forward create failed, err: %v", err)
	}

	root := light.New()

	root.Use(lg.RequestLogger(logger))

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		Debug: true,
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("supersecret"), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	root.Use(jwtMiddleware.Handler)

	root.HandleAll("/*path", func(w http.ResponseWriter, r *http.Request) {
		r.URL = testutils.ParseURI("http://localhost:8327")
		fwd.ServeHTTP(w, r)
	})

	logger.Infof("lightgate HTTP listening at %s", opts.HTTPListenAddr)
	logger.Fatalln(http.ListenAndServe(opts.HTTPListenAddr, root))
}
