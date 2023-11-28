package data_structures

type Transaction struct {
	Date    string  `yaml:"date"`
	Shop    string  `yaml:"shop"`
	Account string  `yaml:"account"`
	Method  string  `yaml:"method"`
	Currency string `yaml:"currency"`
	Sum     float32 `yaml:"sum"`
	Items   []Items `yaml:"items"`
}

type Items struct {
	Name   string  `yaml:"name"`
	Amount float32 `yaml:"amount"`
	Price  float32 `yaml:"price"`
}

func (trans Transaction) Validate() (bool, float32, float32) {

	var balance float32

	for _, item := range trans.Items {
		balance += item.Amount * item.Price
	}

	if balance == trans.Sum {
		return true, 0, 0
	} else {
		return false, trans.Sum, balance
	}

}

func (trans Transaction) GetAccount() (string) {
	return trans.Account
}

func (trans Transaction) GetCurrency() (string) {
	return trans.Currency
}

func (trans Transaction) GetDate() (string) {
	return trans.Date
} 

func (trans Transaction) GetPartnerName() (string) {
	return trans.Shop
}

func (trans Transaction) GetSum() (float32) {
	return trans.Sum
}

func (trans Transaction) GetMethod() (string) {
	return trans.Method
}

