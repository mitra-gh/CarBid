package logging

func mapToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	return params
}

func mapToZeroParams(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := make(map[string]interface{}, 0)
	for k, v := range extra {
		params[string(k)] = v
	}
	return params
}

func prepareLogKey(extra map[ExtraKey]interface{}, category Category, subCategory SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}

	extra["Category"] = category
	extra["SubCategory"] = subCategory

	return mapToZapParams(extra)
}
