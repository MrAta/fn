package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/funcy/functions_go/client/apps"
	"github.com/funcy/functions_go/client/call"
	"github.com/funcy/functions_go/client/routes"
	"github.com/funcy/functions_go/client/tasks"
	"github.com/funcy/functions_go/client/version"
)

// Default functions HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "127.0.0.1:8080"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new functions HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Functions {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new functions HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Functions {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new functions client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Functions {
	cli := new(Functions)
	cli.Transport = transport

	cli.Apps = apps.New(transport, formats)

	cli.Call = call.New(transport, formats)

	cli.Routes = routes.New(transport, formats)

	cli.Tasks = tasks.New(transport, formats)

	cli.Version = version.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Functions is a client for functions
type Functions struct {
	Apps *apps.Client

	Call *call.Client

	Routes *routes.Client

	Tasks *tasks.Client

	Version *version.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Functions) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Apps.SetTransport(transport)

	c.Call.SetTransport(transport)

	c.Routes.SetTransport(transport)

	c.Tasks.SetTransport(transport)

	c.Version.SetTransport(transport)

}
