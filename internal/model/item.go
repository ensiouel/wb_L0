package model

type Item struct {
	ChrtID      int    `db:"chrt_id" json:"chrt_id" validate:"required"`
	TrackNumber string `db:"track_number" json:"track_number" validate:"required"`
	Price       int    `db:"price" json:"price" validate:"gt=0"`
	RID         string `db:"rid" json:"rid" validate:"required"`
	Name        string `db:"name" json:"name" validate:"required"`
	Sale        int    `db:"sale" json:"sale" validate:"gte=0"`
	Size        string `db:"size" json:"size" validate:"required"`
	TotalPrice  int    `db:"total_price" json:"total_price" validate:"gt=0"`
	NmID        int    `db:"nm_id" json:"nm_id" validate:"required"`
	Brand       string `db:"brand" json:"brand" validate:"required"`
	Status      int    `db:"status" json:"status" validate:"required"`
}
