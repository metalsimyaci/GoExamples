package main

import (
	"fmt"
	"strconv"
)

func main() {
	ferrari := NewFerrari("f1","red",320,900000,true)
	CarExecute(ferrari)
	//fmt.Println(ferrari.Information())

	lamborghini := NewLamborghini("l1","blue",300,896000,false)
	CarExecute(lamborghini)
	//fmt.Println(lamborghini.Information())

	mercedes := NewMercedes("l1","blue",360,970000)
	CarExecute(mercedes)
	//fmt.Println(mercedes.Information())

}
func CarExecute(c Carface){
	fmt.Println("\n"+"Araç Bilgi:"+"\n"+c.Information()+"\n")
	var msg string
	isRun :=c.Run()
	if isRun{
		msg="Çalışıyor"
	}else{
		msg="Çalışmıyor"
	}
	fmt.Println("Araç Çalışması:"+msg+".")

	isStop := c.Stop()
	if isStop{
		msg = "durdu"
	}else{
		msg ="duramadı!"
	}
	fmt.Println("Araç durma durumu:"+msg+".")
}


//interface
type Carface interface {
	Run() bool
	Stop() bool
	Information() string
}


type Car struct {
	Brand	string
	Model	string
	Color	string
	Speed 	int
	Price 	float64
}
type SpecialProduction struct {
	Special bool
}

//Ferrari
type Ferrari struct {
	Car
	SpecialProduction
}
func NewFerrari(model, color string, speed int, price float64, isSpecial bool) *Ferrari{
	f :=new(Ferrari)
	f.Brand ="Ferrari"
	f.Model= model
	f.Color= color
	f.Speed= speed
	f.Price= price
	f.Special=isSpecial
	return f
}
func (_ Ferrari) Run() bool  {
	return true
}
func (_ Ferrari) Stop() bool  {
	return true
}
func (f *Ferrari) Information() string  {
	r := "\t"+f.Brand+" "+f.Model+"\t"+"Color:"+f.Color+"\n"+"\t"+"Speed:"+strconv.Itoa(f.Speed)+"\n\t"+"Price:"+strconv.FormatFloat(f.Price,'g',-1,64)
	r +="\n\t"+"Special:"
	if f.Special{
		r+="YES"
	}else{
		r +="NO"
	}
	return r
}

//Lamborghini
type Lamborghini struct {
	Car
	SpecialProduction
}
func NewLamborghini(model, color string, speed int, price float64, isSpecial bool) *Lamborghini{
	f :=new(Lamborghini)
	f.Brand ="Lamborghini"
	f.Model= model
	f.Color= color
	f.Speed= speed
	f.Price= price
	f.Special=isSpecial
	return f
}
func (_ Lamborghini) Run() bool  {
	return false
}
func (_ Lamborghini) Stop() bool  {
	return true
}
func (f *Lamborghini) Information() string  {
	r := "\t"+f.Brand+" "+f.Model+"\t"+"Color:"+f.Color+"\n"+"\t"+"Speed:"+strconv.Itoa(f.Speed)+"\n\t"+"Price:"+strconv.FormatFloat(f.Price,'g',-1,64)
	r +="\n\t"+"Special:"
	if f.Special{
		r+="YES"
	}else{
		r +="NO"
	}
	return r
}

//Mercedes
type Mercedes struct {
	Car
}
func NewMercedes(model, color string, speed int, price float64) *Lamborghini{
	f :=new(Lamborghini)
	f.Brand ="Mercedes"
	f.Model= model
	f.Color= color
	f.Speed= speed
	f.Price= price
	return f
}
func (_ Mercedes) Run() bool  {
	return true
}
func (_ Mercedes) Stop() bool  {
	return true
}
func (f *Mercedes) Information() string  {
	r := "\t"+f.Brand+" "+f.Model+"\t"+"Color:"+f.Color+"\n"+"\t"+"Speed:"+strconv.Itoa(f.Speed)+"\n\t"+"Price:"+strconv.FormatFloat(f.Price,'g',-1,64)

	return r
}

