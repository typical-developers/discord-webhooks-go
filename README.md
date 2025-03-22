# Discord Webhooks
A module to handle sending messages to Discord webhooks in Golang.

## Example Usage
### Sending a Message
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

    payload := webhooks.WebhookPayload{}
    payload.SetContent("Hello, world!")

    err := webhook.Send(&payload)
    if err != nil {
        println(err.Error())
    }
}
```

### Sending a Message (with embeds)
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

    payload := webhooks.WebhookPayload{}
    payload.SetContent("Hello, world!")

    embed := payload.AddEmbed()
    embed.SetTitle("Hello, world!")
    embed.SetDescription("This is a test embed.")
    embed.SetTimestamp(time.Now())

    customFieldOne := embed.AddField()
    customFieldOne.SetName("Custom Field 1")
    customFieldOne.SetValue("This is a custom field.")
    customFieldOne.SetInline(true)

    customFieldTwo := embed.AddField()
    customFieldTwo.SetName("Custom Field 2")
    customFieldTwo.SetValue("This is another custom field.")
    customFieldTwo.SetInline(true)

    err := webhook.Send(&payload)
    if err != nil {
        println(err.Error())
    }
}
```

### Sending a Message (with attachments)
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

    payload := webhooks.WebhookPayload{}
    payload.SetContent("Hello, world!")

    file, err := os.Open("./example.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()

    payload.AddAttachment("example.txt", file)

    err := webhook.Send(&payload)
    if err != nil {
        println(err.Error())
    }
}
```

## In-Progress Features
- [x] Building the client with a URL or the client id and secret.
- [x] Setting a default user profile for the webhook client (username and avatar).
- [x] Sending basic messages.
  - [x] Including embeds.
  - [x] Including attachments.
- [x] Message builders.
- [ ] Editing existing webhook messages.
- [ ] Deleting the webhook.