package controllers

type CommentController struct {
	BaseController
}

func (c *CommentController) Lists() {

}

func (c *CommentController) Create() {

	c.JSONResult(0, "ok")
}

func (c *CommentController) Index() {
	c.Prepare()
	c.TplName = "comment/index.tpl"
}
