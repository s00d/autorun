package autorun

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
)

type AutoRun struct {
	AppName    string // Имя приложения
	Executable string // Путь к исполняемому файлу приложения
}

func NewAutoRun() *AutoRun {
	appName := filepath.Base(os.Args[0])
	executable, _ := os.Executable()

	return &AutoRun{
		AppName:    appName,
		Executable: executable,
	}
}

// Add добавляет команду в автозапуск
func (a *AutoRun) Enable() error {
	switch runtime.GOOS {
	case "linux":
		return a.addToAutoRun()
	case "darwin":
		return a.addToAutoRun()
	case "windows":
		return a.addToAutoRun()
	default:
		return errors.New("unsupported operating system")
	}
}

// Remove удаляет команду из автозапуска
func (a *AutoRun) Disable() error {
	switch runtime.GOOS {
	case "linux":
		return a.removeFromAutoRun()
	case "darwin":
		return a.removeFromAutoRun()
	case "windows":
		return a.removeFromAutoRun()
	default:
		return errors.New("unsupported operating system")
	}
}

func (a *AutoRun) IsEnabled() (bool, error) {
	switch runtime.GOOS {
	case "linux":
		return a.isAutoEnabled()
	case "darwin":
		return a.isAutoEnabled()
	case "windows":
		return a.isAutoEnabled()
	default:
		return false, errors.New("unsupported operating system")
	}
}
