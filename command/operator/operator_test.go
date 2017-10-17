package operator

import (
	"strings"
	"testing"
)

func TestOperatorCommand_noTabs(t *testing.T) {
	if strings.ContainsRune(New().Help(), '\t') {
		t.Fatal("help has tabs")
	}
}
