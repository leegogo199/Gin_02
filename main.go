package main
import(
"github.com/gin-gonic/gin"
"net/http"
)
func main(){
    router:=gin.Default()
    router.GET("/",func(c *gin.Context){
           c.String(http.StatusOK,"it works")
    })
    router.Run(":8000")

}
