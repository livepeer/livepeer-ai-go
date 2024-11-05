# LiveVideoToVideoParams


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `SubscribeURL`                                          | *string*                                                | :heavy_check_mark:                                      | Source URL of the incoming stream to subscribe to.      |
| `PublishURL`                                            | *string*                                                | :heavy_check_mark:                                      | Destination URL of the outgoing stream to publish.      |
| `ModelID`                                               | **string*                                               | :heavy_minus_sign:                                      | Hugging Face model ID used for image generation.        |
| `Params`                                                | [*components.Params](../../models/components/params.md) | :heavy_minus_sign:                                      | Initial parameters for the model.                       |