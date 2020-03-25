package controllers

import (
	"com.waschild/noChaos-Server/models"
	"fmt"
	"github.com/astaxie/beego"
)

type FormController struct {
	beego.Controller
}

// @Title	CreatedServlet
// @Description 创建服务
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /create [Post]
func (f *FormController) Create() {

	form := models.NC_Form{}
	form.Name = f.GetString("name")
	form.IsStore = true

	col0 := models.NC_Field{}
	col0.Name = "id"
	col0.Type = "整数"
	col0.Size = 32
	col0.Default = "0"
	col0.AutoIncrement = true
	col0.IsKey = true

	col1 := models.NC_Field{}
	col1.Name = "长文本"
	col1.Type = "文本"
	col1.Size = 6667652
	col1.NotNull = true
	col1.Default = "默认文本"

	col11 := models.NC_Field{}
	col11.Name = "长文本2"
	col11.Type = "文本"
	col11.Size = 255
	col11.Default = "默认文本2"

	col2 := models.NC_Field{}
	col2.Name = "整数32"
	col2.Type = "整数"
	col2.Size = 32
	col2.Default = "0"
	col2.AutoIncrement = true

	col3 := models.NC_Field{}
	col3.Name = "整数64"
	col3.Type = "整数"
	col3.Size = 64
	col3.Index = true
	col3.Unique = true

	col4 := models.NC_Field{}
	col4.Name = "小数32"
	col4.Type = "小数"
	col4.Size = 32
	col4.Decimal = 2
	col4.Maximum = 6

	col5 := models.NC_Field{}
	col5.Name = "小数64"
	col5.Type = "小数"
	col5.Size = 64
	col5.Decimal = 2
	col5.Maximum = 5

	col6 := models.NC_Field{}
	col6.Name = "布尔"
	col6.Type = "布尔"

	col7 := models.NC_Field{}
	col7.Name = "时间"
	col7.Type = "时间"

	col8 := models.NC_Field{}
	col8.Name = "日期"
	col8.Type = "日期"

	col9 := models.NC_Field{}
	col9.Name = "日期时间"
	col9.Type = "日期时间"

	col10 := models.NC_Field{}
	col10.Name = "时间戳"
	col10.Type = "时间戳"

	//col1.FormId = form.ID
	//col1.FormId = form.ID
	//col1.FormId = form.ID
	//col1.FormId = form.ID
	//col1.FormId = form.ID
	//col1.FormId = form.ID
	//col1.FormId = form.ID

	form.Fields = []models.NC_Field{col0, col1, col2, col3, col4, col5, col6, col7, col8, col9, col10, col11}

	//form.Fields = append(form.Fields, field)

	models.NCDB.Create(&form)
	for _, field := range form.Fields {
		field.FormId = form.ID
		models.NCDB.Create(&field)
	}

	//models.NCDB.Create(&field)
	//models.NCDB.Model(&form).Related(&[]models.NC_Field{field})
	//servlet.Create()
	f.Data["json"] = "create form success"
	f.ServeJSON()
}

// @Title	CreatedServlet
// @Description 创建服务
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /search [Post]
func (f *FormController) Search() {

	//一对多查询
	var form models.NC_Form
	models.NCDB.First(&form)
	models.NCDB.Debug().Model(&form).Related(&form.Fields, "FormId")

	fmt.Println(form)

	f.Data["json"] = "create model success"
	f.ServeJSON()
}

// @Title	CreatedServlet
// @Description 创建服务
// @Param	name	string	true		"服务名称"
// @Success 200 编译成功
// @Failure 403 body is empty
// @router /build [Post]
func (f *FormController) Build() {

	//一对多查询
	var form models.NC_Form
	id, _ := f.GetInt("id")
	form.ID = uint(id)
	models.NCDB.First(&form)
	models.NCDB.Debug().Model(&form).Related(&form.Fields, "FormId")

	form.GetCode()

	fmt.Println(form)

	f.Data["json"] = "create model success"
	f.ServeJSON()
}