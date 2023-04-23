// k8s客户端，必须先调用NewForConfig初始化
// https://github.com/kubernetes/client-go/blob/master/INSTALL.md
// https://github.com/kubernetes/client-go/tree/master/examples
package k8s

import (
	"path/filepath"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var Client *kubernetes.Clientset

// NewForConfig 初始化客户端，使用用户家目录下的.kube/config配置
func NewForConfig() error {
	home := homedir.HomeDir()
	kubeconfig := filepath.Join(home, ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		logrus.Errorln("build config error", err)
		return err
	}
	Client, err = kubernetes.NewForConfig(config)
	return err
}
