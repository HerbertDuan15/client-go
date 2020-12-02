package main

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func listSysPods(){
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/dhjwork/Desktop/k8s/config.dev")
	if err != nil {
		panic(err)
	}
	config.APIPath = "api"
	config.GroupVersion = &corev1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	result := &corev1.PodList{}

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

func main() {
	listSysPods()
}
