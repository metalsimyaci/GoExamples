package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var fileFolderPath ="files\\"
var files = []string{fileFolderPath+"demo.go",fileFolderPath+"note1.txt"}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
func addZIPFile(fileName string,zw *zip.Writer) error{
	file,err := os.Open(fileName)
	checkErr(err)
	defer file.Close()

	wr, err :=zw.Create(fileName)
	checkErr(err)

	if _,err := io.Copy(wr,file); err != nil{
		return fmt.Errorf("%s dosyası ZPI'e yazılırken bir hata oluştu:%s",fileName,err)
	}

	return nil
}
func addTARFile(fileName string,tw *tar.Writer) error{
	file,err := os.Open(fileName)
	checkErr(err)
	defer file.Close()

	stat, err :=file.Stat()
	checkErr(err)

	hdr := &tar.Header{
		ModTime: stat.ModTime(),
		Name: stat.Name(),
		Size: stat.Size(),
		Mode:int64(stat.Mode().Perm()),
	}

	err = tw.WriteHeader(hdr)
	checkErr(err)

	copied, err := io.Copy(tw,file)
	checkErr(err)

	if copied < stat.Size(){
		msg := "%s dosyasına %d kadar veri yazıldı. Ama beklenen veri %d kadardı."
		return fmt.Errorf(msg,fileName,copied,stat)
	}

	return nil
}
func createArchiveTARFile(archiveTarFileName string) int{
	if len(archiveTarFileName) == 0{
		return -1
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveTarFileName+".tar",flags,0644)
	checkErr(err)
	defer file.Close()

	tw := tar.NewWriter(file)
	defer tw.Close()

	for _,fileName := range files{
		if error  := addTARFile(fileName,tw); error != nil{
			log.Fatal("%s dosyası TAR dosyasına eklenemedi",fileName,error)
		}
	}
	return 1
}
func createArchiveZIPFile(archiveTarFileName string) int{
	if len(archiveTarFileName) == 0{
		return -1
	}

	flags := os.O_WRONLY | os.O_CREATE | os.O_TRUNC
	file, err := os.OpenFile(archiveTarFileName+".zip",flags,0644)
	checkErr(err)
	defer file.Close()

	zw := zip.NewWriter(file)
	defer zw.Close()

	for _,fileName := range files{
		if error  := addZIPFile(fileName, zw); error != nil{
			log.Fatal("%s dosyası ZIP dosyasına eklenemedi",fileName,error)
		}
	}
	return 1
}

func CreateTar(){
	result := createArchiveTARFile("dosyaX")
	if result > 0 {
		fmt.Println("İşlem Başarılı:", result)
	}else{
		fmt.Println("İşlem Başarısız oldu:",result)
	}
}
func CreateZip(){
	result := createArchiveZIPFile("dosyaX")
	if result > 0 {
		fmt.Println("İşlem Başarılı:", result)
	}else{
		fmt.Println("İşlem Başarısız oldu:",result)
	}
}

func CreateDirIfNotExists(dir string){
	if _,err :=os.Stat(dir); os.IsNotExist(err){
		err =os.MkdirAll(dir,0755)
		checkErr(err)
	}
}
func zipRead(){
	zr,err :=zip.OpenReader("dosyaX.zip")
	checkErr(err)

	defer zr.Close()

	for _, file := range zr.Reader.File{
		zippedFile,err := file.Open()
		checkErr(err)
		defer zippedFile.Close()
		targetDir := "./unzip/"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		dirName := strings.Split(extractedFilePath,"\\")
		CreateDirIfNotExists(targetDir+dirName[1])

		if file.FileInfo().IsDir(){
			os.MkdirAll(extractedFilePath,file.Mode())
		}else{
			log.Println("Dosya çıkarılıyor",file.Name)

			outFile, err := os.OpenFile(extractedFilePath,
				os.O_WRONLY | os.O_CREATE | os.O_TRUNC,file.Mode())
			checkErr(err)
			defer outFile.Close()
		}

		fmt.Println(extractedFilePath)

	}
}
func main() {
	zipRead()
}

