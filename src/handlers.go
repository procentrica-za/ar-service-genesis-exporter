package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s *Server) handleexportasset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Handle Export Asset Has Been Called...")
		//Get Asset ID from URL
		assettypeid := r.URL.Query().Get("assettypeid")

		//Check if Asset ID provided is null
		if assettypeid == "" {
			w.WriteHeader(500)
			fmt.Fprint(w, "Asset Type ID not properly provided in URL")
			fmt.Println("Asset Type ID not proplery provided in URL")
			return
		}

		//post to crud service
		req, respErr := http.Get("http://" + config.CRUDHost + ":" + config.CRUDPort + "/assetregister?assettypeid=" + assettypeid)

		//check for response error of 500
		if respErr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, respErr.Error())
			fmt.Println("Error in communication with CRUD service endpoint for export asset")
			return
		}
		if req.StatusCode != 200 {
			fmt.Fprint(w, "Request to DB can't be completed...")
			fmt.Println("Unable to asset export")
		}
		if req.StatusCode == 500 {
			w.WriteHeader(500)

			bodyBytes, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatal(err)
			}
			bodyString := string(bodyBytes)
			fmt.Fprintf(w, "Request to DB can't be completed..."+bodyString)
			fmt.Println("Request to DB can't be completed..." + bodyString)
			return
		}

		//close the request
		defer req.Body.Close()

		//create new response struct
		var assetResponse ExportAssetResponse

		//decode request into decoder which converts to the struct
		decoder := json.NewDecoder(req.Body)

		err := decoder.Decode(&assetResponse)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err.Error())
			fmt.Println("Error occured in decoding asset response")
			return
		}
		js, jserr := json.Marshal(assetResponse)
		if jserr != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, jserr.Error())
			fmt.Println("Error occured when trying to marshal the response to export asset")
			return
		}

		//return back to Front-End user
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(js)
	}
}
