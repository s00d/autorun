package autorun

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type AutoStart struct {
	AppName    string // Имя приложения
	Executable string // Путь к исполняемому файлу приложения
}

func NewAutoStart() *AutoStart {
	appName := filepath.Base(os.Args[0])
	executable, _ := os.Executable()

	return &AutoStart{
		AppName:    appName,
		Executable: executable,
	}
}

// Add добавляет команду в автозапуск
func (a *AutoStart) Enable() error {
	switch runtime.GOOS {
	case "linux":
	case "darwin":
	case "windows":
		return a.addToAutoStart()
	default:
		return errors.New("unsupported operating system")
	}
	return errors.New("unsupported operating system")
}

// Remove удаляет команду из автозапуска
func (a *AutoStart) Disable() error {
	switch runtime.GOOS {
	case "linux":
	case "darwin":
	case "windows":
		return a.removeFromAutoStart()
	default:
		return errors.New("unsupported operating system")
	}
	return errors.New("unsupported operating system")
}

func (a *AutoStart) IsEnabled() (bool, error) {
	switch runtime.GOOS {
	case "linux":
	case "darwin":
	case "windows":
		return a.isAutoEnabled()
	default:
		return false, errors.New("unsupported operating system")
	}
	return false, errors.New("unsupported operating system")
}
