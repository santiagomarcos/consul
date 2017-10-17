package keygen

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/mitchellh/cli"
)

func TestKeygenCommand_noTabs(t *testing.T) {
	if strings.ContainsRune(New(nil).Help(), '\t') {
		t.Fatal("help has tabs")
	}
}

func TestKeygenCommand(t *testing.T) {
	ui := cli.NewMockUi()
	cmd := New(ui)
	code := cmd.Run(nil)
	if code != 0 {
		t.Fatalf("bad: %d", code)
	}

	output := ui.OutputWriter.String()
	result, err := base64.StdEncoding.DecodeString(output)
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(result) != 16 {
		t.Fatalf("bad: %#v", result)
	}
}
