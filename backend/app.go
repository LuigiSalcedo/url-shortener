package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"log"
	"net/http"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ErrJson = errors.New("json is not valid")
var client *redis.Client

type URL struct {
	Dir string `json:"url"`
}

func saveUrl(c echo.Context) error {
	url := new(URL)

	err := json.NewDecoder(c.Request().Body).Decode(url)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrJson)
	}

	hash := fnv.New32a()
	hash.Write([]byte(url.Dir))

	realUrl := url.Dir

	urlHashed := fmt.Sprintf("%X", hash.Sum(nil))

	url.Dir = fmt.Sprintf("http://localhost:8888/%s", urlHashed)

	err = client.Set(urlHashed, realUrl, time.Hour*24).Err()

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, url)
}

func getUrl(c echo.Context) error {
	hash := c.Param("hash")

	dir, err := client.Get(hash).Result()

	if err == redis.Nil || dir == "" {
		return c.HTML(http.StatusNotFound, "not found")
	}

	return c.Redirect(http.StatusSeeOther, dir)
}

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if client == nil {
		log.Fatal("redis db not detected")
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST"},
	}))

	e.POST("/save", saveUrl)
	e.GET("/:hash", getUrl)

	log.Fatal(e.Start(":8888"))
}
