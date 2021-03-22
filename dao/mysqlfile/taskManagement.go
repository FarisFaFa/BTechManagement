package mysqlfile

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Task struct {
	TaskId       string `json:"taskId" param:"taskId"`
	Location     string `json:"location" param:"location"`
	Deadline     string `json:"deadline" param:"deadline"`
	InchargeId   string `json:"inchargeId" param:"inchargeId"`
	ContactName  string `json:"contactName" param:"contactName"`
	ContactPhone string `json:"contactPhone" param:"contactPhone"`
	ContactEmail string `json:"contactEmail" param:"contactEmail"`
	Completed    int    `json:"completed" param:"completed"`
}

type Task_table struct {
	TaskId       string `xorm:"varchar(20) pk notnull 'task_id'" json:"taskId" from:"TaskId" binding:"required"`
	Location     string `xorm:"varchar(40) notnull 'location'" json:"location" form:"Location" binding:"required" `
	Deadline     string `xorm:"date notnull 'deadline'" json:"deadline" form:"Deadline" binding:"required"`
	InchargeId   string `xorm:"varchar(20) 'incharge_id'" json:"inchargeId" form:"InchargeId" binding:"required"`
	ContactName  string `xorm:"varchar(20) 'contact_name'" json:"contactName" form:"ContactName"`
	ContactPhone string `xorm:"varchar(10) 'contact_phone'" json:"contactPhone" form:"ContactPhone"`
	ContactEmail string `xorm:"varchar(20) 'contact_email'" json:"contactEmail" form:"ContactEmail"`
	Completed    int    `xorm:"tinyint default 0 notnull 'complete'" json:"completed" form:"Completed"`
}

// sql语句实现获取到所有任务
func GetAllTasks(c *gin.Context) {
	// 写查询语句
	// 将task_table中所有的信息查询出来
	sqlStr := `select task_id,location,deadline,incharge_id,contact_name,contact_phone,contact_email,complete from task_table`
	fmt.Println(db)
	rowsObj, e := db.Query(sqlStr)
	if e != nil {
		fmt.Printf("error is %v", e)
		//return
	}
	// 设定一个计数器，记录数量
	var length = 0
	//遍历查询到的数据

	var allTasks = make([]Task, 0, 400)
	for rowsObj.Next() {
		//结构体存储一行数据
		var oneTask Task
		e := rowsObj.Scan(&oneTask.TaskId, &oneTask.Location, &oneTask.Deadline, &oneTask.InchargeId, &oneTask.ContactName, &oneTask.ContactPhone, &oneTask.ContactEmail, &oneTask.Completed)
		if e != nil {
			fmt.Printf("error is %v", e)
			//return
		} else {
			// 将这行数据添加到切片里面
			allTasks = append(allTasks, oneTask)
			length++
		}
	}
	// 定义一个结构体用于返回多个数据
	type Message struct {
		// 主要数据都放在Data中
		Data []Task `json:"Data"`
	}
	// 生命结构体变量
	var Msg Message
	// 给结构体变量赋值
	Msg.Data = allTasks
	// 解析Json数据，会变为二进制数据
	//result, _ := json.Marshal(Msg)
	// 用Json返回结构体
	c.JSON(200, Msg)
	rowsObj.Close()
}

// xorm实现获取所有数据
func XormGetAllTasks(c *gin.Context) {
	//allTasks := make(map[int64]Task1)
	var allTasks []Task_table
	err := x.Find(&allTasks)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err})
		return
	}
	c.JSON(http.StatusOK, allTasks)
}

// 插入数据
func InsertTask(c *gin.Context) {
	// 声明接受的变量
	var NewTask Task_table
	// 绑定表单传回的数据Bind()默认解析并绑定form格式
	// 根据请求头中content-type自动推断
	if err := c.Bind(&NewTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println(NewTask)
	if _, err := x.Insert(NewTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "添加成功")
}
