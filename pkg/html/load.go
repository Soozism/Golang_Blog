package html

import "github.com/gin-gonic/gin"

func LoadHTML(r *gin.Engine) {
	 
	r.LoadHTMLGlob("internal/**/**/**/*tmpl")
}