package util

type WeatherAPI struct {
	Coord      Weather_Coord
	Weather    Weather_Weather
	Base       Weather_Base
	Main       Weather_Main
	Visibility Weather_Visibility
	Wind       Weather_Wind
	Clouds     Weather_Clouds
	DT         Weather_DT
	SYS        Weather_SYS
	TZ         Weather_TZ
	ID         Weather_ID
	Name       Weather_Name
	COD        Weather_COD
}

type Weather_Coord struct {
	Lon float32 `json:"lon"`
	Lat float32 `json:"lat"`
}

type Weather_Weather struct {
	ID   int    `json:"id"`
	Main string `json:"main"`
	Desc string `json:"description"`
	Icon string `json:"icon"`
}

type Weather_Base struct {
	Base string `json:"base"`
}

type Weather_Main struct {
	Temp float32 `json:"temp"`
	Feel float32 `json:"feels_like"`
	TMin float32 `json:"temp_min"`
	TMax float32 `json:"temp_max"`
	Pres int     `json:"pressure"`
	Humi int     `json:"humidity"`
}

type Weather_Visibility struct {
	Vis int `json:"visibility"`
}

type Weather_Wind struct {
	Spd float32 `json:"speed"`
	Deg int     `json:"deg"`
}

type Weather_Clouds struct {
	All bool `json:"all"`
}

type Weather_DT struct {
	DT int `json:"dt"`
}

type Weather_SYS struct {
	Type int     `json:"type"`
	ID   int     `json:"id"`
	Mess float32 `json:"message"`
	Ctry string  `json:"country"`
	Sunr int     `json:"sunrise"`
	Suns int     `json:"sunset"`
}

type Weather_TZ struct {
	TZ int `json:"timezone"`
}

type Weather_ID struct {
	ID int `json:"id"`
}

type Weather_Name struct {
	Name string `json:"name"`
}

type Weather_COD struct {
	COD int `json:"cod"`
}
