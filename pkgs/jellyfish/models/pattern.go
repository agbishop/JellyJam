package models

import "encoding/json"

type (
	PatterFileListInfo struct {
		Folders  string `json:"folders"`
		Name     string `json:"name"`
		ReadOnly bool   `json:"readOnly"`
	}
	PatterFileList struct {
		PatternFileList []PatterFileListInfo `json:"patternFileList"`
	}
	Zones struct {
		Save  bool            `json:"save"`
		Zones map[string]Zone `json:"zones"`
	}
	Zone struct {
		NumPixels int `json:"numPixels"`
		PortMap   []struct {
			CtlrName        string `json:"ctlrName"`
			PhyEndIdx       int    `json:"phyEndIdx"`
			PhyPort         int    `json:"phyPort"`
			PhyStartIdx     int    `json:"phyStartIdx"`
			ZoneRGBStartIdx int    `json:"zoneRGBStartIdx"`
		} `json:"portMap"`
	}
	PatternData struct {
		Colors              []int  `json:"colors"`
		SpaceBetweenPixels  int    `json:"spaceBetweenPixels"`
		EffectBetweenPixels string `json:"effectBetweenPixels"`
		LedOnPos            struct {
			Field1 []interface{} `json:"0"`
			Field2 []int         `json:"1"`
		} `json:"ledOnPos"`
		Type      string `json:"type"`
		Skip      int    `json:"skip"`
		NumOfLeds int    `json:"numOfLeds"`
		Direction string `json:"direction"`
		RunData   struct {
			Speed       int    `json:"speed"`
			Brightness  int    `json:"brightness"`
			Effect      string `json:"effect"`
			EffectValue int    `json:"effectValue"`
			RgbAdj      []int  `json:"rgbAdj"`
		} `json:"runData"`
	}
	PatternFileData struct {
		Folders     string `json:"folders"`
		RawJsonData string `json:"jsonData"`
		Data        PatternData
		Name        string `json:"name"`
	}
	PatternSettings struct {
		PatternFileData PatternFileData `json:"patternFileData"`
	}
)

func (p *PatternSettings) UnmarshalJSON(data []byte) error {
	type _r PatternSettings
	var res _r
	if err := json.Unmarshal(data, &res); err != nil {
		return err
	}
	if err := json.Unmarshal([]byte(res.PatternFileData.RawJsonData), &res.PatternFileData.Data); err != nil {
		return err
	}
	*p = PatternSettings(res)
	return nil
}
