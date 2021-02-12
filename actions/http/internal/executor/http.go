package executor

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/wesovilabs/orion/internal/errors"
)

var httpClient = http.DefaultClient

var DefaultHTTPConnection = &Connection{
	Timeout: 10 * time.Second,
	Proxy:   "",
}

var DefaultCookie = &Cookie{
	Name:   "",
	Path:   "",
	Domain: "",
	Value:  "",
}

type HTTP struct {
	URL         string
	Method      string
	Body        string
	Host        string
	Headers     map[string][]string
	QueryParams map[string][]string
	Connection  *Connection
	Cookies     []*Cookie
	Auth        *Auth
}

type Cookie struct {
	Name   string
	Path   string
	Domain string
	Value  string
}

type Connection struct {
	Timeout time.Duration
	Proxy   string
}

type Auth struct {
	Kind  string
	Value string
}

func (auth *Auth) headerValue() string {
	return fmt.Sprintf("%s %s", auth.Kind, auth.Value)
}

func (action *HTTP) httpMethod() string {
	return strings.ToUpper(action.Method)
}

func (action *HTTP) buildRequest() (*http.Request, errors.Error) {
	req, err := http.NewRequest(action.httpMethod(), action.URL, strings.NewReader(action.Body))
	if err != nil {
		return nil, errors.Unexpected(err.Error())
	}
	req.Header = action.Headers
	if log.IsLevelEnabled(log.TraceLevel) && len(action.Headers) > 0 {
		for name, values := range action.Headers {
			log.Tracef("Set header '%s' with value: '%v'", name, values)
		}
	}
	if action.Host != "" {
		log.Tracef("Set host to '%s'", action.Host)
		req.Host = action.Host
	}
	if action.Auth != nil {
		log.Tracef("Set auth token to '%s'", action.Auth.headerValue())
		req.Header.Add("Authorization", action.Auth.headerValue())
	}
	q := req.URL.Query()
	for name, value := range action.QueryParams {
		log.Tracef("Add query param '%s' with value '%v'", name, value)
		for i := range value {
			q.Add(name, value[i])
		}
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (action *HTTP) Execute() (Variables, errors.Error) {
	vars := createVariables()
	req, err := action.buildRequest()
	if err != nil {
		return vars, err
	}
	client := action.buildHTTPClient()
	action.setCookies(client, req.URL)
	log.Tracef("%s %s", action.Method, req.URL)
	if res, err := client.Do(req); err != nil {
		return vars, errors.Unexpected(err.Error())
	} else if res != nil {
		log.Tracef("Status code is %d", res.StatusCode)
		defer res.Body.Close()
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Warning(err.Error())
		}
		vars.SetBody(bodyBytes)
		vars.SetHeaders(res.Header)
		vars.SetStatusCode(res.StatusCode)
		bodyString := string(bodyBytes)
		bodyStrLen := math.Min(200, float64(len(bodyString)))
		log.Tracef("%s...", bodyString[:int(bodyStrLen)])
	}
	return vars, nil
}

func (action *HTTP) setCookies(client *http.Client, url *url.URL) {
	if len(action.Cookies) == 0 {
		return
	}

	cookies := make([]*http.Cookie, len(action.Cookies))
	for i := range action.Cookies {
		cookie := action.Cookies[i]
		cookies = append(cookies, &http.Cookie{
			Name:   cookie.Name,
			Value:  cookie.Value,
			Path:   cookie.Path,
			Domain: cookie.Domain,
		})
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Warnf("error creating cookie jar: %s", err.Error())
		return
	}
	jar.SetCookies(url, cookies)
	client.Jar = jar
}

func (action *HTTP) buildHTTPClient() *http.Client {
	timeout := httpClient.Timeout
	if action.Connection != nil {
		timeout = action.Connection.Timeout
	}
	transport := &http.Transport{}
	if action.Connection != nil {
		if action.Connection.Proxy != "" {
			proxyUrl, _ := url.Parse(action.Connection.Proxy)
			if proxyUrl != nil {
				transport.Proxy = http.ProxyURL(proxyUrl)
			}
		}
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	return client
}
