package data_structures


type Transaction struct {
	Date string `yaml:"date"`
	Shop string `yaml:"shop"`
	Account string `yaml:"account"`
	Method string `yaml:"method"`
	Sum string `yaml:"sum"`
	Items []Items `yaml:"items"`
}

type Items struct {
	Name string `yaml:"name"`
	Amount string `yaml:"amount"`
	Price string `yaml:"price"`
}