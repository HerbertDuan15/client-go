package main

import (
	"context"
	"fmt"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

/////////////////////list pods in kube-system namespace
func listSysPods() {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/duanhongjian/config.dev")
	if err != nil {
		panic(err)
	}
	//fmt.Println(config)
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}

	for {
		err = restClient.Get().
			Namespace("kube-system").
			Resource("pods").
			VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
			Do(context.TODO()).
			Into(result)
		if err != nil {
			panic(err)
		}

		for _, item := range result.Items {
			fmt.Printf("Namespace:%v \t status:%+v \t Name:%v\n", item.Namespace, item.Status.Phase, item.Name)
		}
	}

}

////////////////////////end

//////////////////list pods number
func getPodsNum() {
	// var kubeconfig *string
	// if home := homeDir(); home != "" {
	// 	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	// } else {
	// 	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	// }
	// flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", "/home/duanhongjian/config.dev")
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	for {
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
		//time.Sleep(10 * time.Second)

	}
}
func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

/////////////////end

func findPodNotRunning() {
	//  调用 BuildConfigFromFlags 函数获取 rest.Config 对象
	config, err := clientcmd.BuildConfigFromFlags("", "/home/duanhongjian/config.dev")
	if err != nil {
		panic(err)
	}
	// 配置API路径和请求的资源组/资源版本信息
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	// 通过rest.RESTClientFor()生成RESTClient对象
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}
	// 通过RESTClient构建请求参数，查询default空间下所有pod资源
	err = restClient.Get().
		Namespace("").
		Resource("pods").
		VersionedParams(&metav1.ListOptions{Limit: 500}, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result.Items[0].Status.Phase) // 得到 pod 的状态
	for _, pod := range result.Items {
		if podStatus := pod.Status.Phase; podStatus == "Running" { // 如果是 running 则 continue
			continue
		} else { //打印当前的状态以及原因
			fmt.Printf("NameSpace: %s, Pod Name: %s, Pod Status: %s\n", pod.Namespace, pod.Name, podStatus)
			// fmt.Printf("Reason for not running1: %v\n", pod.Status.Conditions)
			// fmt.Printf("Reason for not running2: %v\n", pod.Status.ContainerStatuses)
			// fmt.Printf("Reason for not running3: %v\n", pod.Status.EphemeralContainerStatuses)
			// fmt.Printf("Reason for not running4: %v\n", pod.Status.InitContainerStatuses)
			fmt.Printf("Reason for not running5: %v\n", pod.Status.Message)
			// fmt.Printf("Reason for not running6: %v\n", pod.Status.NominatedNodeName)
			// fmt.Printf("Reason for not running7: %v\n", pod.Status.Phase)
			// fmt.Printf("Reason for not running8: %v\n", pod.Status.PodIP)
			// fmt.Printf("Reason for not running9: %v\n", pod.Status.PodIPs)
			fmt.Printf("Reason for not running10: %v\n", pod.Status.Conditions)
			//fmt.Printf("Reason for not running11: %v\n", pod.Status.QOSClass)
		}
	}
}

func main() {
	//listSysPods()
	//getPodsNum()
	findPodNotRunning()
}
