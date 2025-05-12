package main

import (
	"bufio"
	"fmt"
	"os"
)

func printDhcpPacker() {

	fmt.Print("\n ")
	fmt.Print("\n0         7            15         23         31")
	fmt.Print("\n+----------------------------------------------+")
	fmt.Print("\n|  Op (1)  |  Htype (1)  | Hlen (1) | Hops (1) |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|                    Xid (4)                   |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|        Secs (2)      |      Flags (2)        |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|                 Ciaddr (4)                   |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|                 Yiaddr (4)                   |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|                 Siaddr (4)                   |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|                 Giaddr (4)                   |")
	fmt.Print("\n+----------------------------------------------+")
	fmt.Print("\n|                 Chaddr (16)                  |")
	fmt.Print("\n+----------------------------------------------+")
	fmt.Print("\n|                 Sname (64)                   |")
	fmt.Print("\n+----------------------------------------------|")
	fmt.Print("\n|                 File (128)                   |")
	fmt.Print("\n+----------------------------------------------+")
	fmt.Print("\n|           Options(variable)                  |")
	fmt.Print("\n+----------------------------------------------+\n\n")

}

func dhcpDiscovery() {
	fmt.Println("DHCP DISCOVERY")
}

func dhcpOffer() {
	fmt.Println("DHCP OFFER")
}

func dhcpRequest() {
	fmt.Println("DHCP REQUEST")
}

func dhcpAcknowledge() {
	fmt.Println("DHCP ACKNOWLEDGE")
}

func main() {
	fmt.Println("DHCP Packet simulation.")
	fmt.Println("Press enter to continue...")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
	fmt.Println("Press enter to clear")
	input.Scan()
	fmt.Print("\033[H\033[2J")
	printDhcpPacker()
	dhcpDiscovery()
	dhcpOffer()
	dhcpRequest()
	dhcpAcknowledge()

}
