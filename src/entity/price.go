package entity

type Price struct {
	Amount float32 `bson:"amount"`
	Unit   string  `bson:"unit"`
}
