package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	var port string
	var sharesDir string
	flag.StringVar(&port, "port", "8080", "设置要绑定的TCP端口,默认端口8080")
	flag.StringVar(&sharesDir, "dir", ".", "设置要共享的目录,默认当前目录")
	flag.Parse()

	// 设置请求处理函数
	http.HandleFunc("/download/", func(w http.ResponseWriter, r *http.Request) {
		// 从URL中获取文件路径
		filePath := r.URL.Path[len("/download/"):]

		// 创建完整路径
		fullPath := filepath.Join(sharesDir, filePath)

		// 检查文件是否存在且不是目录
		if stat, err := os.Stat(fullPath); err == nil && !stat.IsDir() {
			// 设置响应头，使浏览器下载文件
			w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(fullPath)))
			http.ServeFile(w, r, fullPath)
		} else {
			// 如果文件不存在或是一个目录，则返回404错误
			http.NotFound(w, r)
		}
	})

	// 确定计算机的局域网IP地址
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	var ip string
	for _, address := range addrs {
		// 检查IP地址确保不是回环地址
		if inet, ok := address.(*net.IPNet); ok && !inet.IP.IsLoopback() {
			if inet.IP.To4() != nil {
				ip = inet.IP.String()
				break
			}
		}
	}

	// 启动HTTP服务器
	addr := fmt.Sprintf("%s:%s", ip, port)
	serverUrl := fmt.Sprintf("http://%s/download/", addr)
	fmt.Printf("文件共享服务器已启动。\n")
	fmt.Printf("共享目录：%s\n", sharesDir)
	fmt.Printf("服务器监听地址：http://%s\n", addr)
	fmt.Printf("要下载文件，请访问：%s<文件名>\n", serverUrl)

	log.Fatal(http.ListenAndServe(addr, nil))
}
