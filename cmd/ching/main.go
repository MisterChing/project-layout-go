package main

import (
	"github.com/MisterChing/go-lib/utils/workflow"
	"log"
	"time"
)

func main() {

	workflow.GoWithRecover(func() {
		for {
			log.Println("main G")
			time.Sleep(time.Second)
			go func() {
				defer func() {
					if err := recover(); err != nil {
						log.Println("fuck recover!")
					}
				}()
				go func() {
					doWork()
					NewPanic()
				}()

			}()

		}
	})

	for {
		log.Println("hahaha")
		time.Sleep(time.Second)
	}

	time.Sleep(time.Second * 1000)
}

func doWork() {

	log.Println("do work")
	time.Sleep(time.Second)
	//NewPanic()

}
func NewPanic() {

	panic("a panic !")

}
