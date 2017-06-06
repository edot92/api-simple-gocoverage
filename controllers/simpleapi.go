package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/edot92/api-simple-gocoverage/models"
)

// SimpelapiController ...
type SimpelapiController struct {
	beego.Controller
}

// Getsimpelapi ...
// @Title Get value of  :)
// @Description Get results for param array
// @Param	body		body 	models.Simpelapi.StructParamDeret true		"value array , ex 1,2,3"
// @Success 200 {object} models.Simpelapi.ResSuccses "res json"
// @Failure 400 value format error
// @router /getsimpelapi [post]
func (u *SimpelapiController) Getsimpelapi() {
	var paramBody models.StructParamDeret
	json.Unmarshal(u.Ctx.Input.RequestBody, &paramBody)
	lengthParam := len(paramBody.ParamDeret)
	fmt.Println(lengthParam)
	if lengthParam < 2 || lengthParam > 3 {
		u.Ctx.Output.Context.WriteString("Array length must 2 or 3 , ex: 1,2,3")
		u.Ctx.Output.SetStatus(400)
		return
	}
	resultsModels, err := models.GetSimpelAPI(paramBody.ParamDeret)
	if err != nil {
		u.Ctx.Output.Context.WriteString(err.Error())
		u.Ctx.Output.SetStatus(400)
		return
	}
	var resJSON models.ResSuccses
	resJSON.Message = "Berhasil :)"
	resJSON.Data = resultsModels
	u.Data["json"] = resJSON
	u.Ctx.ResponseWriter.WriteHeader(200)
	u.Ctx.Output.Header("Access-Control-Allow-Origin", "*")

	u.ServeJSON()
}
