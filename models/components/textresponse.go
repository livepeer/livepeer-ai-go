// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type TextResponse struct {
	Text   string  `json:"text"`
	Chunks []Chunk `json:"chunks"`
}

func (o *TextResponse) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

func (o *TextResponse) GetChunks() []Chunk {
	if o == nil {
		return []Chunk{}
	}
	return o.Chunks
}
