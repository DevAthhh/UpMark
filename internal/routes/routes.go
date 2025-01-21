package routes

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

type Msg struct {
	Body string `json:body`
	Name string `json:name`
}

func Route(rtr *gin.Engine) {
	crud := rtr.Group("/api/v1.0.0")
	{
		crud.POST("/interaction", func(ctx *gin.Context) {
			b := Msg{}
			ctx.BindJSON(&b)
			create_or_update(b)
			ctx.JSON(200, gin.H{
				"msg": b,
			})
		})
		crud.GET("/interaction", func(ctx *gin.Context) {
			name := ctx.Query("name")
			result := search_file(name)
			ctx.JSON(200, gin.H{
				"msg": result,
			})
		})
	}
}

func create_or_update(b Msg) {
	file, err := os.OpenFile("data/"+b.Name, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(b.Body)
	if err != nil {
		panic(err)
	}
}

func search_file(name string) string {
	file, err := os.Open("data/" + name)
	if err != nil {
		return "File doesn't exists!"
	}
	defer file.Close()

	data := make([]byte, 64)
	res := ""

	for {
		n, err := file.Read(data)
		if err == io.EOF {
			break
		}
		res += string(data[:n])
	}
	return res
}
