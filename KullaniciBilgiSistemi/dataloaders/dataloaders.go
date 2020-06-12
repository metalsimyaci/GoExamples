package dataloaders

import (
	. "../models"
	"../utils"
	"encoding/json"
	"path/filepath"
)

const (
	userPath = "./json/users.json"
	interestPath = "./json/interests.json"
	interestMappingPath = "./json/userInterestMappings.json"
)

func LoadUsers() []User{
	var result []User
	absPath, _ := filepath.Abs(userPath)
	//absPath,_ := os.Getwd()
	//path := absPath+userPath
	bytes,err := utils.ReadFile(absPath)
	utils.CheckError(err)

	err =json.Unmarshal([]byte(bytes),&result)
	utils.CheckError(err)

	return result
}
func LoadInterests() []Interest{
	var result []Interest
	absPath, _ := filepath.Abs(interestPath)
	bytes,err := utils.ReadFile(absPath)
	utils.CheckError(err)

	err =json.Unmarshal([]byte(bytes),&result)
	utils.CheckError(err)

	return result
}
func LoadInterestMappings() []InterestMapping{
	var result []InterestMapping
	absPath, _ := filepath.Abs(interestMappingPath)
	bytes,err := utils.ReadFile(absPath)
	utils.CheckError(err)

	err =json.Unmarshal([]byte(bytes),&result)
	utils.CheckError(err)

	return result
}
