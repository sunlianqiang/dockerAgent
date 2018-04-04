package controllers

import (
	"dockerAgent/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ImageController struct {
	beego.Controller
}

// // URLMapping ...
// func (c *ImageController) URLMapping() {
// 	c.Mapping("Post", c.Post)
// 	c.Mapping("Get", c.Get)

// }

func (c *ImageController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

// func (c *AppController) Post() {
// 	var v models.App
// 	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
// 		if _, err := models.AddApp(&v); err == nil {
// 			c.Ctx.Output.SetStatus(201)
// 			c.Data["json"] = v
// 		} else {
// 			c.Data["json"] = err.Error()
// 		}
// 	} else {
// 		c.Data["json"] = err.Error()
// 	}
// 	c.ServeJSON()
// }

func (c *ImageController) Post() {
	logs.Debug("image pull")

	// var image string = "docker.io/library/alpine"
	// docker.io/library/alpine
	logs.Debug("post para:%+v\n", c.Input())
	image := c.Ctx.Input.Param(":name")

	var v = models.Image{
		Name: image,
	}
	logs.Debug("req body:%+v\n", c.Ctx.Input.RequestBody)

	if err := models.ImagePull(image); nil == err {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
		logs.Debug("image pull err:", err.Error())
	}

	c.ServeJSON()
}

// func (c *ImageController) Post() {
// 	logs.Debug("image pull")

// 	// var image string = "docker.io/library/alpine"
// 	// docker.io/library/alpine
// 	logs.Debug("post para:%+v\n", c.Input())
// 	image := c.Input().Get("image")

// 	var v = models.Image{
// 		Name: image,
// 	}
// 	logs.Debug("req body:%+v\n", c.Ctx.Input.RequestBody)

// 	if err := models.ImagePull(image); nil == err {
// 		c.Ctx.Output.SetStatus(201)
// 		c.Data["json"] = v
// 	} else {
// 		c.Data["json"] = err.Error()
// 		logs.Debug("image pull err:", err.Error())
// 	}

// 	c.ServeJSON()
// }
