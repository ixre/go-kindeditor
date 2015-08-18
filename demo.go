/**
 * Copyright 2015 @ z3q.net.
 * name : demo.go
 * author : jarryliu
 * date : 2015-08-18 23:19
 * description :
 * history :
 */
package go_kindeditor

import (
	"github.com/jrsix/gof/web"
	"strings"
	"encoding/json"
	"github.com/jrsix/gof/web/mvc"
)


var _ mvc.Filter = new(editorController)

type editorController struct {
}

func (this *editorController) Requesting(*web.Context) bool{
	//todo: check permission
	return true
}

func (this *editorController) RequestEnd(*web.Context){
}

func (this *editorController) File_manager(ctx *web.Context) {
	d, err := fileManager(ctx.Request,"./uploads/","http://img.abc.com/uploads/")
	ctx.Response.Header().Add("Content-Type", "application/json")
	if err != nil {
		ctx.Response.Write([]byte("{error:'" + strings.Replace(err.Error(), "'", "\\'", -1) + "'}"))
	} else {
		ctx.Response.Write(d)
	}
}

func (this *editorController) File_upload_post(ctx *web.Context) {
	fileUrl, err := fileUpload(ctx.Request,"./uploads/","http://img.abc.com/uploads/")
	var hash map[string]interface{} = make(map[string]interface{})
	if err == nil {
		hash["error"] = 0;
		hash["url"] = fileUrl;
	}else {
		hash["error"] = 1
		hash["message"] = err.Error()
	}
	ctx.Response.Header().Add("Content-Type", "application/json")
	d, _ := json.Marshal(hash)
	ctx.Response.Write(d)
}
