<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	"github.com/my-company/livepeerai"
	"github.com/my-company/livepeerai/models/components"
	"log"
)

func main() {
	s := livepeerai.New(
		livepeerai.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
	)
	request := components.TextToImageParams{
		Prompt: "<value>",
	}
	ctx := context.Background()
	res, err := s.TextToImage(ctx, request)
	if err != nil {
		log.Fatal(err)
	}
	if res.ImageResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->