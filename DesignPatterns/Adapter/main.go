package main

func main() {
	client := new(Client)
	mac := new(Mac)

	client.InsertLightningConnectorIntoComputer(mac)

	windowsMachine := new(Windows)

	windowsMachineAdapter := new(WindowsAdapter)
	windowsMachineAdapter.windowsMachine = windowsMachine

	client.InsertLightningConnectorIntoComputer(windowsMachineAdapter)
}
