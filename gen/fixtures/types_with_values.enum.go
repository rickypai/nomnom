package fixtures

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidCity = errors.New("invalid City")
)

func IsCity(in string) bool {
	switch in {
	case "london":
		return true
	case "oakland":
		return true
	case "portland":
		return true
	case "seattle":
		return true
	case "San Francisco":
		return true
	case `"`:
		return true
	case "sekret":
		return true
	}

	return false
}

func ToCity(in string) (City, bool) {
	switch in {
	case "london":
		return CityLondon, true
	case "oakland":
		return CityOakland, true
	case "portland":
		return CityPortland, true
	case "seattle":
		return CitySeattle, true
	case "San Francisco":
		return CitySanFrancisco, true
	case `"`:
		return CityQuotes, true
	case "sekret":
		return citySekret, true
	}

	return City(""), false
}

func ToCityErr(in string) (City, error) {
	if city, ok := ToCity(in); ok {
		return city, nil
	}

	return City(""), fmt.Errorf("casting `%v`: %w", in, ErrInvalidCity)
}

func MustToCity(in string) City {
	city, err := ToCityErr(in)
	if err != nil {
		panic(err)
	}

	return city
}

type CityValues struct{}

func (_ *CityValues) Values() []string {
	return []string{
		"london",
		"oakland",
		"portland",
		"seattle",
		"San Francisco",
		`"`,
		"sekret",
	}
}
