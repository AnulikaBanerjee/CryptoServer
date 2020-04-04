package main

import (
	"fmt"
	
	"net/http"
	"strings"
	utils "CryptoServer/src/utils"
	constants "CryptoServer/src/constants"
)


func currencyHandlerFunc(w http.ResponseWriter, r *http.Request){
	/* if getting query params
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items, 
	// we only want the single item.
	key := keys[0]

	fmt.Println("Url Param 'key' is: " + string(key))
	*/
	//for path
	path:=r.URL.Path
	//fmt.Println("Path=: "+path)
	input:=strings.TrimPrefix(path,"/currency/")
	if input=="all" {
		input=""
	}
	fmt.Fprintf(w, utils.RunCryptoServerForSymbol(input))

}





func main() {

	http.HandleFunc("/currency/", currencyHandlerFunc)
	http.ListenAndServe(":"+constants.ListenPort,nil)
	
}

