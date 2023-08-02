//go:build darwin
// +build darwin

package autostart

import (
	"os"
)

// addToAutoStart добавляет команду в автозапуск в macOS
func (a *AutoStart) addToAutoStart() error {
	plistFilePath := a.getPlistFilePath()
	plistFileContent := a.getPlistFileContent()

	plistFile, err := os.Create(plistFilePath)
	if err != nil {
		return err
	}
	defer plistFile.Close()

	_, err = plistFile.WriteString(plistFileContent)
	if err != nil {
		return err
	}

	return nil
}

// removeFromAutoStart удаляет команду из автозапуска в macOS
func (a *AutoStart) removeFromAutoStart() error {
	plistFilePath := a.getPlistFilePath()

	err := os.Remove(plistFilePath)
	if err != nil {
		return err
	}

	return nil
}

// getPlistFilePath возвращает путь к файлу .plist в macOS
func (a *AutoStart) getPlistFilePath() string {
	return "~/Library/LaunchAgents/" + a.AppName + ".plist"
}

// getPlistFileContent возвращает содержимое файла .plist в macOS
func (a *AutoStart) getPlistFileContent() string {
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

func (a *AutoStart) isAutoEnabled() (bool, error) {
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
