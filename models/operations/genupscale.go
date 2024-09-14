// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-ai-go/models/components"
)

type GenUpscaleResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ImageResponse *components.ImageResponse
}

func (o *GenUpscaleResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenUpscaleResponse) GetImageResponse() *components.ImageResponse {
	if o == nil {
		return nil
	}
	return o.ImageResponse
}