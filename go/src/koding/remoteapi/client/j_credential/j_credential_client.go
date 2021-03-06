package j_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j credential API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j credential API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJCredentialCloneID post remote API j credential clone ID API
*/
func (a *Client) PostRemoteAPIJCredentialCloneID(params *PostRemoteAPIJCredentialCloneIDParams) (*PostRemoteAPIJCredentialCloneIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialCloneIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialCloneID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.clone/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialCloneIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialCloneIDOK), nil

}

/*
PostRemoteAPIJCredentialCreate post remote API j credential create API
*/
func (a *Client) PostRemoteAPIJCredentialCreate(params *PostRemoteAPIJCredentialCreateParams) (*PostRemoteAPIJCredentialCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialCreate",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.create",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialCreateReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialCreateOK), nil

}

/*
PostRemoteAPIJCredentialDeleteID post remote API j credential delete ID API
*/
func (a *Client) PostRemoteAPIJCredentialDeleteID(params *PostRemoteAPIJCredentialDeleteIDParams) (*PostRemoteAPIJCredentialDeleteIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialDeleteIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialDeleteID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.delete/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialDeleteIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialDeleteIDOK), nil

}

/*
PostRemoteAPIJCredentialFetchDataID Method JCredential.fetchData
*/
func (a *Client) PostRemoteAPIJCredentialFetchDataID(params *PostRemoteAPIJCredentialFetchDataIDParams) (*PostRemoteAPIJCredentialFetchDataIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialFetchDataIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialFetchDataID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.fetchData/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialFetchDataIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialFetchDataIDOK), nil

}

/*
PostRemoteAPIJCredentialFetchUsersID post remote API j credential fetch users ID API
*/
func (a *Client) PostRemoteAPIJCredentialFetchUsersID(params *PostRemoteAPIJCredentialFetchUsersIDParams) (*PostRemoteAPIJCredentialFetchUsersIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialFetchUsersIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialFetchUsersID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.fetchUsers/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialFetchUsersIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialFetchUsersIDOK), nil

}

/*
PostRemoteAPIJCredentialIsBootstrappedID post remote API j credential is bootstrapped ID API
*/
func (a *Client) PostRemoteAPIJCredentialIsBootstrappedID(params *PostRemoteAPIJCredentialIsBootstrappedIDParams) (*PostRemoteAPIJCredentialIsBootstrappedIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialIsBootstrappedIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialIsBootstrappedID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.isBootstrapped/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialIsBootstrappedIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialIsBootstrappedIDOK), nil

}

/*
PostRemoteAPIJCredentialOne post remote API j credential one API
*/
func (a *Client) PostRemoteAPIJCredentialOne(params *PostRemoteAPIJCredentialOneParams) (*PostRemoteAPIJCredentialOneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialOneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialOne",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.one",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialOneReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialOneOK), nil

}

/*
PostRemoteAPIJCredentialShareWithID Method JCredential.shareWith
*/
func (a *Client) PostRemoteAPIJCredentialShareWithID(params *PostRemoteAPIJCredentialShareWithIDParams) (*PostRemoteAPIJCredentialShareWithIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialShareWithIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialShareWithID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.shareWith/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialShareWithIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialShareWithIDOK), nil

}

/*
PostRemoteAPIJCredentialSome post remote API j credential some API
*/
func (a *Client) PostRemoteAPIJCredentialSome(params *PostRemoteAPIJCredentialSomeParams) (*PostRemoteAPIJCredentialSomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialSomeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialSome",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.some",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialSomeReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialSomeOK), nil

}

/*
PostRemoteAPIJCredentialUpdateID post remote API j credential update ID API
*/
func (a *Client) PostRemoteAPIJCredentialUpdateID(params *PostRemoteAPIJCredentialUpdateIDParams) (*PostRemoteAPIJCredentialUpdateIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJCredentialUpdateIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJCredentialUpdateID",
		Method:             "POST",
		PathPattern:        "/remote.api/JCredential.update/{id}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJCredentialUpdateIDReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJCredentialUpdateIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
