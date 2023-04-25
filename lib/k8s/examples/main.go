// 集群内客户端
// https://github.com/kubernetes/client-go/tree/master/examples/in-cluster-client-configuration
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
		panic(err)
	}
	ctx := context.Background()
	for {
		pods, err := k8s.PodsList(ctx, "default")
		if err != nil {
			logrus.Errorln("pods list error", err)
			time.Sleep(3 * time.Second)
			continue
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		time.Sleep(10 * time.Second)
	}
}
