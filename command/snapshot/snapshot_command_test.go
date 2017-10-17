package snapshot

import (
	"strings"
	"testing"
)

func TestSnapshotCommand_noTabs(t *testing.T) {
	if strings.ContainsRune(New().Help(), '\t') {
		t.Fatal("usage has tabs")
	}
}
