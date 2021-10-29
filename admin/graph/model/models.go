package model

type Tx struct {
	Hash string `json:"hash"`
	Raw  string `json:"raw"`
}

type TxOutput struct {
	Hash   string `json:"hash"`
	Index  uint32 `json:"index"`
	Amount int64  `json:"amount"`
	Script string `json:"script"`
}

type TxInput struct {
	Hash      string `json:"hash"`
	Index     uint32 `json:"index"`
	PrevHash  string `json:"prev_hash"`
	PrevIndex uint32 `json:"prev_index"`
}
