package logging

import "fmt"

func mapToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for k, v := range extra {
		params = append(params, string(k))
		params = append(params, v)
	}
	fmt.Println(params)
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
