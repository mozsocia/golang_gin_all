Give me full code for golang gin api - use GORM - file name and folder structure - jwt auth with user model - login, register route - models, controllers, routes folder - config/config.go file for database setup - utils for jwt -.env file

Here is an example of a basic Golang Gin API using GORM for database operations, JWT authentication for user management, and the recommended file structure:

```
project
|__ config
|   |__ config.go
|   |__ .env
|__ controllers
|   |__ authController.go
|__ models
|   |__ user.go
|__ routes
|   |__ authRoutes.go
|__ utils
|   |__ jwtUtils.go
|__ main.go

```

config/config.go:

```go
package config

import (
    "fmt"
    "os"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbName := os.Getenv("DB_NAME")
    dbPassword := os.Getenv("DB_PASSWORD")

    dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return db, nil
}

```

utils/jwtUtils.go:

```go
package utils

import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type Claims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateToken(username string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })

    if err != nil {
        return nil, err
    }

    if _, ok := token.Claims.(*Claims); !ok && !token.Valid {
        return nil, err
    }

    return token, nil
}

```

models/user.go:

```go
package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Username string `gorm:"unique" json:"username"`
    Password string `json:"password"`
}

```

controllers/authController.go:

```go
package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"

    "project/models"
    "project/utils"
)

func Register(c *gin.Context) {
    var user models.User
    c.BindJSON(&user)

    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    user.Password = string(hashedPassword)

    db, err := config.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }
    db.Create(&user)

    token, _ := utils.GenerateToken(user.Username)

    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}


func Login(c *gin.Context) {
    var user models.User
    c.BindJSON(&user)

    db, err := config.ConnectDB()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    var existingUser models.User
    db.Where("username = ?", user.Username).First(&existingUser)

    if existingUser.ID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid username or password",
        })
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Invalid username or password",
        })
        return
    }

    token, _ := utils.GenerateToken(user.Username)

    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}

```

routes/authRoutes.go:

```go
package routes

import (
    "github.com/gin-gonic/gin"

    "project/controllers"
)

func AuthRoutes(route *gin.Engine) {
    auth := route.Group("/auth")
    {
        auth.POST("/register", controllers.Register)
        auth.POST("/login", controllers.Login)
    }
}

```

main.go:

```go
package main

import (
    "github.com/gin-gonic/gin"

    "project/config"
    "project/routes"
)

func main() {
    r := gin.Default()

    // Load environment variables from .env file
    err := godotenv.Load(".env")

    if err != nil {
            log.Fatal("Error loading .env file")
        }

    db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Failed to connect to database")
    }

    // Migrate the schema
    db.AutoMigrate(&models.User{})

    // Initialize routes
    routes.AuthRoutes(r)

    r.Run(":8080")
}

```


is there anyway to connect to db one time and use that in every controller??

Yes, you can create a package-level variable in the `config` package to hold the database connection and initialize it once at the start of the application. Then you can use this variable in each controller function that needs to interact with the database.

For example, in `config/config.go` you can create a package-level variable `DB` like this:

```go
package config

import (
    "fmt"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() error {
    // Load configuration from environment variables
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Create database connection
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    return nil
}

```

Then in each controller function, you can use the `DB` variable to interact with the database, for example:

```go
package controllers

import (
    "net/http"

    "project/config"
    "project/models"
)

func GetUser(c *gin.Context) {
    id := c.Param("id")

    var user models.User
    if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

```

Note that this approach assumes that your application is not running in a `concurrent environment`, because `DB` is a global variable that is shared by all requests. If your application is running in a `concurrent environment`, you may need to use a connection pool or some other mechanism to ensure that each request gets its own database connection.



------------
--------------
------------
---------

how to use this in `concurrent environment`,

To use the `DB` variable in a concurrent environment, you should create a new database connection for each request. One way to do this is to use a connection pool. A connection pool is a collection of database connections that can be reused by multiple requests, instead of creating a new connection for each request.

