package request

import "gokes/pkg/utils"

func PaginationConfig(conditions map[string][]string) utils.PaginationConfig {
	filterable := map[string]string{
		"end_at": utils.DateType,
	}

	return utils.NewRequestPaginationConfig(conditions, filterable)
}
