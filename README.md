## 介绍
这是一个局域网内的文件共享服务，用于将局域网A电脑的某个目录下的文件共享给局域网其它设备。

## 使用方式

### 1. window电脑
如果你使用window电脑，下载此项目下的`fileshare.exe`文件后,在CMD执行
`fileshare.exe -port 8080 -dir C:\path\to\your\shares`后即可启动
- port http服务端口，默认8080
- dir 共享文件夹路径，默认当前目录
> port和dir都是可省略

启动后通过`http://局域网ip地址:port/dowload/<文件名>`执行下载。

一个偷懒的办法是，下载这个文件后，把它和要下载的文件放在一起，这样它就和要下载的文件在同一个目录，直接双击点开运行，ok啦！

### 2. mac电脑
下载此项目下`fileshare`文件，然后终端执行
`./fileshare -port 8080 -dir /path/to/your/shares` 即可启动服务

### 注意点
有可能服务器启动后仍然不能使用，可能是防火墙导致的，可以开放指定端口解决。




