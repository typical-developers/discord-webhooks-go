# Discord Webhooks
Send messages to webhooks on Discord. If you do not want to use discordgo specifically for webhooks, this is a lightweight library that will allow you to send them.

## Initalizing a Webhook Client
The library has two ways of initializing a webhook client. Unlike discordgo, you can also initalize a webhook client from its URL and not just from its ID and Secret.

### ID + Secret
```go
package main

import (
    webhooks "github.com/typical-developers/discord-webhooks-go"
)

func main() {
    id := "0"
    secret := "0"
    client := webhooks.NewWebhookClient(id, secret)
}
```

### URL
```go
package main

import (
    webhooks "github.com/typical-developers/discord-webhooks-go"
)

func main() {
    url := "https://discord.com/api/webhooks/0/0"
    client := webhooks.NewWebhookClientFromURL(url)
}
```

## Executing Webhooks
Messages can be sent through webhooks using the `Execute` method on the client.

### Sending Messages
```go
package main

import (
    "context"

    webhooks "github.com/typical-developers/discord-webhooks-go"
)

func main() {
    ctx := context.Background()

    url := "https://discord.com/api/webhooks/0/0"
    client := webhooks.NewWebhookClientFromURL(url)

    _, _, err := client.Execute(ctx,
        webhooks.MessagePayload{
            Content: "Hello World!",
        },
        nil,
    )
    if err != nil {
        panic(err)
    }
}
```

### Sending Messages with Files
```go
package main

import (
    "context"

    webhooks "github.com/typical-developers/discord-webhooks-go"
)

func main() {
    ctx := context.Background()

    url := "https://discord.com/api/webhooks/0/0"
    client := webhooks.NewWebhookClientFromURL(url)

	file, err := os.Open("file.png")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    _, _, err := client.Execute(ctx,
        webhooks.MessagePayload{
            Content: "Hello World!",
            Files: []webhooks.WebhookFile{
                webhooks.WebhookFile{
                    FileName: file.Name(),
                    Reader:   file,
                },
            },
        },
        nil,
    )
    if err != nil {
        panic(err)
    }
}
```

## Waiting for Response
Adding a `wait=true` query parameter will allow you to wait for the webhook to send and the server to acknowledge it. This will return a `WebhookMessage`, which adds methods that allow you manage the message.

```go
package main

import (
    "context"
	"net/url"

    webhooks "github.com/typical-developers/discord-webhooks-go"
)

func main() {
    ctx := context.Background()

    u := "https://discord.com/api/webhooks/0/0"
    client := webhooks.NewWebhookClientFromURL(u)

    query := url.Values{}
    query.Set("wait", "true")

    message, _, err := client.Execute(ctx,
        webhooks.MessagePayload{
            Content: "Hello World!",
        },
        &query,
    )
    if err != nil {
        panic(err)
    }

    message, _, err := message.Edit(ctx,
        webhooks.EditMessagePayload{
            Content: "Hello World, but edited!!",
        },
        nil,
    )
}
```
