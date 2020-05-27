package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"
)

type InputKey struct {
	Key string `json:"key" biniding:required`
}

type Rules struct {
	Rules []Rule `json:"rules"`
}
type Rule struct {
	ServiceName     string `json:"service_name"`
	RegExpresession string `json:"reg_expression"`
}

func setupRouter() *gin.Engine {
	// gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/static/index.tmpl", gin.H{})
	})

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func cliSecretCheck(c *cli.Context) error {
	secretKey := c.String("key")
	fmt.Printf(secretCheck(secretKey))
	return nil
}

func secretCheck(arg string) string {
	// rulesData, err := os.Open("./static/rules.json")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer rulesData.Close()

	// byteValue, _ := ioutil.ReadAll(rulesData)
	// var rules Rules

	rules := &Rules{}

	json.Unmarshal([]byte(allRules), &rules)

	for i := 0; i < len(rules.Rules); i++ {
		matched, _ := regexp.MatchString(rules.Rules[i].RegExpresession, arg)
		if matched == true {
			return fmt.Sprintf("We found the service name for the key: %v as %v", arg, rules.Rules[i].ServiceName)
		}
	}
	return fmt.Sprintf("We were not able to find the service name for the key: %v", arg)
}

func runServer(c *cli.Context) error {
	r := setupRouter()

	r.POST("/", func(c *gin.Context) {
		var json InputKey
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if json.Key == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Request doesn't have valid key!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": secretCheck(json.Key)})
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8000")
	return nil
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
