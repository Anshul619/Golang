package main

import "log"

type Windows struct{}

func (w *Windows) insertIntoUSBPort() {
	log.Println("USB connector in Windows")
}
