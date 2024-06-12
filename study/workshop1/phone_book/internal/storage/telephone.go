package storage

type telephoneRecord struct {
	Name 	  string 	`json: "name"`
	Telephone string 	`json: "telephone"`
	CreatedAt time.Time `json: "created_at"`
}

func transform(telephone models.Telephone) telephoneRecord {
	return telephoneRecord{
		Name: string(telephone.Name),
		Telephone
	}
}