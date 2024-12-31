// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"github.com/livepeer/livepeer-ai-go/internal/utils"
)

type BodyGenImageToTextImage struct {
	FileName string `multipartForm:"name=fileName"`
	// This field accepts []byte data or io.Reader implementations, such as *os.File.
	Content any `multipartForm:"content"`
}

func (o *BodyGenImageToTextImage) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

func (o *BodyGenImageToTextImage) GetContent() any {
	if o == nil {
		return nil
	}
	return o.Content
}

type BodyGenImageToText struct {
	// Uploaded image to transform with the pipeline.
	Image BodyGenImageToTextImage `multipartForm:"file,name=image"`
	// Text prompt(s) to guide transformation.
	Prompt *string `default:"" multipartForm:"name=prompt"`
	// Hugging Face model ID used for transformation.
	ModelID *string `default:"" multipartForm:"name=model_id"`
}

func (b BodyGenImageToText) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BodyGenImageToText) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *BodyGenImageToText) GetImage() BodyGenImageToTextImage {
	if o == nil {
		return BodyGenImageToTextImage{}
	}
	return o.Image
}

func (o *BodyGenImageToText) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *BodyGenImageToText) GetModelID() *string {
	if o == nil {
		return nil
	}
	return o.ModelID
}
