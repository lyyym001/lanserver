# lanserver
    本地Lanserver

## 编译
    go env -w GOOS=windows
    go env -w GOARCH=amd64
    go env -w CGO_ENABLED=1
    go build -o ./exe/LanServer.exe ./server.go
## 编译linux
    go env -w GOOS=linux
    go env -w GOARCH=amd64
    go env -w CGO_ENABLED=0


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


### 3.添加灵创作品
#### 老师端请求
    参数
    openCode = 3
    subCode = 30018
    data = {
        CourseName string //课程名称
        CourseID   string //灵创课程ID
    }
#### 老师端回执
    参数
    openCode = 3
    subCode = 30018
    data = {
        Code int   //0-作品已经存在 1-成功
        Id   int64 //作品dbid(跟其他第三方一样前端记录dbid用来操作)
    }

### 3.更新灵创作品
#### 老师端请求
    参数
    openCode = 3
    subCode = 30019
    data = {
        CourseName string //课程名称
        CourseID   string //灵创课程ID
        Id         int64  //作品dbid
    }

#### 老师端回执
    参数
    openCode = 3
    subCode = 30019
    data = {
        Code int   //0-更新失败 1-成功
        Id   int64 //作品dbid(跟其他第三方一样前端记录dbid用来操作)
    }

### 4.检查课程是否存在  post
#### 本地测试url:http://192.168.0.22:9001/post 
#### 老师端请求
    参数
    openCode = 5
    subCode = 1
    data = {
        courseid string //base64编码后的课程数据
    }

#### 老师端回执
    {
        openCode = 5
        subCode = 1
        data = ""   //0-参数异常 -1-课程不存在 1-课程存在
    }