package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:0,
	}
	return b
}

//receiver function - format bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0.0


	//list items 
	for key, value := range b.items{
		fs += fmt.Sprintf("%-25v ...$%v\n", key+":", value)
		total +=value
	}

	total += b.tip

	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "Tip: ", b.tip)
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total: ", total)


	return fs
}

//udpate tip

func (b *bill) updateTip(tip float64) {
	(*b).tip += tip
}

//add item
func (b *bill) addItem(name string, price float64){
	b.items[name] = price
}

//save bill
func (b *bill) save(){
	data := []byte(b.format())
	err := os.WriteFile("bill/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file")
}