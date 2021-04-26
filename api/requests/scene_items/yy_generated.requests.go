// This file has been automatically generated. Don't edit it.

package sceneitems

import requests "github.com/andreykaipov/goobs/api/requests"

/*
GetSceneItemPropertiesParams represents the params body for the "GetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesParams struct {
	requests.Params

	// The name of the source.
	Item string `json:"item"`

	// the name of the scene that the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of GetSceneItemPropertiesParams
func (o *GetSceneItemPropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of GetSceneItemPropertiesParams
func (o *GetSceneItemPropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on GetSceneItemPropertiesParams
func (o *GetSceneItemPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
GetSceneItemPropertiesResponse represents the response body for the "GetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#GetSceneItemProperties.
*/
type GetSceneItemPropertiesResponse struct {
	requests.Response

	Bounds struct {
		// Alignment of the bounding box.
		Alignment int `json:"alignment"`

		// Type of bounding box.
		Type string `json:"type"`

		// Width of the bounding box.
		X float64 `json:"x"`

		// Height of the bounding box.
		Y float64 `json:"y"`
	} `json:"bounds"`

	Crop struct {
		// The number of pixels cropped off the bottom of the source before scaling.
		Bottom int `json:"bottom"`

		// The number of pixels cropped off the left of the source before scaling.
		Left int `json:"left"`

		// The number of pixels cropped off the right of the source before scaling.
		Right int `json:"right"`

		// The number of pixels cropped off the top of the source before scaling.
		Top int `json:"top"`
	} `json:"crop"`

	// The name of the source.
	Name string `json:"name"`

	Position struct {
		// The point on the source that the item is manipulated from.
		Alignment float64 `json:"alignment"`

		// The x position of the source from the left.
		X float64 `json:"x"`

		// The y position of the source from the top.
		Y float64 `json:"y"`
	} `json:"position"`

	// The clockwise rotation of the item in degrees around the point of alignment.
	Rotation float64 `json:"rotation"`

	Scale struct {
		// The x-scale factor of the source.
		X float64 `json:"x"`

		// The y-scale factor of the source.
		Y float64 `json:"y"`
	} `json:"scale"`

	// If the source is visible.
	Visible bool `json:"visible"`
}

