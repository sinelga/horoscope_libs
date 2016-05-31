package dbhandler

import (
	"gopkg.in/mgo.v2"
//	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestCheckIfExist(t *testing.T) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	
	result :=CheckIfExist(*session,"test.com","link")
	
	if result {
		t.Error("Cant be TRUE")
		
	}
	
	InsertNewSite(*session, "test.com","link2") 
	

}
