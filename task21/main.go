package main 

import (
	"fmt"
)

type DataSender interface { //target interface which we would like to implement
	getData()
}

type Usb struct {

}

func (u *Usb) getData() { //a struct which already implement target interface
	fmt.Println("Get data by USB")
}

type Thunderbolt struct { //adaptee. a struct which can't implement target interface directly
}

type PCIExpress struct { //adaptee. a struct which can't implement target interface directly

}

type NewEraConn interface { //adaptee interface (both adaptee implement this interface)
	getVeryFastData()
}

type NewEraConnAdapter struct { //adapter - let connect target interface with adaptee structures
	conn NewEraConn
}

func (adapter *NewEraConnAdapter) getData() {
	fmt.Println("Call adapter to get data from new era")
	adapter.conn.getVeryFastData()
}

func (t *Thunderbolt) getVeryFastData() {
	fmt.Println("Get data by thunderbolt")
}

func (pe *PCIExpress) getVeryFastData() {
	fmt.Println("Get data by PCI express")
}


func main() {
	/*
	Реализовать паттерн «адаптер» на любом примере.
	Зачем нужен адаптер -
	1) Совместить легаси с новым кодом без рефакторинга 
	2) Интеграция с внешним API (т.е когад есть несовместимые интерфейсы)
	3) Поддержка нескольких версий интерфейса
	Адаптер можно реализовывать через агрегацию или композицию, реализовано через агрегацию, как более prod-like вариант, 
	который более гибок за счет сохранения 2 SOLID принципов (Open/closed, Interface segregation), чего нет в композиции 
	*/
	senders := []DataSender{
		&Usb{}, 
		&NewEraConnAdapter{conn: &Thunderbolt{}},
		&NewEraConnAdapter{conn: &PCIExpress{}},
	}

	for _, sender := range senders {
		sender.getData()
	}
}
