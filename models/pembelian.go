package models

type pembelian struct {
	ID           int    `json:"id" bson:"_id"`                    // ID pembelian
	Kode_Barang  string `json:"kode_barang" bson:"kode_barang"`   // Kode barang yang dibeli
	Nama_Barang  string `json:"nama_barang" bson:"nama_barang"`   // Nama barang
	Harga_Beli   int    `json:"harga_beli" bson:"harga_beli"`     // H
	Jumlah       int    `json:"jumlah_beli" bson:"jumlah_beli"`   // Jumlah pemb
	Tanggal_Beli string `json:"tanggal_beli" bson:"tanggal_beli"` //
	supplier     string `json:"supplier" bson:"supplier"`         // Nama supplier

}
