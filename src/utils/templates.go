package utils

import (
	"fmt"
	"net/http"

	"github.com/flosch/pongo2/v6"
	"github.com/gin-gonic/gin"
)

func RenderTemplateResponse(ctx *gin.Context,
	templateName string,
	context map[string]interface{},
) {
	template := pongo2.Must(
		pongo2.FromFile(fmt.Sprintf("./src/templates/%s", templateName)),
	)
	err := template.ExecuteWriter(context, ctx.Writer)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}

}
