package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	/*type BOOK struct {
		id, ISBN float64
		name     string
	}
	book1 = BOOK{01, 1231, "貓的經濟學"}
	book2 = BOOK{02, 12341, "雞腿說：說什ㄇ"}
	book3 = BOOK{03, 12384721, "百富單一威士忌"}
	*/
	//GroupTest v1 for method GET

	v1 := router.Group("/v1")
	{
		v1.GET("/index", getWeb)
		v1.GET("/books", books)
		v1.GET("/books/:id", book)
		router.LoadHTMLFiles("index.html")
	}

	//GroupTest v2 for Method PUT/POST
	v2 := router.Group("/v2")
	{
		v2.PUT("/book", book)
		v2.POST("/book", bookAdd)
		router.Run(":8080")
	}

}

func getWeb(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}

func books(c *gin.Context) {
	c.JSON(http.StatusOK, "{id:01, ISBN:1231, name:貓的經濟學/id:02, ISBN:12cs31, name:單一麥芽純麥/id:03, ISBN:18571, name:麥卡倫}")

}

func book(c *gin.Context) {
	id := c.Param("id")

	switch c.Request.Method {
	case http.MethodGet:
		switch id {
		case "01":
			c.String(http.StatusOK, "id:%v, ISBN:1231, name:貓的經濟學\n", id)
		case "02":
			c.String(http.StatusOK, "id:02, ISBN:12cs31, name:單一麥芽純麥")
		case "03":
			c.String(http.StatusOK, "id:03, ISBN:18571, name:麥卡倫")
		}

	case http.MethodPut:
		type BOOK struct {
			ID   float64
			ISBN float64
			Name string
		}
		book1 := BOOK{01, 1231, "貓的經濟學"}
		book2 := BOOK{02, 12341, "雞腿說：說什ㄇ"}
		book3 := BOOK{03, 12384721, "百富單一威士忌"}

		newId := c.PostForm("ID")
		newName := c.PostForm("Name")
		newISBN := c.PostForm("ISBN")

		fmt.Printf("newid = %s\n", newId)
		switch newId {
		case "01":
			book1.ID, _ = strconv.ParseFloat(newId, 64)
			book1.ISBN, _ = strconv.ParseFloat(newISBN, 64)
			book1.Name = newName
		case "02":
			book2.ID, _ = strconv.ParseFloat(newId, 64)
			book2.ISBN, _ = strconv.ParseFloat(newISBN, 64)
			book2.Name = newName
		case "03":
			book3.ID, _ = strconv.ParseFloat(newId, 64)
			book3.ISBN, _ = strconv.ParseFloat(newISBN, 64)
			book3.Name = newName
		}
		var data []BOOK
		fmt.Println(newId)
		c.JSON(http.StatusOK, data)

	}
}

func bookAdd(c *gin.Context) {
	type BOOK struct {
		Id   *int    `json:"id,string"`
		Name *string `json:"name"`
		ISBN *string `json:"isbn"`
	}
	var b BOOK

	err := c.ShouldBindWith(&b, binding.JSON)
	if err != nil {
		log.Print(err)
	}

	id := b.Id
	name := b.Name
	ISBN := b.ISBN
	//id := c.PostForm("id")
	//name := c.PostForm("name")
	//ISBN := c.PostForm("ISBN")
	//res := fmt.Sprintf("thanks for your donation!\nBook_info--\nbook_id:%v\nbook_name:%s\nISBN:%s", *id, *name, *ISBN)
	if id != nil && name != nil && ISBN != nil {
		c.JSON(http.StatusOK, b)
	} else {
		c.String(200, "欄位不可留白")
	}

}
