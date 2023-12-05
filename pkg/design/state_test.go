package design

import (
	"log"
	"testing"
)

func TestVendingMachine(t *testing.T) {
	vendingMachine := NewVendingMachine(1, 10)
	err := vendingMachine.requestItem()
	if err != nil {
		log.Fatal(err)
	}
	err = vendingMachine.insertMoney(10)
	if err != nil {
		log.Fatal(err)
	}
	err = vendingMachine.dispenseItem()
	if err != nil {
		log.Fatal(err)
	}
	err = vendingMachine.addItem(2)
	if err != nil {
		log.Fatal(err)
	}
}
