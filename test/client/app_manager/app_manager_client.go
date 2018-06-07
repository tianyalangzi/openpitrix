// Code generated by go-swagger; DO NOT EDIT.

package app_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new app manager API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for app manager API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateApp creates app
*/
func (a *Client) CreateApp(params *CreateAppParams) (*CreateAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateApp",
		Method:             "POST",
		PathPattern:        "/v1/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAppReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAppOK), nil

}

/*
CreateAppVersion creates app version
*/
func (a *Client) CreateAppVersion(params *CreateAppVersionParams) (*CreateAppVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateAppVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateAppVersion",
		Method:             "POST",
		PathPattern:        "/v1/app_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateAppVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateAppVersionOK), nil

}

/*
DeleteAppVersions deletes app versions
*/
func (a *Client) DeleteAppVersions(params *DeleteAppVersionsParams) (*DeleteAppVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAppVersions",
		Method:             "DELETE",
		PathPattern:        "/v1/app_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAppVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppVersionsOK), nil

}

/*
DeleteApps deletes apps
*/
func (a *Client) DeleteApps(params *DeleteAppsParams) (*DeleteAppsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAppsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteApps",
		Method:             "DELETE",
		PathPattern:        "/v1/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAppsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAppsOK), nil

}

/*
DescribeAppVersions describes app versions with filter
*/
func (a *Client) DescribeAppVersions(params *DescribeAppVersionsParams) (*DescribeAppVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeAppVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeAppVersions",
		Method:             "GET",
		PathPattern:        "/v1/app_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeAppVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeAppVersionsOK), nil

}

/*
DescribeApps describes apps with filter
*/
func (a *Client) DescribeApps(params *DescribeAppsParams) (*DescribeAppsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDescribeAppsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DescribeApps",
		Method:             "GET",
		PathPattern:        "/v1/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DescribeAppsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DescribeAppsOK), nil

}

/*
GetAppVersionPackage gets the package content of app version
*/
func (a *Client) GetAppVersionPackage(params *GetAppVersionPackageParams) (*GetAppVersionPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppVersionPackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppVersionPackage",
		Method:             "GET",
		PathPattern:        "/v1/app_version/package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAppVersionPackageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppVersionPackageOK), nil

}

/*
GetAppVersionPackageFiles gets the package files content of app version
*/
func (a *Client) GetAppVersionPackageFiles(params *GetAppVersionPackageFilesParams) (*GetAppVersionPackageFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAppVersionPackageFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAppVersionPackageFiles",
		Method:             "GET",
		PathPattern:        "/v1/app_version/package/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAppVersionPackageFilesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAppVersionPackageFilesOK), nil

}

/*
ModifyApp modifies app
*/
func (a *Client) ModifyApp(params *ModifyAppParams) (*ModifyAppOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyAppParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyApp",
		Method:             "PATCH",
		PathPattern:        "/v1/apps",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyAppReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyAppOK), nil

}

/*
ModifyAppVersion modifies app version
*/
func (a *Client) ModifyAppVersion(params *ModifyAppVersionParams) (*ModifyAppVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyAppVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyAppVersion",
		Method:             "PATCH",
		PathPattern:        "/v1/app_versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyAppVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyAppVersionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
