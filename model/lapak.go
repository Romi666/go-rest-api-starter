package model

type (
	RequestListLapak struct {
		PasarID   int    `json:"pasar_id"`
		Limit     int    `json:"limit"`
	}

	ListLapak struct {
		LapakID   string `json:"id_lapak" gorm:"column:id"`
		NamaLapak string `json:"nama_lapak" gorm:"column:nama_lapak"`
		Deskripsi string `json:"deskripsi" gorm:"column:deskripsi"`
		Alamat    string `json:"alamat" gorm:"column:alamat"`
		Kategori  string `json:"kategori" gorm:"column:kategori"`
		Image     string `json:"image" gorm:"column:image"`
	}

	ResponseListLapak struct {
		Lapak      []ListLapak `json:"list_lapak"`
	}
)
