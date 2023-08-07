// 打印k8s集群资源信息
// cd tools/docker/resource-info
// set GOOS=linux&set CGO_ENABLED=0&go build -o app .
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/violetpupil/components/lib/k8s"
)

func main() {
	err := k8s.NewInCluster()
	if err != nil {
		logrus.Fatalln("new in cluster error", err)
	}

	ctx := context.Background()
	for {
		pods, err := k8s.PodsList(ctx, k8s.NamespaceDefault)
		if err != nil {
			logrus.Errorln("pods list error", err)
		} else {
			for _, pod := range pods.Items {
				fmt.Println(pod.ObjectMeta.Name)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
