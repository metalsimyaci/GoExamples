package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)
const FileName string = "demo.txt"
const NewPath string  = "./moved/renamedFile.txt"
const CopyPath string = "./copied/copy.txt"
const dataPath string  = "data.csv"
var (
	err 	error
)

func check(e error) {
	if e != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exist.")
		}
		panic(e)
	}
}
func MoveFile()  {
	err := os.Rename(FileName,NewPath)
	check(err)
}
func CopyFile()	  {
	nf, err :=os.OpenFile(FileName,os.O_CREATE | os.O_APPEND,0666)
	defer nf.Close()
	check(err)

	cf, err :=os.OpenFile(CopyPath,os.O_CREATE | os.O_APPEND,0666)
	defer cf.Close()
	check(err)

	_, err2 := io.Copy(cf,nf)
	check(err2)



}
func WriteFile(){
	myFile,err := os.OpenFile(FileName,os.O_CREATE | os.O_APPEND,0666)
	defer  myFile.Close()
	check(err)
	for i := 0; i< 10; i++  {
		content := "writes\n"
		_,err := myFile.WriteString(content)
		check(err)
	}
}
func WriteFileFromByte(){
	myFile,err := os.OpenFile(FileName,os.O_CREATE | os.O_APPEND,0666)
	defer  myFile.Close()
	check(err)

	byteSlice :=[]byte("Bu dosyaya yazılacak")
	_,err = myFile.Write(byteSlice)
	check(err)
}
func WriteFileFromTempDir(){
	tempDirPath, err := ioutil.TempDir("","tempFile")
	check(err)

	tempFile, err := ioutil.TempFile(tempDirPath,"tempFile.txt")
	tempFile.Close()
	check(err)

	fmt.Println(tempFile.Name())

	err = os.Remove(tempFile.Name())
	check(err)
	fmt.Println("Temp Dosya Silindi")

	err = os.Remove(tempDirPath)
	check(err)
	fmt.Println("Temp Klasör Silindi")

}
func FileInfo(){
	fInfo, err := os.Stat(FileName)
	check(err)
	fmt.Printf("File name:%v\n",fInfo.Name())
}
func ReadCsv()  {
	f,err := os.Open(dataPath)
	defer f.Close()
	check(err)

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	check(err)

	for _,row := range  records{
		printRow(row)
	}
}
func printRow(row [] string)  {
	log.Printf("len(row) %d\n",len(row))
	for i,col := range row  {
		log.Printf("[%d]:%s\n",i,col)
	}
}
func main() {
	//WriteFile()
	//WriteFileFromByte()
	//WriteFileFromTempDir()
	//MoveFile()
	//CopyFile()
	ReadCsv()
	log.Println("İşlem tamamlandı")
}
