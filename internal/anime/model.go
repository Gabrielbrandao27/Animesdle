package anime

type Naruto struct {
	ID               int64  `json:"id" db:"id"`
	ImgRef           string `json:"img_ref" db:"img_ref"`
	Name             string `json:"name" db:"name"`
	Species          string `json:"species" db:"species"`
	PlaceOrigin      string `json:"place_origin" db:"place_origin"`
	IntroArc         string `json:"intro_arc" db:"intro_arc"`
	Affiliation      string `json:"affiliation" db:"affiliation"`
	ChakraType       string `json:"chakra_type" db:"chakra_type"`
	KekkeiGenkai     string `json:"kekkei_genkai" db:"kekkei_genkai"`
	JutsuAffinity    string `json:"jutsu_affinity" db:"jutsu_affinity"`
	SpecialAttribute string `json:"special_attribute" db:"special_attribute"`
}

type OnePiece struct {
	ID          int64  `json:"id" db:"id"`
	ImgRef      string `json:"img_ref" db:"img_ref"`
	Name        string `json:"name" db:"name"`
	Species     string `json:"species" db:"species"`
	PlaceOrigin string `json:"place_origin" db:"place_origin"`
	IntroArc    string `json:"intro_arc" db:"intro_arc"`
	Affiliation string `json:"affiliation" db:"affiliation"`
	Bounty      int    `json:"bounty" db:"bounty"`
	Haki        string `json:"haki" db:"haki"`
	DevilFruit  string `json:"devil_fruit" db:"devil_fruit"`
	Height      int    `json:"height" db:"height"`
}