One popular connection pool library for Golang is `sqlx`, which is built on top of the standard library's `database/sql` package. Here's an example of how you could use `sqlx` to create a connection pool:

```go
package config

import (
    "fmt"
    "os"

    "github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDB() error {
    // Load configuration from environment variables
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Create database connection pool
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
    var err error
    DB, err = sqlx.Open("mysql", dsn)
    if err != nil {
        return err
    }

    // Test the connection
    if err := DB.Ping(); err != nil {
        return err
    }

    // Set connection pool options
    DB.SetMaxIdleConns(10)
    DB.SetMaxOpenConns(100)

    return nil
}

```

In the above example, we're using the `sqlx.Open` function to create a new database connection pool. The `SetMaxIdleConns` and `SetMaxOpenConns` methods are used to set the maximum number of idle and open connections in the pool, respectively.

To use the `DB` variable in a controller function, you can create a new database connection from the pool like this:

```go
package controllers

import (
    "net/http"

    "project/config"
    "project/models"
)

func GetUser(c *gin.Context) {
    id := c.Param("id")

    db := config.DB.MustBegin()
    defer db.Close()

    var user models.User
    if err := db.Get(&user, "SELECT * FROM users WHERE id = ?", id); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

```

In this example, we're using the `MustBegin` method to create a new database connection from the pool, and the `defer db.Close()` statement to release the connection back to the pool when the function returns. Note that we're using the `Get` method of the `sqlx.DB` object to execute the database query and retrieve the `User` object.

By using a connection pool and creating a new database connection for each request, we can ensure that our application is able to handle concurrent requests without running out of database connections.


```go
 func Register(c \*gin.Context) {
    ...
    db, err := config.ConnectDB() 
    ...
     db.Create(&user) 
     ...

     } 
```

 in the above code does db, err := config.ConnectDB() stablish a new connection?

Yes, `db, err := config.ConnectDB()` will establish a new connection to the database, assuming that `ConnectDB()` returns a new connection object every time it is called.

If `ConnectDB()` returns a connection pool object, then calling it multiple times will not create new connections every time, but will reuse existing connections in the pool.

In general, it's a good idea to use a connection pool to handle database connections in a web application, since creating new connections for each request can be slow and can lead to resource exhaustion.


how to code config.connecDB to return a connection pool object??

To modify `config.ConnectDB()` to return a connection pool object, you can use the `sql.DB` type from the `database/sql` package, which provides connection pooling by default. Here's an example implementation:

```go
package config

import (
    "database/sql"
    "fmt"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
    // Load configuration from environment variables
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Create database connection pool
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Set connection pool options
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)

    return db, nil
}

```

In this implementation, we're using `sql.Open` to create a new connection pool object, and returning it to the caller. We've also added calls to `SetMaxIdleConns` and `SetMaxOpenConns` to set the maximum number of idle and open connections in the pool, respectively.

With this implementation, calling `config.ConnectDB()` multiple times will return a new connection pool object each time. When you're finished using the connection pool, you should close it using the `db.Close()` method. You should also call `db.Ping()` to test the connection before using it, since the pool may return a closed or invalid connection if the database server goes down.



how to use gorm with database/sql

You can use GORM with `database/sql` by passing a `*sql.DB` connection to `gorm.Open()`. Here's an example implementation:

```go
package config

import (
    "database/sql"
    "fmt"
    "os"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    // Load configuration from environment variables
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    // Create a new database connection
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    // Initialize GORM with the database connection
    gormDB, err := gorm.Open(mysql.New(mysql.Config{
        Conn: db,
    }), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    return gormDB, nil
}

```

In this implementation, we're creating a new `*sql.DB` connection using `sql.Open()`, testing the connection with `db.Ping()`, and passing the connection to `gorm.Open()` by configuring the `mysql` driver. We're also returning a `*gorm.DB` object instead of a `*sql.DB` object.

With this implementation, you can use GORM with `database/sql` to handle database connections and transactions
