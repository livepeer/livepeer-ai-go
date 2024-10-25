// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-ai-go/models/components"
)

type GenImageToTextResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	ImageToTextResponse *components.ImageToTextResponse
}

func (o *GenImageToTextResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenImageToTextResponse) GetImageToTextResponse() *components.ImageToTextResponse {
	if o == nil {
		return nil
	}
	return o.ImageToTextResponse
}
