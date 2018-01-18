package iplocation

import (
	"github.com/zmap/go-iptree/iptree"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

const CITY_INDEX = 10
const COUNTRY_INDEX = 5
const REGION_INDEX = 7

var IPtree *iptree.IPTree
var cityMap map[string]Location
var geoCoordinates map[string]LocalGeo

func StripChars(sourceString string,patternString string)string{

	return strings.Replace(sourceString,patternString,"",-1)

}

/* Parses TSV,CSV file */
func ParseColumnarFile(fileName string,columnSeperator string,rowSeperator string,function func([]string)){
	b, err := ioutil.ReadFile(fileName) // just pass the file name
	if err != nil {
		return
	}
	fileData		:= string(b)
	records			:= strings.Split(fileData, rowSeperator)
	totalRecords 	:= len(records)

	for i := 0; i < (totalRecords-1) ; i ++ {
		tuple := strings.Split(StripChars(records[i],"\""),columnSeperator)
		function(tuple)
	}

}

func ProcessIPTuple(tuple []string) {

	IPtree.AddByString(tuple[0], tuple[1])
	if len(tuple[7]) <= 0 && len(tuple[8]) <= 0 {
		return
	}

	latitude, latError := strconv.ParseFloat(tuple[7], 64)
	longitude, lonError := strconv.ParseFloat(tuple[8], 64)

	if latError != nil || lonError != nil {
		return
	}

	geoCoordinates[tuple[1]] = LocalGeo{Latitude: latitude, Longitude: longitude}
}

func ProcessCityTuple(tuple []string) {

	locationModel := Location{}
	locationModel.city = tuple[CITY_INDEX]
	locationModel.country = tuple[COUNTRY_INDEX]
	locationModel.state = tuple[REGION_INDEX]
	cityMap[tuple[0]] = locationModel

}

func InitIPTree(){
	IPtree = iptree.New()

	cityMap = make(map[string]Location)
	geoCoordinates = make(map[string]LocalGeo)

	fmt.Println("Warming Up IP Block Cache...")
	ParseColumnarFile("../files/GeoIP2-City-Blocks-IPv4.csv", ",", "\n", ProcessIPTuple)

	fmt.Println("Warming up City Cache...")
	//	geoindex.InitSearch()
	ParseColumnarFile("../files/GeoIP2-City-Locations-en.csv", ",", "\n", ProcessCityTuple)
}

func GetLocationByIpTree(key string) (*Location){
	if geoCode, found, err := IPtree.GetByString(key); err == nil && found {
		geoCodeString, _ := geoCode.(string)
		ilocation := cityMap[geoCodeString]
		return &ilocation
	}
	result := new(Location)
	return result
}