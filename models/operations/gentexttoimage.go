// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-ai-go/models/components"
)

type GenTextToImageResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ImageResponse *components.ImageResponse
}

func (o *GenTextToImageResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenTextToImageResponse) GetImageResponse() *components.ImageResponse {
	if o == nil {
		return nil
	}
	return o.ImageResponse
}