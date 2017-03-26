// Package param provides a client for ParamClient.
//
package param

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"../utils"
)

// Verify creates a new instance of the Param client.
//
// Example:
//     // Verify configuration.
//     svc := param.Verify
//
func (p Config) Verify() bool {
	if !utils.EvalBool(utils.IsSet(p.Request.Landscape), "Landscape") ||
		!utils.EvalBool(utils.IsSet(p.Request.Environment), "Environment") ||
		!utils.EvalBool(utils.IsSet(p.Request.Application), "Application") ||
		!utils.EvalBool(utils.IsSet(p.Request.Version), "Version") {
		return false
	}
	return true
}

// RequestHandler makes a HTTP Post request to the param- api service. It
// marshells the param.Request and sends it to the service. It retrieve a json
// formatted response.
//
// Example:
//     // Make a HTTP Post request
//     svc := param.RequestHandler(prefix)
//
func (p Config) RequestHandler(prefix string) (str map[string]string, err error) {
	jsonValue, err := json.Marshal(p.Request)

	if err != nil {
		return p.Response, err
	}

	res, err := http.Post("http://localhost:8080/params", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {
		return p.Response, err
	}

	json.NewDecoder(res.Body).Decode(&p.Response)
	return p.Response, nil
}

// ResponseHandler parses the response from the RequestHandler.
// It will parse through the environment variables and set then using os.Setenv
// or by outputting them to a file.
//
// Example:
//     // Parse the response
//     svc := param.ResponseHandler(prefix)
//
func (p Config) ResponseHandler(prefix string) {
	for k := range p.Response {
		if prefix != "" {
			if strings.HasPrefix(k, prefix) {
				fmt.Printf("Setting ENV var %s : %s\n", k, p.Response[k])
				os.Setenv(k, p.Response[k])
			}
		} else {
			fmt.Printf("Setting ENV var %s : %s\n", k, p.Response[k])
			os.Setenv(k, p.Response[k])
		}
	}
}
