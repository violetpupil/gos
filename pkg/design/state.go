// 状态模式
// https://refactoringguru.cn/design-patterns/state/go/example
// 该模式将与状态相关的行为抽取到独立的状态类中
package design

import "fmt"

// VendingMachine 自动售货机
type VendingMachine struct {
	noItem        State
	hasItem       State
	itemRequested State
	hasMoney      State

	currentState State

	itemCount int
	itemPrice int
}

func NewVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	noItemState := &NoItemState{vendingMachine: v}
	hasItemState := &HasItemState{vendingMachine: v}
	itemRequestedState := &ItemRequestedState{vendingMachine: v}
	hasMoneyState := &HasMoneyState{vendingMachine: v}

	v.setState(hasItemState)
	v.noItem = noItemState
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	return v
}

func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount += count
}

// State 自动售货机在每种状态时，对不同操作的处理
type State interface {
	addItem(int) error     // 添加商品
	requestItem() error    // 选择商品
	insertMoney(int) error // 插入纸币
	dispenseItem() error   // 提供商品
}

// NoItemState 无商品
type NoItemState struct {
	vendingMachine *VendingMachine
}

func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

func (i *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

// HasItemState 有商品
type HasItemState struct {
	vendingMachine *VendingMachine
}

func (i *HasItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	return nil
}

func (i *HasItemState) requestItem() error {
	// 兜底
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Println("Item requested")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

// ItemRequestedState 商品已请求
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less")
	}
	fmt.Println("Money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

// HasMoneyState 收到纸币
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item dispense in progress")
}

func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing item")
	i.vendingMachine.itemCount -= 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
