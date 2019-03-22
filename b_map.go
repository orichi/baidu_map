package baidu_map

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

var conf BDConfig
var apiUrl *url.URL

func init() {

	filename := "configs/bmap.yml"
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		panic(err)
	}

	apiUrl, _ = url.Parse(conf.ApiUrl)

}

func (bMap *BMapItem) CompileLocationInfo() string {
	data, err := json.Marshal(bMap.Location)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
func (bMap *BMapItem) BlankLocation() bool {
	return bMap.Location.Lng == 0 && bMap.Location.Lat == 0
}

func Query(organName string) ([]byte, error) {
	regionName := extractTag(organName)
	return QueryWithRegion(organName, regionName)
}

// 设定查询参数，并查询返回结果
func QueryWithRegion(organName string, region string) ([]byte, error) {
	q := apiUrl.Query()
	q.Set("output", "json")
	q.Set("region", region)
	q.Set("ak", conf.AppKey)
	q.Set("query", organName)
	apiUrl.RawQuery = q.Encode()
	writeLog([]byte(organName + " " + region))
	writeLog([]byte(apiUrl.String()))
	res, err := http.Get(apiUrl.String());
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	result, err := ioutil.ReadAll(res.Body)
	writeLog(result)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func QueryFormatResp(organName string, organRegion string) *BMapResponse {
	repData, err := QueryWithRegion(organName, organRegion)
	if err != nil {
		log.Fatal(err)
	}
	var bmapResponse BMapResponse
	err = json.Unmarshal(repData, &bmapResponse)
	if err != nil {
		log.Fatal(err)
	}
	return &bmapResponse
}

func writeLog(params ...[]byte) {
	f, err := os.OpenFile(conf.QueryLog, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	} else {
		for _, item := range params {
			_, err = f.Write(item)

			if err != nil {
				panic(err)
			}
		}
		f.Write([]byte("\n"))
	}
	defer f.Close()
}
