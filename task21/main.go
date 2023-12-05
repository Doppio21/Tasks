package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type Market struct {
}

func (m *Market) SendOrder() {
	fmt.Println("market send the order")
}

type Client struct {
}

func (c *Client) GetOrder() {
	fmt.Println("client get the order")
}

type OrderAdapter interface{
	Actions()
}

type MarketAdapter struct{
	*Market
}

func(adapter *MarketAdapter) Actions(){
	adapter.SendOrder()
}

func NewMarketAdapter(m *Market) OrderAdapter{
	return &MarketAdapter{m}
}

type ClientAdapter struct{
	*Client
}

func(adapter *ClientAdapter) Actions(){
	adapter.GetOrder()
}

func NewClientAdapter(c *Client) OrderAdapter{
	return &ClientAdapter{c}
}

type Order struct{
}

func(o *Order) Actions(){
	fmt.Println("order is delivered")
}

func main() {
	order:=[]OrderAdapter{NewMarketAdapter(&Market{}), NewClientAdapter(&Client{}), &Order{}}

	for _, j := range order{
		j.Actions()
	}
}
