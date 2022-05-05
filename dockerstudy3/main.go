package main

import (
	"fmt"
	"os"
	"strconv"

	//"github.com/thinkeridea/go-extend/exnet"
	"io"
	"log"
	"net/http"
	"strings"
)

/*
编写一个 HTTP 服务器，大家视个人不同情况决定完成到哪个环节，但尽量把 1 都做完：

	1，接收客户端 request，并将 request 中带的 header 写入 response header
	2，读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	3，Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	4，当访问 localhost/healthz 时，应返回 200
*/

// Main方法入口
func main() {
	println("环境正常")

	// 功能1
	http.HandleFunc("/requestAndResponse", requestAndResponse)

	// 功能2
	http.HandleFunc("/getVersion", getVersion)

	// 功能3
	http.HandleFunc("/ipAndStatus", ipAndStatus) //注册接口句柄

	// 功能4
	http.HandleFunc("/healthz", healthz) //注册接口句柄

	err := http.ListenAndServe(":81", nil) //监听空句柄，80端口被占用，使用81端口
	if nil != err {
		log.Fatal(err) //显示错误日志
	}
}

// 功能1，接收请求及响应
func requestAndResponse(response http.ResponseWriter, request *http.Request) {
	println("调用requestAndResponse接口")
	headers := request.Header //header是Map类型的数据
	println("传入的hander：")
	for header := range headers { //value是[]string
		//println("header的key：" + header)
		values := headers[header]
		for index, _ := range values {
			values[index] = strings.TrimSpace(values[index])
			//println("index=" + strconv.Itoa(index))
			//println("header的value：" + values[index])

		}
		//valueString := strings.Join(values, "")
		//println("header的value：" + valueString)
		println(header + "=" + strings.Join(values, ","))        //打印request的header的k=v
		response.Header().Set(header, strings.Join(values, ",")) // 遍历写入response的Header
		//println()

	}
	fmt.Fprintln(response, "Header全部数据:", headers)
	io.WriteString(response, "succeed")

}

// 功能2，获取环境变量的version
func getVersion(response http.ResponseWriter, request *http.Request) {
	println("调用getVersion接口")
	envStr := os.Getenv("VERSION")
	//envStr := os.Getenv("HADOOP_HOME")
	//println("系统环境变量：" + envStr) //可以看到 C:\soft\hadoop-3.3.1	Win10需要重启电脑才能生效

	response.Header().Set("VERSION", envStr)
	io.WriteString(response, "succeed")

}

// 功能3，输出IP与返回码
func ipAndStatus(response http.ResponseWriter, request *http.Request) {
	println("调用ipAndStatus接口")

	form := request.RemoteAddr
	println("Client->ip:port=" + form) //虚拟机是桥接模式。使用postman返回的全部是127.0.0.1	用手机打开网站192.168.1.139:81/ipAndStatus可以看到新IP
	ipStr := strings.Split(form, ":")
	println("Client->ip=" + ipStr[0]) //打印ip

	// 获取http响应码
	//response.WriteHeader(301) //手动设置响应码，默认200

	//response.WriteHeader(http.StatusOK)//由于默认是调用这个，∴返回码都是这个200【server.go有源码】
	println("Client->response code=" + strconv.Itoa(http.StatusOK))

	//println("response code->：" + code)

	io.WriteString(response, "succeed")
}

// 功能4，连通性测试接口
func healthz(response http.ResponseWriter, request *http.Request) {
	println("调用healthz接口")
	response.WriteHeader(200) //设置返回码200
	//response.WriteHeader(http.StatusOK)//默认会调用这个方法，默认就是200【server.go有源码】
	io.WriteString(response, "succeed")
}

