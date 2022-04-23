package main

import (
	"github.com/jinfwhuang/ds-toolkit/pkg/login"
	log "log"
)

func init() {
	log.SetFlags(log.Llongfile)
}

/*

 */
func main() {
	//ds.Put()

	login.GenPrivateKey()
	log.Println("finish debug main")
}






