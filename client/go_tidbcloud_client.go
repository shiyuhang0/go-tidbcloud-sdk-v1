// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/c4pt0r/go-tidbcloud-sdk-v1/client/backup"
	"github.com/c4pt0r/go-tidbcloud-sdk-v1/client/cluster"
	"github.com/c4pt0r/go-tidbcloud-sdk-v1/client/project"
	"github.com/c4pt0r/go-tidbcloud-sdk-v1/client/restore"
)

// Default go tidbcloud HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.tidbcloud.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new go tidbcloud HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GoTidbcloud {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new go tidbcloud HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GoTidbcloud {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new go tidbcloud client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *GoTidbcloud {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GoTidbcloud)
	cli.Transport = transport
	cli.Backup = backup.New(transport, formats)
	cli.Cluster = cluster.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Restore = restore.New(transport, formats)
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

// GoTidbcloud is a client for go tidbcloud
type GoTidbcloud struct {
	Backup backup.ClientService

	Cluster cluster.ClientService

	Project project.ClientService

	Restore restore.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *GoTidbcloud) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Backup.SetTransport(transport)
	c.Cluster.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Restore.SetTransport(transport)
}
