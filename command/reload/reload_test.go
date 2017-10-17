package reload

import (
	"strings"
	"testing"

	"github.com/hashicorp/consul/agent"
	"github.com/mitchellh/cli"
)

func TestReloadCommand_noTabs(t *testing.T) {
	if strings.ContainsRune(New(cli.NewMockUi()).Help(), '\t') {
		t.Fatal("usage has tabs")
	}
}

func TestReloadCommand(t *testing.T) {
	a := agent.NewTestAgent(t.Name(), ``)
	defer a.Shutdown()

	// Setup a dummy response to errCh to simulate a successful reload
	go func() {
		errCh := <-a.ReloadCh()
		errCh <- nil
	}()

	ui := cli.NewMockUi()
	c := New(ui)
	args := []string{"-http-addr=" + a.HTTPAddr()}

	code := c.Run(args)
	if code != 0 {
		t.Fatalf("bad: %d. %#v", code, ui.ErrorWriter.String())
	}

	if !strings.Contains(ui.OutputWriter.String(), "reload triggered") {
		t.Fatalf("bad: %#v", ui.OutputWriter.String())
	}
}
