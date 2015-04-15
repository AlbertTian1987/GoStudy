package lib

import (
	"errors"
	"fmt"
)

// Goroutine票池的接口。
type GoTickets interface {
	// 拿走一张票。
	Take()
	// 归还一张票。
	Return()
	// 票池是否已被激活。
	Active() bool
	// 票的总数。
	Total() uint32
	// 剩余的票数。
	Remainder() uint32
}

type myGoTickets struct {
	total   uint32
	tickets chan byte
	active  bool
}

func NewGoTickets(total uint32) (GoTickets, error) {
	this := myGoTickets{}
	if !this.init(total) {
		errMsg := fmt.Sprintf("The goruntine ticket pool can NOT be initialized! (total=%d)\n", total)
		return nil, errors.New(errMsg)
	}
	return &this, nil
}

func (this *myGoTickets) init(total uint32) bool {
	if this.active {
		return false
	}

	if total == 0 {
		return false
	}

	ch := make(chan byte, total)
	n := int(total)
	for i := 0; i < n; i++ {
		ch <- 1
	}

	this.tickets = ch
	this.total = total
	this.active = true

	return true
}

func (this *myGoTickets) Take() {
	<-this.tickets
}

func (this *myGoTickets) Return() {
	this.tickets <- 1
}

func (this *myGoTickets) Active() bool {
	return this.active
}

func (this *myGoTickets) Total() uint32 {
	return this.total
}

func (this *myGoTickets) Remainder() uint32 {
	return uint32(len(this.tickets))
}
