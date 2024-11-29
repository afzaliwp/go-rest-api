package helpers

import "log"

func MyLog(v ...interface{}) {
	separator := "----------------------"
	log.Println(separator)
	log.Println(v...)
	log.Println(separator)
}
