
package main
import (
   "net/http"
   "github.com/gin-gonic/gin"
   "github.com/dgrijalva/jwt-go"
   "github.com/auth0/go-jwt-middleware"
)
type NameList struct {
  NameF string  `json: "fName:"`
  NameL string  `json: "lName:"`
  Age  float32 `json: "Age:"`
}
var names = []NameList{
  {NameF: "Quandale", NameL: "Dingle the third", Age: 59},
  {NameF: "James", NameL: "Howlet", Age: 134},
  {NameF: "Bihan", NameL: "Tundra", Age: 38},
}
func getNames(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, names)
}
func handlerFunc(c *gin.Context) {
  c.String(200, "My record of Kira's kills")
}
func main() {
   router := gin.Default()
  router.GET("/", handlerFunc)
  router.GET("/names", getNames)
  router.PUT("/names", putNames)
  router.POST("/names", postNames)
  router.DELETE("/names", deleteNames)
  router.Run()
}
var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("N3qQbwCVqkc7Yg7ciSmc1borLgMfy1jhXUI_9Vb4c_ZSt9xBYHtKx3nf0Pf9TTXd"), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func checkJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMid := *jwtMiddleware
		if err := jwtMid.CheckJWT(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(401)
		}
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func deleteNames(C *gin.Context) {
   var deleteItem NameList
   if err := C.ShouldBindJSON(&deleteItem); err != nil {
     C.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
     return
  }
   var remove bool
   for i, d := range names {
     if d.NameF == deleteItem.NameF && d.NameL == deleteItem.NameL && d.Age == deleteItem.Age {
       names = append(names[:i], names[i+1:]...)
       remove = true
       break
    }
  }
   if remove {
    C.JSON(http.StatusOK, gin.H{"message": "Resource deleted successfully"})
  } else {
    C.JSON(http.StatusNotFound, gin.H{"message": "Data not found"})
  }
}
func putNames(c *gin.Context) {
   var update NameList
   if err := c.ShouldBindJSON(&update); err != nil {
     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
     return
  }
   names = append(names, update)
  c.IndentedJSON(http.StatusCreated, update)
}
func postNames(c *gin.Context) {
   var newName NameList
   if err := c.BindJSON(&newName); err != nil {
     return
  }
   names = append(names, newName)
  c.IndentedJSON(http.StatusCreated, newName)
}