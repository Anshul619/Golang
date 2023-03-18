package main

import "log"

type Client struct{}

func (c Client) InsertLightningConnectorIntoComputer(com Computer) {
	log.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
