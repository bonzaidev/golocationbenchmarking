package iplocation
import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

var Reader *geoip2.Reader

func InitMaxMindGeoIPReader(){
	db, err := geoip2.Open("../files/GeoIP2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	Reader = db
}


func GetLocationByIp(key string) (*Location){
	ip := net.ParseIP(key)
	record, err := Reader.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	loc := new(Location)
	loc.city = record.City.Names["en"]
	loc.state = record.Subdivisions[0].Names["en"]
	loc.country = record.Country.Names["en"]
	return loc
}