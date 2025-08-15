package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"
)

type MyContext struct {
	db *sql.DB
}

func NewMyContext() *MyContext {
	return &MyContext{}
}

func (m *MyContext) handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		f, err := os.Open(
			"chapter05/simple-net-http/content.txt",
		)
		if err != nil {
			http.Error(
				w,
				err.Error(),
				http.StatusInternalServerError,
			)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	default:
	}
}

func main() {
	myctx := NewMyContext()
	http.HandleFunc("/", myctx.handle)
	http.ListenAndServe(":8080", nil)
}
