package pomodoro_test

import (
	"testing"

	"github.com/pavelanni/pragprog-go-cli-book/interactiveTools/pomo/pomodoro"
	"github.com/pavelanni/pragprog-go-cli-book/interactiveTools/pomo/pomodoro/repository"
)

func getRepo(t *testing.T) (pomodoro.Repository, func()) {
	t.Helper()

	return repository.NewInMemoryRepo(), func() {}
}
