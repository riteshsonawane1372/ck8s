package pkg

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ConfigFile(){

	rules:= clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfigrt:= clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules,&clientcmd.ConfigOverrides{})

	config,err:=kubeconfigrt.ClientConfig()
	if err!=nil{
		log.Fatal("An error occured",err)
	}

	clientset:= kubernetes.NewForConfigOrDie(config)

	listPod,err:= clientset.CoreV1().Pods("default").List(context.Background(),v1.ListOptions{})
	if err!=nil{
		log.Fatal("Error While Listing Pods",err)
	}
	
	for _,i:=range (listPod.Items){
		fmt.Println(i.Name)
	}

}