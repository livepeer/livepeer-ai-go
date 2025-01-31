<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	livepeeraigo "github.com/livepeer/livepeer-ai-go"
	"github.com/livepeer/livepeer-ai-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := livepeeraigo.New(
		livepeeraigo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	res, err := s.Generate.TextToImage(ctx, components.TextToImageParams{
		ModelID:            livepeeraigo.String(""),
		Loras:              livepeeraigo.String(""),
		Prompt:             "<value>",
		Height:             livepeeraigo.Int64(576),
		Width:              livepeeraigo.Int64(1024),
		GuidanceScale:      livepeeraigo.Float64(7.5),
		NegativePrompt:     livepeeraigo.String(""),
		SafetyCheck:        livepeeraigo.Bool(true),
		NumInferenceSteps:  livepeeraigo.Int64(50),
		NumImagesPerPrompt: livepeeraigo.Int64(1),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.ImageResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->