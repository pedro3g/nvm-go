package utils_test

import (
	"testing"

	"github.com/pedro3g/nvm-go/utils"
)

func TestGetBaseDir(t *testing.T) {
	_, err := utils.GetBaseDir()

	if err != nil {
		t.Errorf("Expect a base dir, got %v", err)
	}
}
