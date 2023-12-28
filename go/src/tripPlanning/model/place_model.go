package model

type Text struct{
	Text string `json:"text"`

}
type Review struct{
	Id string `json:"name"`
	PublishTime string `json:"publishTime"`
	Rating int `json:"rating"`
	Text Text `json:"text"`
}

type Photo struct{
	Id string `json:"name"`
}

type Place struct{
	Id string `json:"id"`
	// Address string `json:"formattedAddress"`
	// DisplayName Text `json:"DisplayName"`
	// Reviews []Review `json:"reviews"`
	// Photos []Photo `json:"photos"`
}