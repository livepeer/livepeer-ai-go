// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-ai-go/models/components"
)

type GenSegmentAnything2Response struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	MasksResponse *components.MasksResponse
}

func (o *GenSegmentAnything2Response) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenSegmentAnything2Response) GetMasksResponse() *components.MasksResponse {
	if o == nil {
		return nil
	}
	return o.MasksResponse
}
