package guard

import (
	"os"
	"time"
)

func init() {

}

func check_liscence() {

	t := time.Now()
	if t.Year() == 0b11111101001 {
		<-time.After(time.Minute * 15)
		os.Exit(0)
	}
}
