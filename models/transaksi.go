package models

type Transaksi struct {
    ID        int      `json:"id"`
    BarangIDs []int    `json:"barang_ids"`
    Total     float64  `json:"total"`
    Waktu     string   `json:"waktu"` 
}
