// https://github.com/spf13/cobra/blob/main/site/content/user_guide.md
// Commands represent actions, Args are things and Flags are modifiers for those actions.
package cobra

import "github.com/spf13/cobra"

type (
	// 命令
	// Use one-line usage message
	// Short short description shown in the 'help' output
	// Long long message shown in the 'help <this-command>' output
	// RunE 执行函数，返回错误
	//
	// Execute() 执行对应命令
	// ExecuteC() 执行对应命令，并返回命令
	// Flags() 选项
	Command = cobra.Command
)
