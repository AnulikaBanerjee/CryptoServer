package utils

import (
	services "CryptoServer/src/services"
	"fmt"
	"encoding/json"
	constants "CryptoServer/src/constants"
)

//response struct for /currency/{symbol}
type Result struct{
	Id string
	FullName string
	Ask string
	Bid string
	Last string
	Open string
	Low string
	High string
	FeeCurrency string
	id string
	symbol string
	baseCurrency string
}

type ResultAll struct{
	Currencies []Result
}

func RunCryptoServerForSymbol(input string) string{
	var symbolResp,tickerResp,currencyResp map[string]interface{}
	var symbolRespAll,tickerRespAll,currencyRespAll []map[string]interface{}
	var finalResponse string
	if(input!=""){
		symbolResp= services.GetDetailsForSymbols(constants.Symbol,input)
		tickerResp= services.GetDetailsForSymbols(constants.Ticker,input)
		temp,_:=symbolResp["baseCurrency"]
		input=fmt.Sprintf("%v",temp)
		currencyResp= services.GetDetailsForSymbols(constants.Currency,input)
		finalResponse=getParsedCryptoDetails(symbolResp,tickerResp,currencyResp)	
	}else{
		 symbolRespAll= services.GetDetailsForSymbolAll(constants.Symbol)
		 tickerRespAll= services.GetDetailsForSymbolAll(constants.Ticker)
		 currencyRespAll= services.GetDetailsForSymbolAll(constants.Currency)
		 finalResponse=getParsedCryptoDetailsForAll(symbolRespAll,tickerRespAll,currencyRespAll)
	}
	
	return finalResponse
}

func getParsedCryptoDetailsForAll(symbolRespAll,tickerRespAll,currencyRespAll []map[string]interface{}) string{
	var finalResp string
	resultArr := make([]Result,0)
	tempResp:=make(map[string]Result)
	for _,symbol := range symbolRespAll{
		id,_:=symbol["id"]
		fc,_:=symbol["feeCurrency"]
		bc,_:=symbol["baseCurrency"]
		tempResp[fmt.Sprintf("%v",id)]=Result{id:fmt.Sprintf("%v",id),FeeCurrency:fmt.Sprintf("%v",fc),baseCurrency:fmt.Sprintf("%v",bc),}
	}
	for _,tickerResp := range tickerRespAll{
		id,_:=tickerResp["symbol"]
		ask,_:=tickerResp["ask"]
		bid,_:=tickerResp["bid"]
		last,_:=tickerResp["last"]
		open,_:=tickerResp["open"]
		low,_:=tickerResp["low"]
		high,_:=tickerResp["high"]
		res,ok:=tempResp[fmt.Sprintf("%v",id)]
		if ok{
			res.Ask=fmt.Sprintf("%v",ask)
			res.Bid=fmt.Sprintf("%v",bid)
			res.Last=fmt.Sprintf("%v",last)
			res.Open=fmt.Sprintf("%v",open)
			res.Low=fmt.Sprintf("%v",low)
			res.High= fmt.Sprintf("%v",high)
			tempResp[fmt.Sprintf("%v",id)]=res		
		}
	}
	currencyRespMap := make(map[string]string)
	for _,currencyResp:= range currencyRespAll{
		id,_:=currencyResp["id"]
		fN,_:=currencyResp["fullName"]
		currencyRespMap[fmt.Sprintf("%v",id)]=fmt.Sprintf("%v",fN)
	}
	for id,result1 := range tempResp{
		result1.Id=fmt.Sprintf("%v",result1.baseCurrency)
		result1.FullName,_=currencyRespMap[result1.Id]
		tempResp[id]=result1
	}
	for _,result1 := range tempResp{
		resultArr=append(resultArr,result1)
	}
	var jhg ResultAll
	jhg.Currencies=resultArr
	hg,_:=json.Marshal(jhg)
	finalResp=string(hg)

	return finalResp
}

func getParsedCryptoDetails(symbolResp,tickerResp,currencyResp map[string]interface{}) string{
	id,_:=currencyResp["id"]
	fN,_:=currencyResp["fullName"]
	ask,_:=tickerResp["ask"]
	bid,_:=tickerResp["bid"]
	last,_:=tickerResp["last"]
	open,_:=tickerResp["open"]
	low,_:=tickerResp["low"]
	high,_:=tickerResp["high"]
	fc,_:=symbolResp["feeCurrency"]

	resp:=Result{fmt.Sprintf("%v",id),fmt.Sprintf("%v",fN),fmt.Sprintf("%v",ask),fmt.Sprintf("%v",bid),fmt.Sprintf("%v",last),fmt.Sprintf("%v",open),fmt.Sprintf("%v",low),fmt.Sprintf("%v",high),fmt.Sprintf("%v",fc),fmt.Sprintf("%v",id),fmt.Sprintf("%v",id),fmt.Sprintf("%v",id)}
	
	jsonResponse,err:=json.Marshal(resp)
	if err!=nil{
		panic(err)
	}
	return string(jsonResponse)

}