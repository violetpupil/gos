package k8s

import (
	"context"
	"fmt"
	"testing"

	"k8s.io/client-go/util/homedir"
)

func TestHomeDir(t *testing.T) {
	fmt.Println(homedir.HomeDir())
}

func TestPodsList(t *testing.T) {
	err := NewForConfig()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	pods, err := PodsList(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pods)
}

func TestPodsGet(t *testing.T) {
	err := NewForConfig()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	pod, err := PodsGet(ctx, "kube-system", "etcd-docker-desktop")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", pod)
}
