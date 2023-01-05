// 封装 temporal 客户端
// 必须先调用 Init() 初始化客户端，才能调用其它函数
// 可以设置环境变量 TEMPORAL_ADDR，指定 temporal 地址
// 不设置的话，默认为 localhost:7233
package temporal

import (
	"context"
	"os"

	"go.temporal.io/sdk/client"
)

var temporal client.Client

// Init 初始化客户端
func Init() error {
	addr := os.Getenv("TEMPORAL_ADDR")
	options := client.Options{
		HostPort: addr,
	}
	c, err := client.Dial(options)
	if err != nil {
		return err
	}
	temporal = c
	return nil
}

// ExecuteWorkflow 执行工作流
func ExecuteWorkflow(
	ctx context.Context,
	id, queue string,
	workflow any,
	args ...any,
) (client.WorkflowRun, error) {
	options := client.StartWorkflowOptions{
		ID:        id,
		TaskQueue: queue,
	}
	run, err := temporal.ExecuteWorkflow(ctx, options, workflow, args...)
	return run, err
}

// CancelWorkflow 取消工作流
func CancelWorkflow(ctx context.Context, workflowID string) error {
	err := temporal.CancelWorkflow(ctx, workflowID, "")
	return err
}
