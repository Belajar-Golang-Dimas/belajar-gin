package helper

import "github.com/organizations/Belajar-Golang-Dimas/repositories/new/domain/domain"

func MappingToAddress(data map[string]string) domain.Address {
	return domain.Address{
		Street:   data["street"],
		City:     data["city"],
		Province: data["province"],
		Country:  data["country"],
	}
}
