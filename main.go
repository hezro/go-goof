package main
import (
	"fmt"
	"github.com/pstember/go-goof/hello"
    "go-goof/handlers"
    "github.com/gogs/gogs/pkg/tool"
    "github.com/gin-gonic/gin"
)
func main() {
		fmt.Println(tool.MD5(hello.Hello()))
        r := gin.Default()
        r.GET("/ping", handlers.Ping)
        r.Run() // listen and serve on 0.0.0.0:8080
	//test
}
