// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/my-company/livepeerai/internal/utils"
)

type BodyImageToVideoImageToVideoPostImage struct {
	FileName string `multipartForm:"name=image"`
	Content  []byte `multipartForm:"content"`
}

func (o *BodyImageToVideoImageToVideoPostImage) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *BodyImageToVideoImageToVideoPostImage) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

type BodyImageToVideoImageToVideoPost struct {
	// Uploaded image to generate a video from.
	Image BodyImageToVideoImageToVideoPostImage `multipartForm:"file"`
	// Hugging Face model ID used for video generation.
	ModelID *string `default:"" multipartForm:"name=model_id"`
	// The height in pixels of the generated video.
	Height *int64 `default:"576" multipartForm:"name=height"`
	// The width in pixels of the generated video.
	Width *int64 `default:"1024" multipartForm:"name=width"`
	// The frames per second of the generated video.
	Fps *int64 `default:"6" multipartForm:"name=fps"`
	// Used for conditioning the amount of motion for the generation. The higher the number the more motion will be in the video.
	MotionBucketID *int64 `default:"127" multipartForm:"name=motion_bucket_id"`
	// Amount of noise added to the conditioning image. Higher values reduce resemblance to the conditioning image and increase motion.
	NoiseAugStrength *float64 `default:"0.02" multipartForm:"name=noise_aug_strength"`
	// Perform a safety check to estimate if generated images could be offensive or harmful.
	SafetyCheck *bool `default:"true" multipartForm:"name=safety_check"`
	// Seed for random number generation.
	Seed *int64 `multipartForm:"name=seed"`
	// Number of denoising steps. More steps usually lead to higher quality images but slower inference. Modulated by strength.
	NumInferenceSteps *int64 `default:"25" multipartForm:"name=num_inference_steps"`
}

func (b BodyImageToVideoImageToVideoPost) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BodyImageToVideoImageToVideoPost) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BodyImageToVideoImageToVideoPost) GetImage() BodyImageToVideoImageToVideoPostImage {
	if o == nil {
		return BodyImageToVideoImageToVideoPostImage{}
	}
	return o.Image
}

func (o *BodyImageToVideoImageToVideoPost) GetModelID() *string {
	if o == nil {
		return nil
	}
	return o.ModelID
}

func (o *BodyImageToVideoImageToVideoPost) GetHeight() *int64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *BodyImageToVideoImageToVideoPost) GetWidth() *int64 {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *BodyImageToVideoImageToVideoPost) GetFps() *int64 {
	if o == nil {
		return nil
	}
	return o.Fps
}

func (o *BodyImageToVideoImageToVideoPost) GetMotionBucketID() *int64 {
	if o == nil {
		return nil
	}
	return o.MotionBucketID
}

func (o *BodyImageToVideoImageToVideoPost) GetNoiseAugStrength() *float64 {
	if o == nil {
		return nil
	}
	return o.NoiseAugStrength
}

func (o *BodyImageToVideoImageToVideoPost) GetSafetyCheck() *bool {
	if o == nil {
		return nil
	}
	return o.SafetyCheck
}

func (o *BodyImageToVideoImageToVideoPost) GetSeed() *int64 {
	if o == nil {
		return nil
	}
	return o.Seed
}

func (o *BodyImageToVideoImageToVideoPost) GetNumInferenceSteps() *int64 {
	if o == nil {
		return nil
	}
	return o.NumInferenceSteps
}
