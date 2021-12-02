package login

import (
	"github.com/dengpju/higo-gin/higo"
)

type LoginController struct {
}

func NewLoginController() *LoginController {
	return &LoginController{}
}

func (this *LoginController) New() higo.IClass {
	return NewLoginController()
}

func (this *LoginController) Route(hg *higo.Higo) {
	/**
	    //TODO::example
		//route example
		hg.Get("/relative1", this.Example1, hg.Flag("this.Example1"), hg.Desc("Example1"))
		hg.Get("/relative2", this.Example2, hg.Flag("this.Example2"), hg.Desc("Example2"))
		hg.Get("/relative3", this.Example3, hg.Flag("this.Example3"), hg.Desc("Example3"))
		hg.Get("/relative4", this.Example4, hg.Flag("this.Example4"), hg.Desc("Example4"))
		hg.Get("/relative5", this.Example5, hg.Flag("this.Example5"), hg.Desc("Example5"))
	    //route group example
	    hg.AddGroup("/group_prefix", func() {
	    	hg.Get("/relative6", this.Example6, hg.Flag("this.Example6"), hg.Desc("Example6"))
	    })
	*/
}

/**
func (this *{{.Name}}) Example1() {
    //TODO::example code
	ctx := request.Context()
	//get parameter
    name := ctx.Query("name")
    //responser
    responser.SuccessJson("success", 10000, name)
}
*/

/**
//responser string
func (this *LoginController) Example2() string {
    //TODO::example code
    ctx := request.Context()
    //get parameter
    name := ctx.Query("name")
    return name
}
*/

/**
//responser interface{}
func (this *LoginController) Example3() interface{} {
    //TODO::example code
    ctx := request.Context()
    //get parameter
    name := ctx.Query("name")
    return name
}
*/

/**
//example Model
type LoginControllerExampleModel struct {
	Id   int
	Name string
}

func (this *LoginControllerExampleModel) New() higo.IClass {
	return &LoginControllerExampleModel{}
}

func (this *LoginControllerExampleModel) Mutate(attrs ...higo.Property) higo.Model {
	higo.Propertys(attrs).Apply(this)
	return this
}
*/

/**
//responser Model
func (this *LoginController) Example4(ctx *gin.Context) higo.Model {
    //TODO::example code
	model := &LoginControllerExampleModel{Id: 1, Name: "foo"}
	return model
}
*/

/**
//responser Models
func (this *LoginController) Example5(ctx *gin.Context) higo.Models {
    //TODO::example code
	var models []*LoginControllerExampleModel
	models = append(models, &LoginControllerExampleModel{Id: 1, Name: "foo"}, &LoginControllerExampleModel{Id: 2, Name: "bar"})
	return higo.MakeModels(models)
}
*/

/**
//responser Json
func (this *LoginController) Example6(ctx *gin.Context) higo.Json {
    //TODO::example code
	var models []*LoginControllerExampleModel
	models = append(models, &LoginControllerExampleModel{Id: 1, Name: "foo"}, &LoginControllerExampleModel{Id: 2, Name: "bar"})
	return models
}
*/
