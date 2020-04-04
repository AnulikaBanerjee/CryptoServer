package services

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func GetDetailsForSymbolAll(url string) []map[string]interface{}{
	response,err:=http.Get(url);
	if(err!=nil){
		fmt.Println(err)
	}
	
	defer response.Body.Close()
	responseBody,err:=ioutil.ReadAll(response.Body)

	byteData:=[]byte(responseBody)
	
	var data []map[string]interface{}
	if err:= json.Unmarshal(byteData, &data);err!=nil{
		panic(err)
	}

	return data
		
}

func GetDetailsForSymbols(url,symbol string) map[string]interface{}{
	response,err:=http.Get(url+symbol);
	if(err!=nil){
		fmt.Println(err)
	}
	defer response.Body.Close()
	responseBody,err:=ioutil.ReadAll(response.Body)
	byteData:=[]byte(responseBody)
	
	
	var data  map[string]interface{}
	if err:= json.Unmarshal(byteData, &data);err!=nil{
		panic(err)
	}

	return data
		
}