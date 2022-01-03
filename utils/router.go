package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

type Request struct {
	Original *http.Request
	Params   map[string]string
}

type HttpVerb string

const (
	Get    HttpVerb = "GET"
	Post   HttpVerb = "POST"
	Put    HttpVerb = "PUT"
	Delete HttpVerb = "DELETE"
	Patch  HttpVerb = "PATCH"
)

type Route struct {
	Path    string
	Verb    HttpVerb
	Handler func(http.ResponseWriter, *Request)
}

const uriParameter = `(/\w+)/(?P<params>{(?P<key>\w+)\:(?P<reg>[^{}]+)})?`

func getParameters(uri, pattern string) (map[string]string, error, int) {
	parameters := make(map[string]string)
	passes, ok := regexp.Match(uriParameter, []byte(pattern))
	if ok != nil {
		return nil, fmt.Errorf("Bad construction of regexp: %s", ok), http.StatusInternalServerError
	}
	if !passes {
		return nil, fmt.Errorf("The regexp pattern doesn't match with uri parameter rules"), http.StatusInternalServerError
	}

	re := regexp.MustCompile(uriParameter)
	keys := strings.Split(strings.Trim(re.ReplaceAllString(pattern, "$3 "), " "), " ")
	uriToMatch := fmt.Sprintf("^%s$", re.ReplaceAllString(pattern, "$1/($4)"))
	passes, _ = regexp.Match(uriToMatch, []byte(uri))

	if !passes {
		return nil, fmt.Errorf("Route not found"), http.StatusNotFound
	}

	re = regexp.MustCompile(uriToMatch)
	parsed := re.FindStringSubmatch(uri)[1:]
	for i, key := range keys {
		parameters[key] = parsed[i]
	}

	return parameters, nil, 0
}

type errorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func GetErrorResponse(w http.ResponseWriter, err string, status int) {
	serialized, _ := json.Marshal(errorResponse{Error: err, Code: status})

	http.Error(w, string(serialized), status)
}

func Bypass(route Route) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if route.Verb != HttpVerb(r.Method) {
			GetErrorResponse(w, fmt.Sprintf("Method %s not allowed", r.Method), http.StatusMethodNotAllowed)
			return
		}
		parameters, ok, code := getParameters(r.URL.Path, route.Path)

		if ok != nil {
			GetErrorResponse(w, ok.Error(), code)
			return
		}
		request := Request{Original: r, Params: parameters}
		route.Handler(w, &request)
	}
}
