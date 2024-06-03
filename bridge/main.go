package main

import "fmt"

// abstraction
type IStore interface {
	pay(float32)
	setPaymentMethod(IPaymentMethod)
}

// Refined abstraction
type Storelocal struct {
	payment IPaymentMethod
}

func (s Storelocal) pay(price float32) {
	s.payment.ExecutePayment(price)
}

func (s *Storelocal) setPaymentMethod(paymentMethod IPaymentMethod) {
	s.payment = paymentMethod
}

// Refined abstraction
type StoreMall struct {
	payment IPaymentMethod
}

func (s StoreMall) pay(price float32) {
	s.payment.ExecutePayment(price)
}

func (s *StoreMall) setPaymentMethod(paymentMethod IPaymentMethod) {
	s.payment = paymentMethod
}

// implementation
type IPaymentMethod interface {
	ExecutePayment(float32)
}

// concrete implementation
type PaymentCard struct{}

func (p PaymentCard) ExecutePayment(value float32) {
	fmt.Printf("paying with card %0.2f \n", value)
}

// concrete implementation
type PaymentCash struct{}

func (p PaymentCash) ExecutePayment(value float32) {
	fmt.Printf("paying with cash %0.2f \n", value)
}

func main() {

	// local store or mall store can be mixed with different payment methods
	paymentCard := PaymentCard{}
	PaymentCash := PaymentCash{}

	storeLocal1 := Storelocal{} // store local + payment card
	storeLocal1.setPaymentMethod(paymentCard)

	storeLocal2 := Storelocal{} // store local + payment cash
	storeLocal2.setPaymentMethod(PaymentCash)

	storeMall1 := StoreMall{} // store mall + payment card
	storeMall1.setPaymentMethod(paymentCard)

	storeLocal1.pay(40.3)
	storeLocal2.pay(232.4556)
	storeMall1.pay(256.89)

}
