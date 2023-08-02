//go:build windows
// +build windows

package autorun

import (
	"golang.org/x/sys/windows/registry"
)

// addToAutoStart добавляет команду в автозапуск в Windows
func (a *AutoStart) addToAutoStart() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.SetStringValue(a.AppName, a.Executable)
	if err != nil {
		return err
	}

	return nil
}

// removeFromAutoStart удаляет команду из автозапуска в Windows
func (a *AutoStart) removeFromAutoStart() error {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.ALL_ACCESS)
	if err != nil {
		return err
	}
	defer key.Close()

	err = key.DeleteValue(a.AppName)
	if err != nil {
		return err
	}

	return nil
}

func (a *AutoStart) isAutoEnabled() (bool, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, `SOFTWARE\Microsoft\Windows\CurrentVersion\Run`, registry.READ)
	if err != nil {
		return false, err
	}
	defer key.Close()

	value, _, err := key.GetStringValue(a.AppName)
	if err != nil {
		if strings.Contains(err.Error(), "The system cannot find the file specified.") {
			return false, nil
		}
		return false, err
	}

	return value == a.Executable, nil
}
