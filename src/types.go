package main

import "github.com/gorilla/mux"

//router service struct
type Server struct {
	router *mux.Router
}

//config struct
type Config struct {
	CRUDHost            string
	CRUDPort            string
	GENESISEXPORTERPort string
}

type ExportAsset struct {
	AssetTypeID string `json:"id"`
}

type ExportAssetResponse struct {
	AssettypeLevelID      string `json:"level"`
	Code                  string `json:"code"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	IsUTC                 bool   `json:"isutc"`
	SizeUnit              string `json:"sizeunit"`
	TypeLookup            string `json:"typelookup"`
	SizeLookup            string `json:"sizelookup"`
	Dimension1Name        string `json:"dimension1name"`
	Dimension1Description string `json:"dimension1description"`
	Dimension1Unit        string `json:"dimension1unit"`
	Dimension2Name        string `json:"dimension1name"`
	Dimension2Description string `json:"dimension2description"`
	ExtentFormula         string `json:"extentformula"`
	DepreciationModel     string `json:"depreciationmodel"`
	DepreciationMethod    string `json:"depreciationmethod"`
	ISActive              bool   `json:"isactive"`
}
