package repo

type Credential struct {
	Id   string
	Code string
}

func NewCredential(id string, code string) *Credential {
	return &Credential{Id: id, Code: code}
}
