package utils

import (
	"encoding/json"
	"fmt"
	"log"
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

type RouteParameter struct {
	Path    string
	Verb    HttpVerb
	Handler func(http.ResponseWriter, *Request)
}
type Route struct {
	Path       string
	Parameters []string
	Handler    func(http.ResponseWriter, *Request)
}

type errorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

type Router struct {
	routes map[HttpVerb][]Route
}

func checkRoute(uri string, r Route) (map[string]string, error) {
	passes, _ := regexp.Match(r.Path, []byte(uri))

	if !passes {
		return nil, fmt.Errorf("Route not found")
	}

	re := regexp.MustCompile(r.Path)
	parsed := re.FindStringSubmatch(uri)[1:]
	parameters := make(map[string]string)

	for i, key := range r.Parameters {
		parameters[key] = parsed[i]
	}

	return parameters, nil
}

func getParameters(pattern string) (string, []string, error) {
	const uriParameter = `(/\w+)/(?P<params>{(?P<key>\w+)\:(?P<reg>[^{}]+)})?`
	var keys []string = []string{}
	passes, ok := regexp.Match(uriParameter, []byte(pattern))
	if ok != nil {
		return "", keys, fmt.Errorf("Bad construction of regexp: %s", ok)
	}
	if !passes {
		return "", keys, fmt.Errorf("The regexp pattern doesn't match with uri parameter rules")
	}

	re := regexp.MustCompile(uriParameter)
	uriToMatch := fmt.Sprintf("^%s$", re.ReplaceAllString(pattern, "$1/($4)"))
	keys = strings.Split(strings.Trim(re.ReplaceAllString(pattern, "$3 "), " "), " ")

	return uriToMatch, keys, nil
}

func GetErrorResponse(w http.ResponseWriter, err string, status int) {
	serialized, _ := json.Marshal(errorResponse{Error: err, Code: status})

	http.Error(w, string(serialized), status)
}

func (r *Router) ProcessRoutes(routes map[string][]RouteParameter) *Router {
	r.routes = make(map[HttpVerb][]Route)
	for key, routeModule := range routes {
		log.Println("Injecting route module " + key)
		for _, routeParameter := range routeModule {
			uriToMatch, parameters, ok := getParameters(routeParameter.Path)
			if ok != nil {
				log.Println("Error processing route " + routeParameter.Path)
				break
			}

			route := Route{
				Path:       uriToMatch,
				Parameters: parameters,
				Handler:    routeParameter.Handler,
			}

			if r.routes[routeParameter.Verb] == nil {
				r.routes[routeParameter.Verb] = make([]Route, 0)
			}
			r.routes[routeParameter.Verb] = append(r.routes[routeParameter.Verb], route)
		}
	}
	return r
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	route := req.URL.Path
	for _, r := range r.routes[HttpVerb(req.Method)] {
		parameters, err := checkRoute(route, r)
		if err != nil {
			continue
		}
		request := Request{Original: req, Params: parameters}
		r.Handler(w, &request)
		return
	}

	GetErrorResponse(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}
