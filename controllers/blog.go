package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lifei6671/mindoc/models"
	"strconv"
)

type BlogController struct {
	BaseController
}

func SubStr(str string, charnum int) string {
	return string(str[:charnum])
}

func (c *BlogController) Book() {
	c.Prepare()
	c.TplName = "blog/book.tpl"
	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(beego.URLFor("AccountController.Login"), 302)
	}

	book_id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	pageIndex, _ := c.GetInt("page", 1)
	pageSize := 10

	doc, total, err := models.NewDocument().FindToPagerByBookId(pageIndex, pageSize, book_id)

	if err != nil {
		c.Data["DocList"] = make([]*models.Document, 0)
	} else {
		c.Data["DocList"] = doc
	}

	member_id, err := models.NewBook().Find(book_id)
	c.Data["BookType"] = member_id.BookName
	books, err := models.NewBook().FindUser(member_id.MemberId)

	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["Lists"] = books

	member, err := models.NewMember().Find(member_id.MemberId)
	c.Data["Member"] = member

	pp := pageIndex - 1
	if pp < 1 {
		pp = 1
	}

	pn := pageIndex + 1
	if pn > (total/pageSize + 1) {
		pn = total/pageSize + 1
	}

	c.Data["PagePrev"] = pp
	c.Data["PageNext"] = pn
}

func (c *BlogController) User() {
	c.Prepare()
	c.TplName = "blog/user.tpl"
	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(beego.URLFor("AccountController.Login"), 302)
	}

	member_id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	pageIndex, _ := c.GetInt("page", 1)
	pageSize := 10

	books, err := models.NewBook().FindUser(member_id)

	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["Lists"] = books

	doc, total, err := models.NewDocument().FindToPagerByUserId(pageIndex, pageSize, member_id)

	if err != nil {
		c.Data["DocList"] = make([]*models.Document, 0)
	} else {
		c.Data["DocList"] = doc
	}

	member, err := models.NewMember().Find(member_id)
	c.Data["Member"] = member

	pp := pageIndex - 1
	if pp < 1 {
		pp = 1
	}

	pn := pageIndex + 1
	if pn > (total/pageSize + 1) {
		pn = total/pageSize + 1
	}

	c.Data["PagePrev"] = pp
	c.Data["PageNext"] = pn
}

func (c *BlogController) Doc() {
	c.Prepare()
	c.TplName = "blog/doc.tpl"
	doc_id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(beego.URLFor("AccountController.Login"), 302)
	}

	doc, err := models.NewDocument().Find(doc_id)
	if err != nil {
		c.Data["Doc"] = make([]*models.Document, 0)
	} else {
		c.Data["Doc"] = doc
	}

	books, _, err := models.NewBook().FindToPager(1, 15, doc.MemberId)
	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}
	c.Data["Lists"] = books

	member, err := models.NewMember().Find(doc.MemberId)
	c.Data["Member"] = member

	doc, err = models.NewDocument().FindPrev(doc_id)
	if err != nil {
		c.Data["DocPrev"] = make([]*models.Document, 0)
	} else {
		c.Data["DocPrev"] = doc
	}

	doc, err = models.NewDocument().FindNext(doc_id)
	if err != nil {
		c.Data["DocNext"] = make([]*models.Document, 0)
	} else {
		c.Data["DocNext"] = doc
	}
}

func (c *BlogController) Index() {
	c.Prepare()
	c.TplName = "blog/index.tpl"
	//如果没有开启匿名访问，则跳转到登录页面
	if !c.EnableAnonymous && c.Member == nil {
		c.Redirect(beego.URLFor("AccountController.Login"), 302)
	}

	member_id := 0
	pageIndex, _ := c.GetInt("page", 1)
	pageSize := 10

	if c.Member != nil {
		member_id = c.Member.MemberId
	}
	books, _, err := models.NewBook().FindForHomeToPager(1, 15, member_id)

	if err != nil {
		beego.Error(err)
		c.Abort("500")
	}

	c.Data["Lists"] = books

	doc, total, err := models.NewDocument().FindToPager(pageIndex, pageSize)

	if err != nil {
		c.Data["DocList"] = make([]*models.Document, 0)
	} else {
		c.Data["DocList"] = doc
	}

	pp := pageIndex - 1
	if pp < 1 {
		pp = 1
	}

	pn := pageIndex + 1
	if pn > (total/pageSize + 1) {
		pn = total/pageSize + 1
	}
	c.Data["PagePrev"] = pp
	c.Data["PageNext"] = pn
}
