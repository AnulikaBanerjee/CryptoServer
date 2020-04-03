package services

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

func GetDetailsForSymbolAll(url,symbol string) []map[string]interface{}{
	resp3,err:=http.Get(url+symbol);
	if(err!=nil){
		fmt.Println(err)
	}
	defer resp3.Body.Close()
	body3,err:=ioutil.ReadAll(resp3.Body)
	byt3:=[]byte(body3)
	
	var dat3 []map[string]interface{}
	if err:= json.Unmarshal(byt3,&dat3);err!=nil{
		panic(err)
	}

	return dat3
		
}

func GetDetailsForSymbols(url,symbol string) map[string]interface{}{
	resp3,err:=http.Get(url+symbol);
	if(err!=nil){
		fmt.Println(err)
	}
	defer resp3.Body.Close()
	body3,err:=ioutil.ReadAll(resp3.Body)
	byt3:=[]byte(body3)
	
	
	var dat3  map[string]interface{}
	if err:= json.Unmarshal(byt3,&dat3);err!=nil{
		panic(err)
	}

	return dat3
		
}