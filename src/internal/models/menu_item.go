package models

type MenuItem struct {
	ID            int    `db:"id" json:"id"`
	Title         string `db:"title" json:"title"`
	Description   string `db:"description" json:"description"`
	Weigth        int    `db:"weight" json:"weight"`
	WeigthMeasure string `db:"weight_measure" json:"weightMeasure"`
}
