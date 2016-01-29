package types

type Foo struct {
	ID    int64  `json:"id",sql:"AUTO_INCREMENT"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
