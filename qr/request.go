package qr

import (
	"github.com/MuhammadSuryono/module-golang-server/http/validate"
	"github.com/gin-gonic/gin"
)

func QrValidateRequest(ctx *gin.Context) {
	var request requestUploadFile
	_ = ctx.Bind(&request)
	valid := validate.NewValidate()
	valid.ValidationStruct(request).JsonResponse(ctx)
}
