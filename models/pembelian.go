package models

import "time"

type Pembelian struct {
	ID           int       `json:"id" bson:"_id"`
	Kode_Barang  string    `json:"kode_barang" bson:"kode_barang"`
	Nama_Barang  string    `json:"nama_barang" bson:"nama_barang"`
	Harga_Beli   int       `json:"harga_beli" bson:"harga_beli"`
	Jumlah       int       `json:"jumlah_beli" bson:"jumlah_beli"`
	Tanggal_Beli time.Time `json:"tanggal_beli" bson:"tanggal_beli"`
	Supplier     string    `json:"supplier" bson:"supplier"`
}
