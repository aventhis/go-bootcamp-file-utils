package main

import "flag"

func main() {
	flagA := flag.Bool("a", false, "flag indicating the directory for archives")
	flag.Parse()

}
