package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main()  {
	var urlData = "https://raw.githubusercontent.com/brenik/GO-Unmarshal-JSON-with-some-unknown-field/master/data.json"

	resp, _ := http.Get(urlData)
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var decoded []map[string]interface{}
	err = json.Unmarshal(data, &decoded)
	if err != nil {
		panic(err)
	}

	for i, val := range decoded{
		if value, inMap := val["global_id"]; inMap {
			fmt.Printf("global_id № %+v  = %.f\n",i,value)
		}
	}

}