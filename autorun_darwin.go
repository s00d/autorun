//go:build darwin
// +build darwin

package autorun

import (
	"os"
)

// addToAutoRun добавляет команду в автозапуск в macOS
func (a *AutoRun) addToAutoRun() error {
	plistFilePath := a.getPlistFilePath()
	plistFileContent := a.getPlistFileContent()

	plistFile, err := os.Create(plistFilePath)
	if err != nil {
		return err
	}
	defer func(plistFile *os.File) {
		err := plistFile.Close()
		if err != nil {

		}
	}(plistFile)

	_, err = plistFile.WriteString(plistFileContent)
	if err != nil {
		return err
	}

	return nil
}

// removeFromAutoRun удаляет команду из автозапуска в macOS
func (a *AutoRun) removeFromAutoRun() error {
	plistFilePath := a.getPlistFilePath()

	err := os.Remove(plistFilePath)
	if err != nil {
		return err
	}

	return nil
}

// getPlistFilePath возвращает путь к файлу .plist в macOS
func (a *AutoRun) getPlistFilePath() string {
	return "~/Library/LaunchAgents/" + a.AppName + ".plist"
}

// getPlistFileContent возвращает содержимое файла .plist в macOS
func (a *AutoRun) getPlistFileContent() string {
	return `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>` + a.AppName + `</string>
	<key>ProgramArguments</key>
	<array>
		<string>` + a.Executable + `</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
</dict>
</plist>`
}

func (a *AutoRun) isAutoEnabled() (bool, error) {
	plistFilePath := a.getPlistFilePath()

	_, err := os.Stat(plistFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
