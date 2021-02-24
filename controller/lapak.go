package controller

import (
	"go-rest-api-starter/model"
)

//GetListLapakDipasar method to get List Lapak By Pasar
func (c *Controller) GetListLapakDipasar(storeID, limit int) (listLapak model.ResponseListLapak, err error) {
	var (
		_listLapak	[]model.ListLapak
	)

	rows, err := c.DB.Table("lapak").Select("lapak.id, lapak.nama_lapak, lapak.deskripsi, pasar.alamat, kategori.deskripsi as kategori, COALESCE(lapak.image, '')").Joins("JOIN pasar on pasar.id = lapak.idpasar").Joins("JOIN link_kategori_lapak on link_kategori_lapak.idlapak = lapak.id").Joins("JOIN kategori on kategori.id = link_kategori_lapak.idkategori").Where("lapak.idpasar = ?", storeID).Group("lapak.id").Limit(limit).Rows()

	defer rows.Close()

	if err != nil {
		return
	}
	for rows.Next() {
		var getLapak model.ListLapak

		err = rows.Scan(&getLapak.LapakID, &getLapak.NamaLapak, &getLapak.Deskripsi, &getLapak.Alamat, &getLapak.Kategori, &getLapak.Image)
	}

	listLapak.Lapak = _listLapak
	return
}
