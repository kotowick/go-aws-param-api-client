// Package main is the main client for thr param-api client.
package main

import (
	"./lib/param"
	"gopkg.in/alecthomas/kingpin.v2"
)

//
type paramRequest struct {
	Application string
	Environment string
	Version     string
	Landscape   string
}

var (
	// General Flags
	landscape   = kingpin.Flag("landscape", "Landscape").Short('l').Envar("PARAM_API_LANDSCAPE").String()
	environment = kingpin.Flag("environment", "Environment").Short('e').Envar("PARAM_API_ENVIRONMENT").String()
	application = kingpin.Flag("app", "Application").Short('a').Envar("PARAM_API_APPLICATION").String()
	version     = kingpin.Flag("env-version", "Version").Short('v').Envar("PARAM_API_VERSION").String()
	prefix      = kingpin.Flag("prefix", "Prefix").Short('p').Envar("PARAM_API_PREFIX").String()
	output      = kingpin.Flag("output", "Output").Short('o').Envar("PARAM_API_OUTPUT_FILE").String()

	//Global Vars
	cliVersion = "0.0.1"
)

//	Entry Controller
func main() {
	kingpin.Version(cliVersion)

	kingpin.Parse()

	paramConfig := param.New(
		*landscape,
		*application,
		*environment,
		*version,
	)

	if val := paramConfig.Verify(); !val {
		return
	}

	httpResp, err := paramConfig.RequestHandler(*prefix)

	if err != nil {
		return
	}

	paramConfig.Response = httpResp

	paramConfig.ResponseHandler(*prefix)

	return
}
