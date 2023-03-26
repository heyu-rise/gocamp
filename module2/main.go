package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/healthz", version)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func version(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	// 接收客户端 request，并将 request 中带的 header 写入 response header
	for s, strings := range header {
		for _, s2 := range strings {
			w.Header().Add(s, s2)
		}
	}
	// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	versionOs := os.Getenv("VERSION")
	w.Header().Add("VERSION", versionOs)
	// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	statusCode := 200
	statusCodeStr := "200"
	ip := r.RemoteAddr
	log.Printf("ip: %s, statusCode: %s", ip, statusCodeStr)
	myMap := make(map[string]string, 10)
	myMap["ip"] = ip
	myMap["statusCode"] = statusCodeStr
	// 当访问 localhost/healthz 时，应返回 200
	w.WriteHeader(statusCode)
	jsons, _ := json.Marshal(myMap)
	io.WriteString(w, string(jsons))
}
