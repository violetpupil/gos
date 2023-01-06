// 封装 temporal 客户端
// 必须先调用 Init 初始化客户端，才能调用其它函数
package temporal

import (
	"context"

	"go.temporal.io/sdk/client"
)

var temporal client.Client

// Init 初始化客户端
// addr 传空的话，默认连接 localhost:7233
func Init(addr string) error {
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
