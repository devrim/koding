package j_reward

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new j reward API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for j reward API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
PostRemoteAPIJRewardAddCustomReward post remote API j reward add custom reward API
*/
func (a *Client) PostRemoteAPIJRewardAddCustomReward(params *PostRemoteAPIJRewardAddCustomRewardParams) (*PostRemoteAPIJRewardAddCustomRewardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJRewardAddCustomRewardParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJRewardAddCustomReward",
		Method:             "POST",
		PathPattern:        "/remote.api/JReward.addCustomReward",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJRewardAddCustomRewardReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJRewardAddCustomRewardOK), nil

}

/*
PostRemoteAPIJRewardFetchCustomData post remote API j reward fetch custom data API
*/
func (a *Client) PostRemoteAPIJRewardFetchCustomData(params *PostRemoteAPIJRewardFetchCustomDataParams) (*PostRemoteAPIJRewardFetchCustomDataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJRewardFetchCustomDataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJRewardFetchCustomData",
		Method:             "POST",
		PathPattern:        "/remote.api/JReward.fetchCustomData",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJRewardFetchCustomDataReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJRewardFetchCustomDataOK), nil

}

/*
PostRemoteAPIJRewardFetchEarnedAmount Method JReward.fetchEarnedAmount
*/
func (a *Client) PostRemoteAPIJRewardFetchEarnedAmount(params *PostRemoteAPIJRewardFetchEarnedAmountParams) (*PostRemoteAPIJRewardFetchEarnedAmountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJRewardFetchEarnedAmountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJRewardFetchEarnedAmount",
		Method:             "POST",
		PathPattern:        "/remote.api/JReward.fetchEarnedAmount",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJRewardFetchEarnedAmountReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJRewardFetchEarnedAmountOK), nil

}

/*
PostRemoteAPIJRewardSome post remote API j reward some API
*/
func (a *Client) PostRemoteAPIJRewardSome(params *PostRemoteAPIJRewardSomeParams) (*PostRemoteAPIJRewardSomeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostRemoteAPIJRewardSomeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostRemoteAPIJRewardSome",
		Method:             "POST",
		PathPattern:        "/remote.api/JReward.some",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostRemoteAPIJRewardSomeReader{formats: a.formats},
		Context:            params.Context,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostRemoteAPIJRewardSomeOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
