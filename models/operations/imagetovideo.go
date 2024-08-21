// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/my-company/livepeerai/models/components"
)

type ImageToVideoResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	VideoResponse *components.VideoResponse
}

func (o *ImageToVideoResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ImageToVideoResponse) GetVideoResponse() *components.VideoResponse {
	if o == nil {
		return nil
	}
	return o.VideoResponse
}
