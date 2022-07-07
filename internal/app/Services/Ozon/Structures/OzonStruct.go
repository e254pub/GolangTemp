package Structures

type OzonAuth struct {
	ApiKey   string
	ClientId string
}

type OzonMethod struct {
	Method string
	Url    string
}

type Filter struct {
	Visibility string   `json:"visibility"`
	ProductId  []string `json:"product_id"`
	OfferId    []string `json:"offer_id"`
}

type AttributeStruct struct {
	//поля бд
}

type ProductStruct struct {
}

type UniqueAttribute struct {
}
