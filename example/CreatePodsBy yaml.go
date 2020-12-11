package example

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"k8s-client/lib"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/yaml"
)

func CreatepodsByyaml(){
	nginxDep:=&v1.Deployment{}
	//把yaml读取出来 b
	b,_ :=ioutil.ReadFile("yamls/nginx.yaml")
	//读出来了之后抓换成json
	nginxjsonyaml,_ :=yaml.ToJSON(b)
	//然后在解析到jsontoyaml的结构体里面
	err :=json.Unmarshal(nginxjsonyaml,nginxDep)
	if err !=nil{
		fmt.Println("json.Unmarshal faild")
	}
	//_:=json.Unmarshal(nginxjsonyaml,nginxDep) //已经把值传递给了nginxDep  然后下面直接创建就可以了

	_,err =lib.KubeClient.AppsV1().Deployments("vizion").Create(nginxDep)
	if err !=nil{
		fmt.Println("KubeClient.AppsV1().Deployments faild",err)
	}
}