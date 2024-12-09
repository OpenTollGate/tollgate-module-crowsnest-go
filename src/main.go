package main

import (
	"bufio"
	"fmt"
	"github.com/theojulienne/go-wireless"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")

	ifaces := wireless.Interfaces()
	fmt.Println(ifaces)

	//wc, _ := wireless.NewClient(ifaces[1])
	//defer wc.Close()
	//
	//conn, err := wireless.Dial(ifaces[0])
	//if err != nil {
	//	fmt.Printf("failed to dial network: %s\n", err)
	//	os.Exit(1)
	//}

	wc, err := wireless.NewClient(ifaces[1])
	if err != nil {
		fmt.Printf("failed to create client: %s\n", err)
		os.Exit(1)
	}
	defer wc.Close()
	//fmt.Println(ifaces)
	//fmt.Println(aps)
	//fmt.Println(ap)
	//fmt.Println(ok)

	//setWifiSSID()
	//doNostrStuff()

	fmt.Println("Bye, world!")
}

func setWifiSSID() {

	var availableInterface = getAvailableWirelessInterfaceName()
	fmt.Println(availableInterface)

	// get respective modes

	// take the one with 'ap'
	iwlistCmd := exec.Command("iwlist", availableInterface, "scan")
	iwlistCmdOut, err := iwlistCmd.Output()
	if err != nil {
		fmt.Println(err, "Error when getting the interface information.")
	} else {
		fmt.Println(string(iwlistCmdOut))
	}
}

func getAvailableWirelessInterfaceName() string {
	var allInterfaces = getWifiInterfaces()

	var found []string

	scanner := bufio.NewScanner(strings.NewReader(allInterfaces))
	for scanner.Scan() {
		var line = scanner.Text()
		if strings.HasSuffix(line, "wifi-iface") {

			var interfaceName string = strings.Split(line, "=")[0]

			found = append(found, interfaceName)
		}
	}

	for _, interfaceName := range found {
		var search string = "wireless." + string(interfaceName) + ".mode='ap'"

		scanner := bufio.NewScanner(strings.NewReader(allInterfaces))
		for scanner.Scan() {
			var line = scanner.Text()

			if line == search {
				return interfaceName
			}
		}
	}

	return "?"
}

func getWifiInterfaces() string {
	iwlistCmd := exec.Command("uci", "show", "wireless")
	iwlistCmdOut, err := iwlistCmd.Output()
	if err != nil {
		fmt.Println(err, "Error when getting the interface information.")
	} else {
		return string(iwlistCmdOut)
	}

	return "?"
}
