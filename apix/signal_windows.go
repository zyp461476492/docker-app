/*
 * Revision History:
 *     Initial: 2018/10/24        Li Zebang
 */

package server

import (
	"os/signal"
	"syscall"
)

func (ep *EntryPoint) configureSignals() {
	signal.Notify(ep.signals, syscall.SIGINT, syscall.SIGTERM)
}

func (ep *EntryPoint) listenSignals() {
	for {
		sig := <-ep.signals

		switch sig {
		default:
			ep.Stop()
			return
		}
	}
}
