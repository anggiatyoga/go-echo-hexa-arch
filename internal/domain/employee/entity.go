package employee

type Employee struct {
	ID      int    `json:"id"`
	Name    string `json:"nama"`
	Address string `json:"alamat"`
	Phone   string `json:"telepon"`
}
