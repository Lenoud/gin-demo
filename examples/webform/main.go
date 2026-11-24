package main

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//go:embed templates/*.html static/*
var pages embed.FS

func newRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.SetHTMLTemplate(template.Must(template.ParseFS(pages, "templates/*.html")))
	static, err := fs.Sub(pages, "static")
	if err != nil {
		panic(err)
	}
	router.StaticFS("/static", http.FS(static))

	var accounts sync.Map
	showLogin := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	}
	router.GET("/", showLogin)
	router.GET("/login", showLogin)
	router.GET("/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", nil)
	})

	router.POST("/register", func(c *gin.Context) {
		username, password, ok := readCredentials(c)
		if !ok {
			c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "请输入用户名和密码，密码长度最多为 72 字节。"})
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "register.html", gin.H{"error": "注册暂时失败，请重试。"})
			return
		}
		if _, exists := accounts.LoadOrStore(username, hash); exists {
			c.HTML(http.StatusConflict, "register.html", gin.H{"error": "该用户名已注册。"})
			return
		}
		c.Redirect(http.StatusSeeOther, "/login")
	})

	router.POST("/login", func(c *gin.Context) {
		username, password, ok := readCredentials(c)
		hash, exists := accounts.Load(username)
		if !ok || !exists || bcrypt.CompareHashAndPassword(hash.([]byte), []byte(password)) != nil {
			c.HTML(http.StatusUnauthorized, "index.html", gin.H{"error": "用户名或密码错误。"})
			return
		}
		c.HTML(http.StatusOK, "index.html", gin.H{"message": "登录校验成功。"})
	})
	return router
}

func readCredentials(c *gin.Context) (string, string, bool) {
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, 4096)
	if err := c.Request.ParseForm(); err != nil {
		return "", "", false
	}
	username := strings.TrimSpace(c.Request.PostForm.Get("username"))
	password := c.Request.PostForm.Get("password")
	return username, password, username != "" && len(username) <= 128 && len(password) > 0 && len(password) <= 72
}

func main() {
	if err := newRouter().Run("127.0.0.1:38081"); err != nil {
		log.Fatal(err)
	}
}
