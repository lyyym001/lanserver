# lanserver
    本地Lanserver

## 编译
    go build -o ./exe/LanServer.exe ./server.go


## 接口
### 1.开启关闭 游戏状态
#### 老师端请求
    参数
    openCode = 2
    subCode = 20036
    data = {
	    Status int //状态 0-关闭 1-开启
	    Type   int //1-静音 2-黑屏 3-护眼
    }
#### 老师端回执
    参数
    openCode = 2
    subCode = 20036
    data = {
	    Code int //0-重复设置 -1-参数错误 1-设置成功
    }
#### 学生端回执
    参数
    openCode = 2
    subCode = 20036
    data = {
	    Status int //状态 0-关闭 1-开启
	    Type   int //1-静音 2-黑屏 3-护眼
    }

### 2.登录状态刷新
#### 老师端学生端回执
    参数
    openCode = 1
    subCode = 10002
    data = {
        CtrlFlag      string
        Code          string
        Mute          int //静音 默认不静音为0 (新增)
        BlackScreen   int //黑屏 默认不黑屏为0 (新增)
        EyeProtection int //护眼 默认不护眼为0 (新增)
    }
    