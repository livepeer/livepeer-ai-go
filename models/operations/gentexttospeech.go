// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/livepeer/livepeer-ai-go/models/components"
)

type GenTextToSpeechResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Successful Response
	AudioResponse *components.AudioResponse
}

func (o *GenTextToSpeechResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GenTextToSpeechResponse) GetAudioResponse() *components.AudioResponse {
	if o == nil {
		return nil
	}
	return o.AudioResponse
}