// GetMessageID returns the MessageID of GetSceneItemPropertiesResponse
func (o *GetSceneItemPropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of GetSceneItemPropertiesResponse
func (o *GetSceneItemPropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of GetSceneItemPropertiesResponse
func (o *GetSceneItemPropertiesResponse) GetError() string {
	return o.Error
}

// GetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) GetSceneItemProperties(
	params *GetSceneItemPropertiesParams,
) (*GetSceneItemPropertiesResponse, error) {
	params.RequestType = "GetSceneItemProperties"
	data := &GetSceneItemPropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
SetSceneItemPropertiesParams represents the params body for the "SetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSceneItemProperties.
*/
type SetSceneItemPropertiesParams struct {
	requests.Params

	Bounds struct {
		// The new alignment of the bounding box. (0-2, 4-6, 8-10)
		Alignment int `json:"alignment"`

		// The new bounds type of the source.
		Type string `json:"type"`

		// The new width of the bounding box.
		X float64 `json:"x"`

		// The new height of the bounding box.
		Y float64 `json:"y"`
	} `json:"bounds"`

	Crop struct {
		// The new amount of pixels cropped off the bottom of the source before scaling.
		Bottom int `json:"bottom"`

		// The new amount of pixels cropped off the left of the source before scaling.
		Left int `json:"left"`

		// The new amount of pixels cropped off the right of the source before scaling.
		Right int `json:"right"`

		// The new amount of pixels cropped off the top of the source before scaling.
		Top int `json:"top"`
	} `json:"crop"`

	// The name of the source.
	Item string `json:"item"`

	Position struct {
		// The new alignment of the source.
		Alignment float64 `json:"alignment"`

		// The new x position of the source.
		X float64 `json:"x"`

		// The new y position of the source.
		Y float64 `json:"y"`
	} `json:"position"`

	// The new clockwise rotation of the item in degrees.
	Rotation float64 `json:"rotation"`

	Scale struct {
		// The new x scale of the item.
		X float64 `json:"x"`

		// The new y scale of the item.
		Y float64 `json:"y"`
	} `json:"scale"`

	// the name of the scene that the source item belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`

	// The new visibility of the source. 'true' shows source, 'false' hides source.
	Visible bool `json:"visible"`
}

// GetRequestType returns the RequestType of SetSceneItemPropertiesParams
func (o *SetSceneItemPropertiesParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of SetSceneItemPropertiesParams
func (o *SetSceneItemPropertiesParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on SetSceneItemPropertiesParams
func (o *SetSceneItemPropertiesParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
SetSceneItemPropertiesResponse represents the response body for the "SetSceneItemProperties" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#SetSceneItemProperties.
*/
type SetSceneItemPropertiesResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of SetSceneItemPropertiesResponse
func (o *SetSceneItemPropertiesResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of SetSceneItemPropertiesResponse
func (o *SetSceneItemPropertiesResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of SetSceneItemPropertiesResponse
func (o *SetSceneItemPropertiesResponse) GetError() string {
	return o.Error
}

// SetSceneItemProperties sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) SetSceneItemProperties(
	params *SetSceneItemPropertiesParams,
) (*SetSceneItemPropertiesResponse, error) {
	params.RequestType = "SetSceneItemProperties"
	data := &SetSceneItemPropertiesResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
ResetSceneItemParams represents the params body for the "ResetSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem.
*/
type ResetSceneItemParams struct {
	requests.Params

	// Name of the source item.
	Item string `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	SceneName string `json:"scene-name"`
}

// GetRequestType returns the RequestType of ResetSceneItemParams
func (o *ResetSceneItemParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of ResetSceneItemParams
func (o *ResetSceneItemParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on ResetSceneItemParams
func (o *ResetSceneItemParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
ResetSceneItemResponse represents the response body for the "ResetSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#ResetSceneItem.
*/
type ResetSceneItemResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of ResetSceneItemResponse
func (o *ResetSceneItemResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of ResetSceneItemResponse
func (o *ResetSceneItemResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of ResetSceneItemResponse
func (o *ResetSceneItemResponse) GetError() string {
	return o.Error
}

// ResetSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) ResetSceneItem(params *ResetSceneItemParams) (*ResetSceneItemResponse, error) {
	params.RequestType = "ResetSceneItem"
	data := &ResetSceneItemResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
DeleteSceneItemParams represents the params body for the "DeleteSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DeleteSceneItem.
*/
type DeleteSceneItemParams struct {
	requests.Params

	Item struct {
		// id of the scene item.
		Id int `json:"id"`

		// name of the scene item (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene the source belongs to. Defaults to the current scene.
	Scene string `json:"scene"`
}

// GetRequestType returns the RequestType of DeleteSceneItemParams
func (o *DeleteSceneItemParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of DeleteSceneItemParams
func (o *DeleteSceneItemParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on DeleteSceneItemParams
func (o *DeleteSceneItemParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DeleteSceneItemResponse represents the response body for the "DeleteSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DeleteSceneItem.
*/
type DeleteSceneItemResponse struct {
	requests.Response
}

// GetMessageID returns the MessageID of DeleteSceneItemResponse
func (o *DeleteSceneItemResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of DeleteSceneItemResponse
func (o *DeleteSceneItemResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of DeleteSceneItemResponse
func (o *DeleteSceneItemResponse) GetError() string {
	return o.Error
}

// DeleteSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DeleteSceneItem(params *DeleteSceneItemParams) (*DeleteSceneItemResponse, error) {
	params.RequestType = "DeleteSceneItem"
	data := &DeleteSceneItemResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}

/*
DuplicateSceneItemParams represents the params body for the "DuplicateSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DuplicateSceneItem.
*/
type DuplicateSceneItemParams struct {
	requests.Params

	// Name of the scene to copy the item from. Defaults to the current scene.
	FromScene string `json:"fromScene"`

	Item struct {
		// id of the scene item.
		Id int `json:"id"`

		// name of the scene item (prefer `id`, including both is acceptable).
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene to create the item in. Defaults to the current scene.
	ToScene string `json:"toScene"`
}

// GetRequestType returns the RequestType of DuplicateSceneItemParams
func (o *DuplicateSceneItemParams) GetRequestType() string {
	return o.RequestType
}

// GetMessageID returns the MessageID of DuplicateSceneItemParams
func (o *DuplicateSceneItemParams) GetMessageID() string {
	return o.MessageID
}

// SetMessageID sets the MessageID on DuplicateSceneItemParams
func (o *DuplicateSceneItemParams) SetMessageID(x string) {
	o.MessageID = x
}

/*
DuplicateSceneItemResponse represents the response body for the "DuplicateSceneItem" request.

Generated from https://github.com/Palakis/obs-websocket/blob/4.5.0/docs/generated/protocol.md#DuplicateSceneItem.
*/
type DuplicateSceneItemResponse struct {
	requests.Response

	Item struct {
		// New item ID
		Id int `json:"id"`

		// New item name
		Name string `json:"name"`
	} `json:"item"`

	// Name of the scene where the new item was created
	Scene string `json:"scene"`
}

// GetMessageID returns the MessageID of DuplicateSceneItemResponse
func (o *DuplicateSceneItemResponse) GetMessageID() string {
	return o.MessageID
}

// GetStatus returns the Status of DuplicateSceneItemResponse
func (o *DuplicateSceneItemResponse) GetStatus() string {
	return o.Status
}

// GetError returns the Error of DuplicateSceneItemResponse
func (o *DuplicateSceneItemResponse) GetError() string {
	return o.Error
}

// DuplicateSceneItem sends the corresponding request to the connected OBS WebSockets server.
func (c *Client) DuplicateSceneItem(
	params *DuplicateSceneItemParams,
) (*DuplicateSceneItemResponse, error) {
	params.RequestType = "DuplicateSceneItem"
	data := &DuplicateSceneItemResponse{}
	if err := requests.WriteMessage(c.conn, params, data); err != nil {
		return nil, err
	}
	return data, nil
}
