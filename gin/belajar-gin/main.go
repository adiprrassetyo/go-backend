package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type User struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Address Address `json:"address"`
}

type Login struct {
	Username `json:"username" binding:"required"`
	Password `json:"password" binding:"required"`
}

// middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// set dummy example dengan value '12345'
		c.Set("example", "12345")

		// before request
		c.Next()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

// multiple middleware
func KeyAccess() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.Query("key")

		if key == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}

		c.Next()
	}
}

func getting(c *gin.Context) {
	//5// membuat data user dummy
	user := User{
		ID:   1,
		Name: "afis",
		Age:  20,
		Address: Address{
			Street:  "Jl. Jendral Sudirman",
			City:    "Jakarta",
			Country: "Indonesia",
		},
	}
	c.JSON(http.StatusOK, gin.H{"message": "someGet", "user": user})
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "somePost"})
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "somePut"})
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "someDelete"})
}

func main() {
	// membuat router gin
	r := gin.Default()

	// 7// menggunakan middleware
	r.Use(Logger())

	// example //
	r.GET("/example", func(c *gin.Context) {
		// get value dari middleware
		example := c.MustGet("example").(string)

		// menampilkan value dari middleware
		c.JSON(http.StatusOK, gin.H{"example": example})
	})

	// 8// menggunakan middleware untuk group route
	v1 := r.Group("/v1", Logger(), KeyAccess())
	{
		v1.GET("/example", func(c *gin.Context) {
			example := c.MustGet("example").(string)
			c.JSON(http.StatusOK, gin.H{"example": example})
		})
		v1.GET("/hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "hello"})
		})
	}

	// 5// html rendering
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"Title": "Home Page"})
	})

	//1// menambah router dengan method GET, POST, PUT, DELETE
	r.GET("/", getting)
	r.POST("/", posting)
	r.PUT("/", putting)
	r.DELETE("/", deleting)
	//2// menambahkan parameter pada path => : name
	r.GET("/user/:name", func(c *gin.Context) {
		// mengambil parameter name
		name := c.Param("name")

		// anggap saja ada logic untuk mengambil data user berdasarkan name

		// menampilkan data user
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + name})

		// http://localhost:8080/user/adit // => {"message": "Hello adit"}
	})

	//3// menambahkan parameter pada query => ? name
	r.GET("/user", func(c *gin.Context) {
		// mengambil value dari query name dan age
		name := c.Query("name")
		age := c.Query("age")

		// anggap saja ada logic untuk mengambil data user berdasarkan name dan age

		// menampilkan data user
		c.JSON(http.StatusOK, gin.H{"message": "Hello " + name + " umur " + age})

		// http://localhost:8080/user?name=afis&age=20 // => {"message": "Hello afis umur 20"}

	})

	//4// grouping route prefix '/transaction'
	transaction := r.Group("/transaction")
	{
		// route '/transaction' dengan method GET (mengambil data)
		transaction.GET("/", getting)
		// route '/transaction' dengan method POST (menambah data)
		transaction.POST("/", posting)
		// route '/transaction' dengan method PUT (mengubah data)
		transaction.PUT("/", putting)
		// route '/transaction' dengan method DELETE (menghapus data)
		transaction.DELETE("/", deleting)
	}

	// 6// login
	r.POST("/login", func(c *gin.Context) {
		var login Login

		err := c.ShouldBindJSON(&login) // otomatis melakukan validasi
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		// anggap ada logic untuk mengecek data user login

		// dummy response berhasil login
		c.JSON(http.StatusOK, gin.H{
			"message": "success login",
		})
	})

	// menjalankan server di port :8080
	r.Run(":8080")
}
