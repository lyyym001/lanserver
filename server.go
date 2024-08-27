package main

import (
	"fmt"
	"lanserver/lframework/utils"
	"lanserver/lframework/ziface"
	"lanserver/lframework/znet"
	"lanserver/pb"

	"lanserver/api"
	"lanserver/core"
	"os"
	"strings"
)

//业务Api 这里定义跟客户都安通信的业务关联
//1	-	登录账号相关
//2 - 	房间业务

// 当客户端建立连接的时候的hook函数
func OnConnecionAdd(conn ziface.IConnection) {
	//创建一个玩家
	player := core.NewPlayer(conn)

	//同步当前玩家的初始化坐标信息给客户端，走MsgID:200消息
	//player.BroadCastStartPosition()

	//将当前新上线玩家添加到worldManager中
	core.WorldMgrObj.AddPlayer(player)

	//将该连接绑定属性PID
	conn.SetProperty("pID", player.PID)

	//同步周边玩家上线信息，与现实周边玩家信息
	//player.SyncSurrounding()

	//同步当前的PlayerID给客户端， 走MsgID:1 消息 这里需要客户端回执 登录信息
	player.SyncPID()

	fmt.Println("=====> Player pIDID = ", player.PID, " arrived ====")
}

// 当客户端断开连接的时候的hook函数
func OnConnectionLost(conn ziface.IConnection) {

	//获取当前连接的PID属性
	pID, _ := conn.GetProperty("pID")
	//fmt.Println("pID = " , pID)
	//根据pID获取对应的玩家对象
	player := core.WorldMgrObj.GetPlayerByPID(pID.(int32))
	if player != nil {
		fmt.Println(player)
		fmt.Println("Player Lost  player= ", player.CID, " Room = ", player.TID)
		//触发玩家下线业务
		if pID != nil {
			fmt.Println("Player Lost  pID= ", pID)
			player.LostConnection()
		}
	}

	//fmt.Println("====> Player ", pID, " left =====")
	//fmt.Println("123")
}

func main() {
	//创建服务器句柄
	s := znet.NewServer()

	//注册客户端连接建立和丢失函数
	s.SetOnConnStart(OnConnecionAdd)
	s.SetOnConnStop(OnConnectionLost)

	//启动本地老师端
	print("Start Up TClient\n")
	_, err := os.StartProcess("StartTClient.bat", nil, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if err != nil {
		print("TClient Started Error\n")
	} else {
		print("TClient Started Succ\n")
	}
	//test()
	//注册路由

	//登录路由
	s.AddRouter(1, &api.AccountApi{})
	//聊天路由
	s.AddRouter(2, &api.RoomApi{})
	//课程路由
	s.AddRouter(3, &api.CourseApi{})
	//启动服务
	s.Serve()

}

func test() {

	var setupDate, courseMode string

	var sData pb.AllStudyInfoData

	courseName := "交通安全"

	fmt.Println("cRoom.AllCourses = ", "getData.CourseID = ", "1010002", "getData.CourseID = ", "1010002", " p.TID = ", "xiaoshoubu27")

	var studentRecord pb.SingleStudyInfoData
	var allStudentRecord []pb.SingleStudyInfoData
	db := utils.GlobalObject.SqliteInst.GetDB()
	rows, err := db.Query("select SetupDate,CourseMode from tb_recore where CourseId=? and Tid=? group by SetupDate order by CreateDate desc limit 10", "1010002", "xiaoshoubu27")
	if err != nil {
		fmt.Println("Sqlite Handle_onStudentRecord Query DB Err")
	} else {
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&setupDate, &courseMode); err == nil {
				setupDate = strings.Replace(setupDate, "T", " ", -1)
				setupDate = strings.Replace(setupDate, "Z", "", -1)
				fmt.Println("setupDate = ", setupDate)
				var studyTotalTime, studyScore int
				var pname, snum, sclass, courseId, studyTime, studyMode, studyAbility string
				var srData pb.StudentRecordData
				var allSrData []pb.StudentRecordData
				rows, err := db.Query("select ts.snum,ts.pname,ts.class,tr.CourseId,tr.SetupDate,tr.CourseMode,tr.StuTimeLong,tr.Score,tr.Ability from tb_recore tr inner join tb_snum ts on tr.SetupDate=? and tr.Tid=? and ts.SNum=tr.SNum", setupDate, "xiaoshoubu27")

				defer rows.Close()
				if err == nil {
					for rows.Next() {
						if err := rows.Scan(&snum, &pname, &sclass, &courseId, &studyTime, &studyMode, &studyTotalTime, &studyScore, &studyAbility); err == nil {
							srData.StudentSnum = snum
							srData.StudentName = pname
							srData.StudentClass = sclass
							srData.CourseID = courseId
							srData.StudyTime = studyTime
							srData.CourseName = courseName
							srData.StudyMode = studyMode
							srData.StudyTotalTime = studyTotalTime
							srData.StudyScore = studyScore
							srData.StudyAbility = studyAbility
							allSrData = append(allSrData, srData)
						}
					}
				}
				//if allSrData != nil {
				studentRecord.StudyTime = setupDate
				studentRecord.StudyCourseMode = courseMode
				studentRecord.StudentRecordData = allSrData
				allStudentRecord = append(allStudentRecord, studentRecord)
				//}
			} else {
				fmt.Println("Handle_onStudyRecordByCourse,", err)
			}
		}

		//
		sData.StudyInfoData = allStudentRecord
		fmt.Println("sData = ", sData)
	}

}
