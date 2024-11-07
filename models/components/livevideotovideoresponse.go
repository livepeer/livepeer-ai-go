// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

// LiveVideoToVideoResponse - Response model for live video-to-video generation.
type LiveVideoToVideoResponse struct {
	// Source URL of the incoming stream to subscribe to
	SubscribeURL string `json:"subscribe_url"`
	// Destination URL of the outgoing stream to publish to
	PublishURL string `json:"publish_url"`
}

func (o *LiveVideoToVideoResponse) GetSubscribeURL() string {
	if o == nil {
		return ""
	}
	return o.SubscribeURL
}

func (o *LiveVideoToVideoResponse) GetPublishURL() string {
	if o == nil {
		return ""
	}
	return o.PublishURL
}
