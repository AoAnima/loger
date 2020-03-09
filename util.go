package loger

import (
	"fmt"
	"os"
	"strings"
)
// Описание цветов ANSII
//http://www.lihaoyi.com/post/BuildyourownCommandLinewithANSIescapecodes.html
func инфо(формат string,  данные ...interface{}) {

	формат = strings.ReplaceAll(формат, "%+v", "\u001b[38;5;48m %+v  \u001b[0m\u001b[38;5;75m")

	str := fmt.Sprintf("\u001b[0m \u001b[36m ИНФО: \u001b[38;5;75m " + формат +" \u001b[0m \n", данные...)
	_, _ = os.Stdout.WriteString(str)
	//if err != nil {
	//	log.Printf("err %+v\n", err)
	//}
}
func ошибка(формат string,  данные ...interface{}) {
	формат = strings.ReplaceAll(формат, "%+v", "\u001b[38;5;204m %+v  \u001b[0m\u001b[38;5;1m")

	str := fmt.Sprintf("\u001b[48;5;124m ОШИБКА: \u001b[0m  \u001b[38;5;1m  " + формат +" \u001b[0m \n", данные...)

	_, _ = os.Stdout.WriteString(str)
	//if err != nil {
	//	log.Printf("err %+v\n", err)
	//}
}

