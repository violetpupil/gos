// 封装 temporal 客户端
// 必须先调用 Init 初始化客户端，才能调用其它函数
package temporal

import (
	"context"

	"github.com/sirupsen/logrus"
	"go.temporal.io/api/enums/v1"
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
		logrus.Error("dial error ", err)
		return err
	}
	temporal = c
	return nil
}

// ExecuteWorkflow 异步执行工作流
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

// ExecuteWorkflowSync 同步执行工作流，停止任务id相同的旧工作流
func ExecuteWorkflowSync(
	ctx context.Context,
	id, queue string,
	workflow any,
	args ...any,
) (client.WorkflowRun, error) {
	options := client.StartWorkflowOptions{
		ID:                    id,
		TaskQueue:             queue,
		WorkflowIDReusePolicy: enums.WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING,
	}
	run, err := temporal.ExecuteWorkflow(ctx, options, workflow, args...)
	return run, err
}

// CancelWorkflow 取消工作流
func CancelWorkflow(ctx context.Context, workflowID string) error {
	err := temporal.CancelWorkflow(ctx, workflowID, "")
	return err
}
