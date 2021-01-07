package userstory

import "github.com/yanyundata/woodpecker/busybox"

type BusyBox struct {
	session Session
}

func newBusyBox() BusyBox {
	return BusyBox{session: make(Session)}
}

func (bb BusyBox) GrpcHelper() *busybox.GrpcHelper {
	return &busybox.GrpcHelper{}
}

func (bb BusyBox) Session() Session {
	return bb.session
}
