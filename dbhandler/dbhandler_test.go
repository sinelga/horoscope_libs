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

	result := CheckIfExist(*session, "test.com", "link3")

	if result {
		//		t.Error("Cant be TRUE")

		CheckIfLinksExist(*session, "test.com", "link3")

	} else {

		InsertNewSite(*session, "test.com", "link3")
	}
	
	ZodiacContents(*session, "test.com")
	

}
