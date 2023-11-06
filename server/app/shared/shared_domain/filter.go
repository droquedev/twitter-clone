package shared_domain

import "fmt"

type Filters struct {
	Field    string
	Operator string
	Value    string
	Logic    string
}

func GenerateFilter(filters []Filters) (string, []interface{}) {
	query := ""
	args := make([]interface{}, len(filters))

	for i, filter := range filters {

		query += fmt.Sprintf("%s %s $%d", filter.Field, filter.Operator, i+1)

		if i != len(filters)-1 {
			query += fmt.Sprintf(" %s ", filter.Logic)
		}

		args[i] = filter.Value
	}

	return query, args
}
