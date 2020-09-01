package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"runtime"
	"strings"
)

const logo = `
 _____ _          ___        _ 
|  __ (_)        / _ \      (_)
| |  \/_ _ __   / /_\ \_ __  _ 
| | __| | '_ \  |  _  | '_ \| |
| |_\ \ | | | | | | | | |_) | |
 \____/_|_| |_| \_| |_/ .__/|_|
                      | |      
                      |_|      `

func welcome(endPoint string) {
	fmt.Println(strings.Replace(logo, "*", "`", -1))
	fmt.Println("")
	fmt.Println(fmt.Sprintf("Server      Name:     %s", os.Getenv("APP_NAME")))
	fmt.Println(fmt.Sprintf("System      Name:     %s", runtime.GOOS))
	fmt.Println(fmt.Sprintf("Go          Version:  %s", runtime.Version()[2:]))
	fmt.Println(fmt.Sprintf("Gin         Version:  %s", gin.Version))
	fmt.Println(fmt.Sprintf("Listen      Address:  %s", endPoint))
}
