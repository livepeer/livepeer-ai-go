// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/livepeer/livepeer-ai-go/internal/utils"
)

type Image struct {
	FileName string `multipartForm:"name=image"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *Image) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *Image) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type BodyGenImageToImage struct {
	// Text prompt(s) to guide image generation.
	Prompt string `multipartForm:"name=prompt"`
	// Uploaded image to modify with the pipeline.
	Image Image `multipartForm:"file"`
	// Hugging Face model ID used for image generation.
	ModelID *string `default:"" multipartForm:"name=model_id"`
	// Degree of transformation applied to the reference image (0 to 1).
	Strength *float64 `default:"0.8" multipartForm:"name=strength"`
	// Encourages model to generate images closely linked to the text prompt (higher values may reduce image quality).
	GuidanceScale *float64 `default:"7.5" multipartForm:"name=guidance_scale"`
	// Degree to which the generated image is pushed towards the initial image.
	ImageGuidanceScale *float64 `default:"1.5" multipartForm:"name=image_guidance_scale"`
	// Text prompt(s) to guide what to exclude from image generation. Ignored if guidance_scale < 1.
	NegativePrompt *string `default:"" multipartForm:"name=negative_prompt"`
	// Perform a safety check to estimate if generated images could be offensive or harmful.
	SafetyCheck *bool `default:"true" multipartForm:"name=safety_check"`
	// Seed for random number generation.
	Seed *int64 `multipartForm:"name=seed"`
	// Number of denoising steps. More steps usually lead to higher quality images but slower inference. Modulated by strength.
	NumInferenceSteps *int64 `default:"100" multipartForm:"name=num_inference_steps"`
	// Number of images to generate per prompt.
	NumImagesPerPrompt *int64 `default:"1" multipartForm:"name=num_images_per_prompt"`
}

func (b BodyGenImageToImage) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BodyGenImageToImage) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BodyGenImageToImage) GetPrompt() string {
	if o == nil {
		return ""
	}
	return o.Prompt
}

func (o *BodyGenImageToImage) GetImage() Image {
	if o == nil {
		return Image{}
	}
	return o.Image
}

func (o *BodyGenImageToImage) GetModelID() *string {
	if o == nil {
		return nil
	}
	return o.ModelID
}

func (o *BodyGenImageToImage) GetStrength() *float64 {
	if o == nil {
		return nil
	}
	return o.Strength
}

func (o *BodyGenImageToImage) GetGuidanceScale() *float64 {
	if o == nil {
		return nil
	}
	return o.GuidanceScale
}

func (o *BodyGenImageToImage) GetImageGuidanceScale() *float64 {
	if o == nil {
		return nil
	}
	return o.ImageGuidanceScale
}

func (o *BodyGenImageToImage) GetNegativePrompt() *string {
	if o == nil {
		return nil
	}
	return o.NegativePrompt
}

func (o *BodyGenImageToImage) GetSafetyCheck() *bool {
	if o == nil {
		return nil
	}
	return o.SafetyCheck
}

func (o *BodyGenImageToImage) GetSeed() *int64 {
	if o == nil {
		return nil
	}
	return o.Seed
}

func (o *BodyGenImageToImage) GetNumInferenceSteps() *int64 {
	if o == nil {
		return nil
	}
	return o.NumInferenceSteps
}

func (o *BodyGenImageToImage) GetNumImagesPerPrompt() *int64 {
	if o == nil {
		return nil
	}
	return o.NumImagesPerPrompt
}