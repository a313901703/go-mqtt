package help

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/labstack/gommon/log"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func GenerateRandomString(length int) (result string) {
	characters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(characters))
		result += string(characters[randomIndex])
	}
	return
}

func ToCamelCase(s string) string {
	var str string
	cases := cases.Title(language.English)
	for _, v := range strings.Split(s, "_") {
		str += cases.String(v)
	}
	return str
}
func WriteLog(data string, filename string) {
	currentTime := time.Now()
	file, err := os.OpenFile("./logs/"+currentTime.Format("2006-01-02")+"-"+filename+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	switch filename {
	case "error":
	case "panic":
		log.Error(data)
	default:
		log.Info(data)
	}
	file.WriteString(data + "\n")
}

func ErrorLog(data string) {
	WriteLog(data, "error")
}
func PanicLog(data string) {
	WriteLog(data, "panic")
}
