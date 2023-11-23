package cobra

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

func TestCommand_ExecuteC(t *testing.T) {
	cmd := cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("run")
			return nil
		},
	}
	cmd.ExecuteC()
}
