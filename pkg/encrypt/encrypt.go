package ds

import (
	tmplog "log"
)

func init() {
	tmplog.SetFlags(tmplog.Llongfile)
}

func Put() {
	tmplog.Println("ds put")
}






