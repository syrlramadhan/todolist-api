package util

import "log"

func SendPanicIfError(err error) {
	if err != nil {
		log.Print("err {} ", err.Error())
		panic(err)
	}
}
