package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"student/model"
)

const (
	dsn string = "root:huanyuan@tcp(127.0.0.1:3306)/go_student_db?charset=utf8mb4&parseTime=True&loc=Local"
)

var (
	db *gorm.DB
)

// 要求用户从命令行输入用户名，密码，确认密码
// 如果输入不为空且两次密码相同，则打印注册成功并结束程序，否则根据情况提示用户输入不能为空或两次密码不一致，并且要求用户重新输入
func main() {
	for {
		showMenu()
		var userSelected uint8
		fmt.Println("请输入您的选择:")
		fmt.Scanln(&userSelected)
		switch userSelected {
		case 1:
			showStudentAllDialog()
		case 2:
			addStudentDialog()
		case 3:
			showStudentDialog()
		case 4:
			deleteStudentDialog()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("无此选项!")
		}
	}
}
func deleteStudentDialog() {
	var studentId int
	fmt.Scanln(&studentId)
	rows := deleteStudentById(studentId)
	if rows > 0 {
		fmt.Println("删除完毕!")
	}

}

// 查看某个学生
func showStudentDialog() {
	var studentName string
	fmt.Println("请输入学生的名字: ")
	fmt.Scanln(&studentName)
	if studentName == "" {
		fmt.Println("学生的名字请填写!")
	}
	student := getStudentByName(studentName)
	fmt.Printf("学生姓名: %s 学生年龄: %d 学生性别: %d 添加时间: %s 修改时间: %s\n", student.Name, student.Age, student.Gender, student.CreatedAt.Format("2006/01/02 15:04:05"), student.UpdatedAt.Format("2006/01/02 15:04:05"))
}

// 添加学生信息
func addStudentDialog() {
	fmt.Println("---------------添加学生信息-----------------")
	for {
		student := new(model.Student)
		var studentName string
		var studentAge, studentGender uint8
		fmt.Println("请输入学生姓名: ")
		fmt.Scanln(&studentName)
		fmt.Println("请输入学生年龄: ")
		fmt.Scanln(&studentAge)
		fmt.Println("请输入学生性别: ")
		fmt.Scanln(&studentGender)

		// 校验

		if studentName == "" || studentAge == 0 || studentGender == 0 {
			fmt.Println("添加学生信息失败!请填写完整")
			continue
		}

		student.Name = studentName
		student.Age = studentAge
		student.Gender = studentGender
		rows := addStudent(student)
		if rows > 0 {
			fmt.Printf("添加[%s]完成!\n", student.Name)
			break
		}
	}
}

func showStudentAllDialog() {
	students := getStudentAll()
	for i := 0; i < len(students); i++ {
		fmt.Printf("学生姓名: %s 学生年龄: %d 学生性别: %d 添加时间: %s 修改时间: %s\n", students[i].Name, students[i].Age, students[i].Gender, students[i].CreatedAt.Format("2006/01/02 15:04:05"), students[i].UpdatedAt.Format("2006/01/02 15:04:05"))
	}

}

func showMenu() {
	fmt.Println("-------------------------------------")
	fmt.Println("1. 显示所有学生列表")
	fmt.Println("2. 添加学生")
	fmt.Println("3. 查看某学生")
	fmt.Println("4. 删除某学生")
	fmt.Println("5. 退出")
	fmt.Println("-------------------------------------")
}

// addStudent
// 添加学生
// return: 影响行数
func addStudent(student *model.Student) int64 {
	result := db.Create(&student)
	return result.RowsAffected
}

// deleteStudentById
// 删除学生
func deleteStudentById(studentId int) int64 {
	return db.Delete(&model.Student{}, studentId).RowsAffected
}

// getUserById
// 根据id获取用户
// return: 查询得到用户结构体
func getStudentById(studentId uint) *model.Student {
	student := model.Student{
		Model: gorm.Model{
			ID: studentId,
		},
	}
	db.Find(&student)
	return &student
}

// getUserByName
// 根据名字获取学生
// bug: name不是唯一
// return: 查询得到学生
func getStudentByName(studentName string) *model.Student {
	student := model.Student{
		Name: studentName,
	}
	db.Where("name = ?", studentName).Find(&student)
	return &student
}

func getStudentAll() []model.Student {
	var students []model.Student
	db.Find(&students)
	return students
}

func init() {
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
