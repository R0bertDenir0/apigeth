package models

//type Block
type Block struct {
	BlockNumber       int64         `json:"block_number"`
	TimeStamp         uint64        `json:"time_stamp"`
	Difficulty        uint64        `json:"difficulty"`
	Hash              string        `json:"hash"`
	TransactionsCount int           `json:"transactions_count"`
	Transactions      []Transaction `json:"transactions"`
}

//type Transaction
type Transaction struct {
	Hash     string `json:"hash"`
	Value    string `json:"value"`
	Gas      uint64 `json:"gas"`
	GasPrice uint64 `json:"gas_price"`
	Nonce    uint64 `json:"nonce"`
	To       string `json:"to"`
	Pending  bool   `json:"pending"`
}

//type Transfer Eth
type TransferEth struct {
	PrivKey string `json:"priv_key"`
	To      string `json:"to"`
	Amount  int64  `json:"amount"`
}

//type hash response
type HashResponse struct {
	Hash string `json:"hash_response"`
}

//type BalanceResponse
type BalanceResponse struct {
	Adress  string `json:"adress"`
	Balance string `json:"balance"`
	Symbol  string `json:"symbol"`
	Unit    string `json:"unit"`
}

//Error Data structure
type Error struct {
	Code    uint64 `json:"code"`
	Message string `json:"message"`
}
