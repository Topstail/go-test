package main

func main() {

	host1 := Host{
		num:      3,
		hostname: "icx127",
	}
	println("host1.ipaddr=", host1.ipaddr, "---over")

	println("W33")
}

type Host struct {
	ipaddr   string
	hostname string
	num      int8
}
