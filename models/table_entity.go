package models

var TableEntity = map[string][]string{
	"countries": {"guid", "title", "code", "continent"},
	"city":      {"guid", "title", "country_id", "city_code", "latitude", "longitude", `"offset"`, "timezone_id", "country_name"},
}
