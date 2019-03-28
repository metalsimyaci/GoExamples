package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Kullanıcı oluşturma alanı v1")
	user1 := &User{
		Id:			1,
		FirstName:  "Hasan",
		LastName:	"URAL",
		UserName:	"hasan.ural",
		Age:		30,
		Pay:		&Payment{
			Salary:		4500,
			Bonus:		2000,
		},
	}

	fmt.Println(user1.GetFullName())
	fmt.Println(user1.GetPayment())
	fmt.Println("Maaş:"+strconv.FormatFloat(user1.Pay.Salary,'g',-1,64))

	fmt.Println("Kullanıcı oluşturma alanı v2")
	user2 :=NewUser()
	user2.FirstName ="Hakan"
	user2.LastName="URAL"
	user2.UserName="hakan.ural"
	user2.Age=29
	user2.Pay=&Payment{775,950}
	fmt.Println(user2.GetFullName())
	fmt.Println(user2.GetPayment())
	fmt.Println("Maaş:"+strconv.FormatFloat(user2.Pay.Salary,'g',-1,64))
}
type User struct {
	Id 			int
	FirstName	string
	LastName	string
	UserName	string
	Age			int
	Pay			*Payment
}
type Payment struct {
	Salary	float64
	Bonus	float64
}

func NewUser() *User  {
	u :=new(User)
	u.Pay = NewPayment()
	return u
}
func NewPayment() *Payment  {
	p :=new(Payment)
	return p
}
func (u User) GetFullName() string  {
	return u.FirstName+" "+u.LastName
}
func (u *User) GetUserName() string  {
	return u.UserName
}
func (u *User) GetPayment() float64  {
	return u.Pay.Salary+u.Pay.Bonus
}

