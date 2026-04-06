package util

// 核心：将 map 转为 []interface{} 后打散传递
func mapToArgs(m map[string]interface{}) []interface{} {
	var args []interface{}
	for k, v := range m {
		args = append(args, k, v)
	}
	return args
}
