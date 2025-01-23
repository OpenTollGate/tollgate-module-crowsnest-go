package main

import (
	"errors"
	"fmt"
	"github.com/digineo/go-uci"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func stringToHex(s string) string {
	hexStr := ""
	for _, b := range []byte(s) {
		hexStr += fmt.Sprintf("%02x", b)
	}
	return hexStr
}

const oiu = "212121"     // TollGate custom OIU
const elementType = "01" // TollGate custom elementType
func createVendorElement(payload string) (string, error) {

	if len(payload) > 247 {
		return "", errors.New("payload cannot exceed 247 characters to stay within vendor_elements max size of 256 chars")
	}

	var vendorElementPayload = oiu + elementType + payload

	payloadLengthInBytesHex := strconv.FormatInt(int64(len(vendorElementPayload)), 16)
	payloadHex := stringToHex(vendorElementPayload)

	return "dd" + payloadLengthInBytesHex + payloadHex, nil
}

func setupBroadcast() {
	interfaces := []string{"default_radio0", "default_radio1"} // TODO: from config

	// TODO from config:		Pubkey of: merchant nsec17jlyx05kfqpyhrfuu6450x8shzlaslpngjnr8fe27raacmp49tzsvfaz9v
	var tollgateVersion = "v0.1.0"
	var gatewayIp = "192.168.1.1"
	var pubkey = "c1f4c025e746fd307203ac3d1a1886e343bea76ceec5e286c96fb353be6cadea"
	var allocationType = "KiB" // or min
	var priceAllocationPer1024 = "1049000"
	var priceUnit = "sat"

	salesPitch := []string{tollgateVersion, pubkey, allocationType, priceAllocationPer1024, priceUnit, gatewayIp}
	var salesPitchString = strings.Join(salesPitch, "|")

	log.Println(salesPitchString)
	vendorElement, err := createVendorElement(salesPitchString)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = setVendorElements(interfaces, vendorElement)
	reloadWifi()

	if err != nil {
		log.Fatal(err)
	}
}

func reloadWifi() {
	log.Println("reloading wifi")

	cmd := exec.Command("wifi", "reload")
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err, "Error reloading wifi")
	} else {
		log.Println("wifi reloaded")
	}
}

func main() {
	log.Println("Starting Tollgate - CrowsNest")

	setupBroadcast()

	log.Println("Shutting down Tollgate - CrowsNest")

	//ifaces := wireless.Interfaces()
	//fmt.Println(ifaces)
	//
	//wifis, err := wifiscan.Scan()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for _, w := range wifis {
	//	fmt.Println(w.SSID, w.RSSI)
}

func setVendorElements(interfaces []string, vendorElement string) error {
	for _, iface := range interfaces {
		if vendorElement == "" {
			return errors.New("vendorElement cannot be empty string")
		}

		u := uci.NewTree("/etc/config")
		u.Del("wireless", iface, "vendor_elements")

		log.Println("add element to inteface " + iface + ", element: " + vendorElement)
		if ok := u.Set("wireless", iface, "vendor_elements", vendorElement); !ok {
			return errors.New("could not add vendor_elements to interface " + iface + ": " + vendorElement)
		}

		// Save edits
		err := u.Commit()
		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	return nil
}

//func setWifiSSID() {
//
//	var availableInterface = getAvailableWirelessInterfaceName()
//	fmt.Println(availableInterface)
//
//	// get respective modes
//
//	// take the one with 'ap'
//	iwlistCmd := exec.Command("iwlist", availableInterface, "scan")
//	iwlistCmdOut, err := iwlistCmd.Output()
//	if err != nil {
//		fmt.Println(err, "Error when getting the interface information.")
//	} else {
//		fmt.Println(string(iwlistCmdOut))
//	}
//}
//
//func getAvailableWirelessInterfaceName() string {
//	var allInterfaces = getWifiInterfaces()
//
//	var found []string
//
//	scanner := bufio.NewScanner(strings.NewReader(allInterfaces))
//	for scanner.Scan() {
//		var line = scanner.Text()
//		if strings.HasSuffix(line, "wifi-iface") {
//
//			var interfaceName string = strings.Split(line, "=")[0]
//
//			found = append(found, interfaceName)
//		}
//	}
//
//	for _, interfaceName := range found {
//		var search string = "wireless." + string(interfaceName) + ".mode='ap'"
//
//		scanner := bufio.NewScanner(strings.NewReader(allInterfaces))
//		for scanner.Scan() {
//			var line = scanner.Text()
//
//			if line == search {
//				return interfaceName
//			}
//		}
//	}
//
//	return "?"
//}
//
//func getWifiInterfaces() string {
//	iwlistCmd := exec.Command("uci", "show", "wireless")
//	iwlistCmdOut, err := iwlistCmd.Output()
//	if err != nil {
//		fmt.Println(err, "Error when getting the interface information.")
//	} else {
//		return string(iwlistCmdOut)
//	}
//
//	return "?"
//}
