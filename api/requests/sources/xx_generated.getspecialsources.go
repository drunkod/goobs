// This file has been automatically generated. Don't edit it.

package sources

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSpecialSourcesParams represents the params body for the "GetSpecialSources" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesParams struct {
	requests.Params
}

// GetRequestType returns the RequestType of GetSpecialSourcesParams
func (o *GetSpecialSourcesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSpecialSourcesParams
func (o *GetSpecialSourcesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSpecialSourcesParams
func (o *GetSpecialSourcesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSpecialSourcesResponse represents the response body for the "GetSpecialSources" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSpecialSources.
*/
type GetSpecialSourcesResponse struct {
	requests.Response

	// Name of the first Desktop Audio capture source.
	Desktop1 string `json:"desktop-1"`

	// Name of the second Desktop Audio capture source.
	Desktop2 string `json:"desktop-2"`

	// Name of the first Mic/Aux input source.
	Mic1 string `json:"mic-1"`

	// Name of the second Mic/Aux input source.
	Mic2 string `json:"mic-2"`

	// NAme of the third Mic/Aux input source.
	Mic3 string `json:"mic-3"`
}

// GetMessageID returns the MessageID of GetSpecialSourcesResponse
func (o *GetSpecialSourcesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSpecialSourcesResponse
func (o *GetSpecialSourcesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSpecialSourcesResponse
func (o *GetSpecialSourcesResponse) GetError() string {
	return o.Error
}

// GetSpecialSources sends the corresponding request to the connected OBS WebSockets server. Note
// the variadic arguments as this request doesn't require any parameters.
func (c *Client) GetSpecialSources(
	paramss ...*GetSpecialSourcesParams,
) (*GetSpecialSourcesResponse, error) {
	if len(paramss) == 0 {
		paramss = []*GetSpecialSourcesParams{{}}
	}
	params := paramss[0]
	params.RequestType = "GetSpecialSources"
	data := &GetSpecialSourcesResponse{}
	if err := requests.WriteMessage(c.Conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
