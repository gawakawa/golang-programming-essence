package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type routerParam map[string]string

type routerFunc func(routerParam, http.ResponseWriter, *http.Request)

type routerItem struct {
	// HTTP method
	method string

	// matcher with request URI
	matcher *regexp.Regexp

	// called function
	fnc routerFunc
}

type router struct {
	items []routerItem
}

func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) POST(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodPost,
		matcher: regexp.MustCompile(prefix),
		fnc:     fnc,
	})
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range rt.items {
		// request methods correspond and request URIs match
		if v.method == r.Method &&
			v.matcher.MatchString(r.RequestURI) {
			// get path parameters
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			param := make(routerParam)
			for i, name := range v.matcher.SubexpNames() {
				param[name] = match[i]
			}
			v.fnc(param, w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	rt := router{}

	rt.GET(
		`^/$`,
		func(p routerParam, w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "Hello")
		},
	)

	rt.GET(
		`^/(?P<name>\w+)$`,
		func(p routerParam, w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello: %v\n", p["name"])
		},
	)

	rt.POST(
		`^/api$`,
		func(p routerParam, w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "text/json")
			fmt.Fprintln(w, `{"status": "OK"}`)
		})

	http.ListenAndServe(":8080", &rt)
}
