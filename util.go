package loger

import (
	"fmt"
	"log"
	"os"
	"strings"
)
// Описание цветов ANSII
//http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
var StdLog = log.New(os.Stderr, "", log.Llongfile)
//var Stdout = log.New(os.Stdout, "", log.Llongfile)
func Инфо(формат string,  данные ...interface{}) {
	StdLog.SetFlags(log.Lshortfile|log.Ltime)
	формат = strings.ReplaceAll(формат, "%+v", "\u001b[38;5;48m %+v  \u001b[0m\u001b[38;5;75m")

	str := fmt.Sprintf("\u001b[0m\u001b[36mИНФО: \u001b[38;5;75m " + формат +" \u001b[0m \n", данные...)
	err := StdLog.Output(2, str)

	if err != nil {
		log.Printf("%+v", err)
	}

}

func Ошибка(формат string,  данные ...interface{}) {
	StdLog.SetFlags(log.Lshortfile|log.Ltime)
	формат = strings.ReplaceAll(формат, "%+v", "\u001b[38;5;204m %+v  \u001b[0m\u001b[38;5;1m")

	str := fmt.Sprintf("\u001b[48;5;124m ОШИБКА >> \u001b[0m  \u001b[38;5;1m  " + формат +" \u001b[0m \n", данные...)

	err := StdLog.Output(2, str)
	if err != nil {
		log.Printf("%+v", err)
	}
}
