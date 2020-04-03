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

}
func RunCryptoServerForSymbol(input string){
	var symbolResp,tickerResp,currencyResp map[string]interface{}
	//var symbolRespAll,tickerRespAll,currencyRespAll []map[string]interface{}
	var finalResponse string
	if(input!=""){
		symbolResp= services.GetDetailsForSymbols(constants.Symbol,input)
		tickerResp= services.GetDetailsForSymbols(constants.Ticker,input)
		currencyResp= services.GetDetailsForSymbols(constants.Currency,input)
		finalResponse=getParsedCryptoDetails(symbolResp,tickerResp,currencyResp)	
	}else{
		// symbolRespAll= services.GetDetailsForSymbolAll(constants.Symbol,input)
		// tickerRespAll= services.GetDetailsForSymbolAll(constants.Ticker,input)
		// currencyRespAll= services.GetDetailsForSymbolAll(constants.Currency,input)
		// finalResponse=getParsedCryptoDetailsForAll(symbolRespAll,tickerRespAll,currencyRespAll)
	}
	
	fmt.Println(finalResponse)
}

// func getParsedCryptoDetailsForAll(symbolRespAll,tickerRespAll,currencyRespAll []map[string]interface{}) string{
// 	return "a"
// }

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

	resp:=Result{fmt.Sprintf("%v",id),fmt.Sprintf("%v",fN),fmt.Sprintf("%v",ask),fmt.Sprintf("%v",bid),fmt.Sprintf("%v",last),fmt.Sprintf("%v",open),fmt.Sprintf("%v",low),fmt.Sprintf("%v",high),fmt.Sprintf("%v",fc)}
	
	jsonResponse,err:=json.Marshal(resp)
	if(err!=nil){
		panic(err)
	}
	return string(jsonResponse)

}