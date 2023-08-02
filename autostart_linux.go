//go:build linux
// +build linux

package autorun

import (
	"os"
)

// addToAutoStart добавляет команду в автозапуск в Linux
func (a *AutoStart) addToAutoStart() error {
	desktopFilePath := a.getDesktopFilePath()
	desktopFileContent := a.getDesktopFileContent()

	desktopFile, err := os.Create(desktopFilePath)
	if err != nil {
		return err
	}
	defer desktopFile.Close()

	_, err = desktopFile.WriteString(desktopFileContent)
	if err != nil {
		return err
	}

	return nil
}

// removeFromAutoStart удаляет команду из автозапуска в Linux
func (a *AutoStart) removeFromAutoStart() error {
	desktopFilePath := a.getDesktopFilePath()

	err := os.Remove(desktopFilePath)
	if err != nil {
		return err
	}

	return nil
}

// getDesktopFilePath возвращает путь к файлу .desktop в Linux
func (a *AutoStart) getDesktopFilePath() string {
	return "~/.config/autostart/" + a.AppName + ".desktop"
}

// getDesktopFileContent возвращает содержимое файла .desktop в Linux
func (a *AutoStart) getDesktopFileContent() string {
	return `[Desktop Entry]
Type=Application
Name=` + a.AppName + `
Exec=` + a.Executable + `
Hidden=false
X-GNOME-Autostart-enabled=true`
}

func (a *AutoStart) isAutoEnabled() (bool, error) {
	desktopFilePath := a.getDesktopFilePath()

	_, err := os.Stat(desktopFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
