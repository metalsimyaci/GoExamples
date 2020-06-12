package main
import (
	"fmt"
	"github.com/go-yaml/yaml"
)

func main(){
	p := Person{"Hasan","URAL",30}
	y, err := yaml.Marshal(p)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(y))
}
// Person :This is a Person
type Person struct {
	FirstName string `yaml:"first_name"`
	LastName string `yaml:"last_name"`
	Age int `yaml:"age"`
}