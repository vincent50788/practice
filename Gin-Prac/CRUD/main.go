package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//Book struct 存放格式---------------------------------------------------------------
type Book struct {
	ID   uint   `db:"id" json:"id" binding:"id"`
	Name string `db:"name" json:"name" binding:"name"`
	ISBN string `db:"ISBN" json:"ISBN" binding:"ISBN"`
}

//AddBook 存放格式---------------------------------------------------------------
type AddBook struct {
	Name string `db:"name" json:"name" binding:"name"`
	ISBN string `db:"ISBN" json:"ISBN" binding:"ISBN"`
}

// type Data struct{
// 	Data []*Book
// }

//DB連線--------------------------------------------------------------------------
var db, err = sql.Open("mysql", "root:vincent50788@tcp(127.0.0.1:3306)/CRUD_BOOK")

//enterPoint router--------------------------------------------------------------
func main() {
	router := gin.Default()
	router.GET("/books", books)
	router.GET("/books/:id", getBook)
	router.POST("/books", addBooks)
	router.PUT("/books/:id", putBook)
	router.DELETE("/books/:id", deleteBook)
	router.PATCH("/books/:id", patchBook)
	router.Run(":8080")
}

//PATCH 更新書籍 by ID------------------------------------------------------------
func patchBook(c *gin.Context) {
	ParamID := c.Param("id")
	name := c.PostForm("name")
	ISBN := c.PostForm("ISBN")
	//檢查JSON是否完整
	if name == "" || ISBN == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Please check all the JSON fields"),
		})
		return
	}
	//更新書籍
	update, err := db.Prepare("UPDATE Book SET name= ?, ISBN= ? where id= ?;")
	_, err = update.Exec(name, ISBN, ParamID)
	if err != nil {
		c.String(200, "資料發生衝突")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Successfully updated to %s", name),
		})
	}
}

//DELETE刪除書籍 by ID-----------------------------------------------------------
func deleteBook(c *gin.Context) {
	ParamID := c.Param("id")
	delete, err := db.Prepare("DELETE FROM Book WHERE id= ?;")
	if err != nil {
		fmt.Print(err.Error())
	}
	_, err = delete.Exec(ParamID)
	if err != nil {
		fmt.Print(err.Error())
	}
	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted BookID = %s", ParamID),
	})

}

//PUT修改書籍 by ID--------------------------------------------------------------
func putBook(c *gin.Context) {
	ParamID := c.Param("id")
	var books []AddBook
	update, err := db.Prepare("UPDATE Book SET name= ?, ISBN= ? where id= ?;")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(body, &books)
	if err != nil {
		fmt.Println(err.Error())
	}
	for i, e := range books {
		//檢查JSON是否完整
		if e.Name == "" || e.ISBN == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Please check all the JSON fields"),
			})
			return
		}
		_, err = update.Exec(books[i].Name, books[i].ISBN, ParamID)
		if err != nil {
			c.String(200, "資料發生衝突")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Successfully updated to %s", books[0].Name),
			})

		}

	}

}

//POST新增書籍------------------------------------------------------------------
func addBooks(c *gin.Context) {
	//var book Book
	var books []AddBook
	//sql新增書籍準備
	INSERT, err := db.Prepare("INSERT INTO Book(name, ISBN) VALUES(?, ?);")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(body, &books)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(books)
	for i, e := range books {
		//檢查JSON是否完整
		if e.Name == "" || e.ISBN == "" {
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("Please check all the JSON fields"),
			})
			return
		}
		_, err = INSERT.Exec(books[i].Name, books[i].ISBN)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "successfully created"})

		}

	}
}

//GET取得所有書籍----------------------------------------------------------------
func books(c *gin.Context) {
	var (
		book  Book
		books []Book
	)
	rows, err := db.Query("SELECT * FROM Book;")

	if err != nil {
		fmt.Print(err.Error())
	}
	//使用Next()將資料裝進 []books
	for rows.Next() {
		rows.Scan(&book.ID, &book.Name, &book.ISBN)
		books = append(books, book)
	}

	defer rows.Close()
	c.JSON(200, books)
}

//ＧＥＴ 取得單一書籍 by ID---------------------------------------------
func getBook(c *gin.Context) {
	var (
		book Book
	)
	id := c.Param("id")
	row := db.QueryRow("SELECT * FROM Book WHERE id=?;", id)
	err = row.Scan(&book.ID, &book.Name, &book.ISBN)
	if err != nil {
		c.JSON(http.StatusOK, nil)
	} else {
		c.JSON(http.StatusOK, book)
	}

}
