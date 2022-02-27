package handler

import (
	"boos/pkg/database"
	"boos/pkg/model"
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 增
func Post(c *gin.Context) {
	//连接数据库
	db := database.NewDB()
	//定义一个数据
	resume := model.Resume{}
	//获取浏览器数据
	err := c.ShouldBind(&resume)
	if err != nil {
		panic(err)
	}
	//判断手机号
	if !(regexp.MustCompile("^1[0-9]{10}$").Match([]byte(resume.Phone))) {
		c.String(200, "输入手机号错误")
		return
	}
	resume.CreateAt = time.Now()
	fmt.Println(resume)
	//添加数据
	if err := db.Create(&resume).Error; err != nil {
		fmt.Println("插入失败", err)
		return
	}
	c.String(200, fmt.Sprintf("success insert resume for %s", resume.Name))

}

func Get(c *gin.Context) {
	//连接数据库
	db := database.NewDB()
	//判断查询类型
	phone, ok := c.GetQuery("phone")
	if ok {
		//创建一个容器
		var resume model.Resume
		//查询一个数据
		err := db.Where("phone=?", phone).Take(&resume).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("查询不到数据")
			c.String(200, "找不到")
			return
		} else if err != nil {
			fmt.Println("查询失败", err)
			c.String(200, "未知错误")
			return
		}
		c.JSON(200, resume)
	}
	var resumes []model.Resume
	err := db.Find(&resumes).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("查询不到数据")
		c.String(200, "找不到")
		return
	} else if err != nil {
		fmt.Println("查询失败", err)
		c.String(200, "未知错误")
		return
	}
	fmt.Println(resumes)
	c.JSON(200, resumes)
}
func Put(c *gin.Context) {
	//连接数据库
	db := database.NewDB()
	resume := model.Resume{}
	//要修改的数据
	phone := c.Query("phone")
	//先查询一条数据
	err := db.Where("phone=?", phone).Take(&resume).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("修改的数据不存在")
		c.String(200, "修改的数据不存在")
		return
	} else if err != nil {
		fmt.Println("查询失败", err)
		c.String(200, "未知错误")
		return
	}
	//获取修改数据的值
	c.ShouldBind(&resume)
	//更新数据
	db.Where("phone", phone).Save(&resume)

	c.String(200, fmt.Sprintf("success update resume for %s", resume.Phone))
}

func DELETE(c *gin.Context) {
	//删除数据
	phone := c.Query("phone")
	//建立数据库连接
	db := database.NewDB()
	//删除指定数据
	number := db.Where("phone=?", phone).Delete(&model.Resume{}).RowsAffected
	if number == 0 {
		c.String(200, "数据已删除")
	} else {
		c.String(200, fmt.Sprintf("success delete resume for %s", phone))
	}

}
