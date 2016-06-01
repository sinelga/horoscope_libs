package dbhandler

import (
	"fmt"
	"github.com/sinelga/horoscope_libs/domains"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func UpdateContents(session mgo.Session, site string, link string, zodiacs []domains.Zodiac) {
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	var sitetocheck domains.Fortuneresors
	sitebson := bson.M{"site.site": site}
	c.Find(sitebson).One(&sitetocheck)

	for i, linkinfo := range sitetocheck.Links {

		if linkinfo.Link == link {

			sitetocheck.Links[i].Zodiacs = zodiacs
			c.Update(sitebson, sitetocheck)			

		}

	}

}

func ZodiacContents(session mgo.Session, site string) []string {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	var sitetocheck domains.Fortuneresors
	sitebson := bson.M{"site.site": site}

	c.Find(sitebson).One(&sitetocheck)

	//	fmt.Println(sitetocheck)

	var tofeedzodiaclinks []string

	for _, linkinfo := range sitetocheck.Links {

		if len(linkinfo.Zodiacs) == 0 {

			fmt.Println(linkinfo.Link)
			tofeedzodiaclinks = append(tofeedzodiaclinks, linkinfo.Link)

		}

	}

	return tofeedzodiaclinks

}

func CheckIfExist(session mgo.Session, site string, link string) bool {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	count, err := c.Find(bson.M{"site.site": site}).Limit(1).Count()

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

	fortuneresors.Site.Site = site

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

}

func CheckIfLinksExist(session mgo.Session, site string, link string) {

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("horoscope").C("arch")

	var sitetocheck domains.Fortuneresors
	var now = time.Now()

	sitebson := bson.M{"site.site": site}
	c.Find(sitebson).One(&sitetocheck)

	set := make(map[string]struct{})

	for _, links := range sitetocheck.Site.Links {

		fmt.Println("links", links.Link)
		set[links.Link] = struct{}{}

	}

	if _, ok := set[link]; ok {

		fmt.Println("element found")
	} else {
		fmt.Println("element not found")

		linkinfo := domains.Linkinfo{
			Created_at: now,
			Type:       "daily_horoscope",
			Link:       link,
		}

		modsite := sitetocheck

		modsite.Links = append(modsite.Links, linkinfo)

		c.Update(sitebson, modsite)

	}

}
