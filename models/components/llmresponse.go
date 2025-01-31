// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LLMResponse struct {
	ID      string        `json:"id"`
	Model   string        `json:"model"`
	Created int64         `json:"created"`
	Usage   LLMTokenUsage `json:"usage"`
	Choices []LLMChoice   `json:"choices"`
}

func (o *LLMResponse) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LLMResponse) GetModel() string {
	if o == nil {
		return ""
	}
	return o.Model
}

func (o *LLMResponse) GetCreated() int64 {
	if o == nil {
		return 0
	}
	return o.Created
}

func (o *LLMResponse) GetUsage() LLMTokenUsage {
	if o == nil {
		return LLMTokenUsage{}
	}
	return o.Usage
}

func (o *LLMResponse) GetChoices() []LLMChoice {
	if o == nil {
		return []LLMChoice{}
	}
	return o.Choices
}
