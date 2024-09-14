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
	s := livepeeraigo.New(
		livepeeraigo.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)

	ctx := context.Background()
	res, err := s.Generate.TextToImage(ctx, components.TextToImageParams{
		Prompt: "<value>",
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