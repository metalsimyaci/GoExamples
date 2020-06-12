package handlers

import (
	. "../dataloaders"
	. "../models"
	. "../utils"
	"encoding/json"
	"net/http"
)

func Run(){
	http.HandleFunc("/",handler)
	http.ListenAndServe(":8000",nil)
}
func handler(wr http.ResponseWriter, r *http.Request)  {
	page := Page{ID:7, Name:"Kullanıcılar",Description:"Kullanıcı Listesi", URI:"/users"}

	users := LoadUsers()
	interests := LoadInterests()
	interestMappings := LoadInterestMappings()

	var newUsers []User
	for _, user := range users {
		for _, interestMapping := range interestMappings{
			if user.ID == interestMapping.UserID {
				for _, interest := range interests{
					if interestMapping.InterestID == interest.ID{
						user.Interests = append(user.Interests, interest)
					}
				}
			}
		}
		newUsers = append(newUsers,user)
	}
	viewModel := UserViewModel{Page:page, Users:newUsers}
	data, err := json.Marshal(viewModel)
	CheckError(err)
	wr.Write([]byte(data))
}