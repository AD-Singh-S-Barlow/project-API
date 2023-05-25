
package main
import (
   "net/http"
   "github.com/gin-gonic/gin"
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