package gfull

import (
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var onCloseFunc []func()
var onCloseEnable = false

// AddOnClose f
func AddOnClose(f func()) {
	if onCloseEnable == false {
		onCloseEnable = true
		// Close Database
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
		go func() {
			<-sigchan
			println("\nGFULL: Bye...")
			// recorres las funciones
			for ind := range onCloseFunc {
				onCloseFunc[ind]()
			}
			os.Exit(0)
		}()
	}
	onCloseFunc = append(onCloseFunc, f)
}

// StrToUint : converte string in uint
func StrToUint(val string) (uint, error) {
	n, err := strconv.ParseUint(val, 0, 64)
	return uint(n), err
}
