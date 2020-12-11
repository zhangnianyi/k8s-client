package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s.io/api/apps/v1"
	"k8s-client/lib"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func main(){
	//var kubeconfig *string
	//if home := homedir.HomeDir(); home != "" {
	//	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//} else {
	//	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//}
	//flag.Parse()
	//
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//if err != nil {
	//	panic(err)
	//}
	//clientset, err := kubernetes.NewForConfig(config)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(clientset.ServerVersion())
	////ctx :=context.Background()
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
	nginxDep:=&v1.Deployment{}
	//把yaml读取出来 b
	b,_ :=ioutil.ReadFile("yamls/nginx.yaml")
	//读出来了之后抓换成json
	nginxjsonyaml,_ :=yaml.ToJSON(b)
	//然后在解析到jsontoyaml的结构体里面
	_=json.Unmarshal(nginxjsonyaml,nginxDep) //已经把值传递给了nginxDep  然后下面直接创建就可以了

	_,err =lib.KubeClient.AppsV1().Deployments("vizion").Create(nginxDep)
	if err !=nil{
		fmt.Println("KubeClient.AppsV1().Deployments faild",err)
	}
}
