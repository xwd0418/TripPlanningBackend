package model

type Text struct {
	Text string `json:"text"`
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Review struct {
	Id          string `json:"name"`
	PublishTime string `json:"publishTime"`
	Rating      int    `json:"rating"`
	Text        Text   `json:"text"`
}

type Photo struct {
	Id string `json:"name"`
}

type Place struct {
	Id          string   `json:"id"`
	DisplayName Text     `json:"DisplayName"`
	PlaceType   Text     `json:"primaryTypeDisplayName"`
	Address     string   `json:"formattedAddress"`
	Location    Location `json:"location"`
	Reviews     []Review `json:"reviews"`
	Photos      []Photo  `json:"photos"`
}
