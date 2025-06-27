package models

type Barang struct {
    ID     int     `json:"id"`
    Nama   string  `json:"nama"`
    Harga  float64 `json:"harga"`
    Stok   int     `json:"stok"`
}
