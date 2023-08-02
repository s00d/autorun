//go:build darwin
// +build darwin

package autorun

import (
	"github.com/kardianos/service"
)

// addToAutoRun добавляет команду в автозапуск в macOS
func (a *AutoRun) addToAutoRun() error {
	config := &service.Config{
		Name:        a.AppName,
		DisplayName: a.AppName,
		Executable:  a.Executable,
	}

	// Создание новой службы
	s, err := service.New(nil, config)
	if err != nil {
		return err
	}

	// Установка службы
	err = s.Install()
	if err != nil {
		return err
	}

	return nil
}

// removeFromAutoRun удаляет команду из автозапуска в macOS
func (a *AutoRun) removeFromAutoRun() error {
	config := &service.Config{
		Name:        a.AppName,
		DisplayName: a.AppName,
		Executable:  a.Executable,
	}

	// Создание новой службы
	s, err := service.New(nil, config)
	if err != nil {
		return err
	}

	// Удаление службы
	err = s.Uninstall()
	if err != nil {
		return err
	}

	return nil
}

// isAutoEnabled проверяет, включена ли автозагрузка в macOS
func (a *AutoRun) isAutoEnabled() (bool, error) {
	config := &service.Config{
		Name:        a.AppName,
		DisplayName: a.AppName,
		Executable:  a.Executable,
	}

	// Создание новой службы
	s, err := service.New(nil, config)
	if err != nil {
		return false, err
	}

	// Проверка статуса службы
	status, err := s.Status()
	if err != nil {
		if err == service.ErrNotInstalled {
			return false, nil
		}
		return false, err
	}

	return status == service.StatusRunning, nil
}
