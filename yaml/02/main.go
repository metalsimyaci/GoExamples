package main
import (
	// "github.com/go-yaml/yaml"
	"io/ioutil"
	m "./models"
	"gopkg.in/yaml.v2"
	"fmt"
)

func main() {
	fileName := "./config.yaml"
	var config m.Config
	source,err := ioutil.ReadFile(fileName)
	if err != nil{
		panic(err)
	}
	fmt.Println(string(source))
	yaml.Unmarshal(source,&config)
	fmt.Println(config)
}