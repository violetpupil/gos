// k8s客户端，必须先调用NewForConfig初始化
// https://github.com/kubernetes/client-go/blob/master/INSTALL.md
// https://github.com/kubernetes/client-go/tree/master/examples
package k8s

import (
	"context"
	"errors"
	"path/filepath"

	"github.com/sirupsen/logrus"
	api "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var Client *kubernetes.Clientset

// NewForFile 初始化客户端，使用用户家目录下的.kube/config配置
func NewForFile() error {
	home := homedir.HomeDir()
	if home == "" {
		return errors.New("get home directory fail")
	}

	kubeconfig := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logrus.Errorln("build config error", err)
		return err
	}
	Client, err = kubernetes.NewForConfig(config)
	return err
}

func NewInCluster() error {
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Errorln("in cluster config error", err)
		return err
	}
	Client, err = kubernetes.NewForConfig(config)
	return err
}

// PodsList 获取所有pod信息
func PodsList(ctx context.Context) (*api.PodList, error) {
	opts := meta.ListOptions{}
	// 所有命名空间
	pods, err := Client.CoreV1().Pods("").List(ctx, opts)
	return pods, err
}

// PodsGet 获取单个pod信息
func PodsGet(ctx context.Context, namespace, name string) (*api.Pod, error) {
	opts := meta.GetOptions{}
	pod, err := Client.CoreV1().Pods(namespace).Get(ctx, name, opts)
	return pod, err
}
