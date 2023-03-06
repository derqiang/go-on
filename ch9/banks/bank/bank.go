package bank

import "sync"

var balance int

type Bank int64

func (Bank) Deposit(amount int) {
	balance = balance + amount
}

func (Bank) Balance() int {
	return balance
}

// + build ignore

type Bank1 int64

var deposits = make(chan int)
var balances = make(chan int)

func (Bank1) Deposit(amount int) {
	deposits <- amount
}

func (Bank1) Balance() int {
	return <-balances
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
}

// - build ignore

var sema = make(chan struct{}, 1)

type Bank3 int64

func (Bank3) Deposit(amount int) {
	sema <- struct{}{}
	balance = balance + amount
	<-sema
}

func (Bank3) Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

/////

var mu sync.Mutex

type Bank4 int64

func (Bank4) Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}

func (Bank4) Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
