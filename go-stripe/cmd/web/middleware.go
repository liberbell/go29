package main

import "net/http"

func SessionLoad(next http.Hnadler) http.Handler {
	return session.LoadAndSave(next)
}
