package types

type Entry struct {
	ID    int    `json:"id",sql:"AUTO_INCREMENT"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
