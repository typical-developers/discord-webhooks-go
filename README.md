# Discord Webhooks
A module to handle sending messages to Discord webhooks in Golang.

# Example Usage
## Sending a Message
```go
package main

import "github.com/typical-developers/discord-webhooks-go"

func main() {
    clientId := "YOUR_CLIENT_ID"
    secret := "YOUR_CLIENT_SECRET"

    webhook := webhooks.NewWebhook(webhooks.NewWebhookArgs{
		ClientID: &clientId,
		Secret:   &secret,
	})

    content := "Hello, world!"

    payload := webhooks.WebhookPayload{
        Content: &content,
    }

    err := webhook.Send(&payload)
    if err != nil {
        println(err.Error())
    }
}
```

## Attaching a File
```go
package main

import "github.com/typical-developers/discord-webhooks-go"

func main() {
    clientId := "YOUR_CLIENT_ID"
    secret := "YOUR_CLIENT_SECRET"

    webhook := webhooks.NewWebhook(webhooks.NewWebhookArgs{
		ClientID: &clientId,
		Secret:   &secret,
	})

    content := "Hello, world!"
    file, err := os.Open("./example.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    payload := webhooks.WebhookPayload{
        Content: &content,
        Files: []*webhooks.WebhookFile{
            {
                Name:   "example.txt",
                Reader: file,
            },
        },
    }

    err := webhook.Send(&payload)
    if err != nil {
        println(err.Error())
    }
}
```

# In-Progress Features
- [x] Building the client with a URL or the client id and secret.
- [x] Setting a default user profile for the webhook client (username and avatar).
- [x] Sending basic messages.
  - [x] Including embeds.
  - [x] Including attachments.
- [ ] Message builders.
- [ ] Editing existing webhook messages.
- [ ] Deleting the webhook.