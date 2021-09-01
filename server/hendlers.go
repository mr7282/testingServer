package server

import "net/http"

func (serv *Server) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!!!\n"))
}