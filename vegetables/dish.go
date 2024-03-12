package vegetables

import (
	"fmt"
	"time"
)

func Meat_Dish(msg string) {
	t := time.Now()
	fmt.Printf("%s,开始炒 %s \n", t.Format("2006-01-02 15:04:05"), msg)
}
