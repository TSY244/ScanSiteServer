package main

import (
	"ScanSiteServer/poc"
	"ScanSiteServer/scanner"
)

func main() {
	pocs := poc.GetPocs()
	url := "http://www.moonlab.com"
	scanner.Run(url, pocs)
}
