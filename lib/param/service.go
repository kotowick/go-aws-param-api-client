// Package param provides a client for ParamClient.
//
package param

// Config struct
//
// Config makes it easy for you to make a HTTP request to ParamApi Services
// and to parse the response. It will either set the environment variables using
// os.Setenv(), or outputing them to a file.
//
type Config struct {
	Request    Request
	Response   map[string]string
	ServiceURL string
}

// Request struct
//
// The variables that are needed to be passed in a POST request to the param-api
// service. This is marshalled into JSON
//
type Request struct {
	Application string
	Environment string
	Version     string
	Landscape   string
}

// New creates a new instance of the Param client.
//
// Example:
//     // Create a Param client.
//     svc := param.New(landscape, application, environment, version)
//
func New(landscape, application, environment, version, serviceURL string) *Config {
	return newClient(landscape, application, environment, version, serviceURL)
}

// newClient creates, initializes and returns a new Param service.
//
func newClient(landscape, application, environment, version, serviceURL string) *Config {
	svc := &Config{
		Request:    Request{Application: application, Environment: environment, Version: version, Landscape: landscape},
		ServiceURL: serviceURL,
	}

	return svc
}
