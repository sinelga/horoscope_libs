package dbhandler

import (
	"fmt"
	"github.com/sinelga/horoscope_libs/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func CheckIfExist(session mgo.Session, site string, link string) bool {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	count, err := c.Find(bson.M{"Site": site}).Limit(1).Count()

	if err != nil {

		log.Fatal(err)
	}
	if count == 0 {
		fmt.Println("not exists")
		return false
	}

	return true
}

func InsertNewSite(session mgo.Session, site string, link string) {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	fortuneresors := &domains.Fortuneresors{}

	fortuneresors.Site.Site = "test.com"

	var now = time.Now()
	var arrLinkinfo []domains.Linkinfo

	linkinfo := domains.Linkinfo{
		Created_at: now,
		Type:       "daily_horoscope",
		Link:       link,
	}
	arrLinkinfo = append(arrLinkinfo, linkinfo)

	fortuneresors.Links = arrLinkinfo

	err := c.Insert(fortuneresors)

	if err != nil {

		log.Fatal(err)
	}

	//	item :=c.Find(bson.M{"Site": site}).Limit(1)

}
