package example

import (
	"fmt"
	"k8s-client/lib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetdeploymentList(){
	listOpt :=metav1.ListOptions{}
	////获取vizion这个namespace下面的所有service的列表
	//list,_ :=clientset.CoreV1().Services("vizion").List(listOpt)
	//for _,item :=range list.Items{
	//	fmt.Println(item.Name)
	//}
	//获取deployment的列表
	delist,err :=lib.KubeClient.AppsV1().Deployments("vizion").List(listOpt)
	if err !=nil{
		fmt.Println("get k8s deployment list faild")
	}
	//获取到deployment的详细信息
	for _,v  :=range delist.Items{
		fmt.Println(v.Name,v.Status)
	}
}
