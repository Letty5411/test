package main

import (
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	"context"
	"encoding/json"
)

func PrettyDump(hint string, obj interface{}) string {
	ret, _ := json.MarshalIndent(obj, "", "  ")
	fmt.Printf("%s: %s\n", hint, ret)
	return fmt.Sprintf("%s:%s", hint, string(ret))
}
func main() {
	// 使用 kubeconfig 获取 Kubernetes 配置。
	kubeconfig := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// 创建 Kubernetes 客户端。
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pod, err := clientset.CoreV1().Pods("e2e-virtual-kubelet").Get(context.TODO(), "e2e-testinstanceupdateimagecheckip-qn4lp", metav1.GetOptions{})
	if err != nil {
		fmt.Printf("failed to get pod:%s\n", err.Error())
	}
	PrettyDump("beforupdate", pod)
	pod.Spec.Containers[0].Image = "busybox:abc"
	pod, err = clientset.CoreV1().Pods("e2e-virtual-kubelet").Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		fmt.Printf("failed to update pod:%s\n", err.Error())
		os.Exit(1)
	}
	PrettyDump("after update", pod)

	// 列出所有 Pod（在所有命名空间中）。
	//pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	//// 打印出每个 Pod 的名称和命名空间。
	//for _, pod := range pods.Items {
	//	fmt.Printf("Name: %s, Namespace: %s\n", pod.Name, pod.Namespace)
	//}
}
