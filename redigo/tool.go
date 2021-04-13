package redigo

// 判断字符串是否在切片中
func inSlice(val string, slice []string) bool {
	for _, b := range slice {
		if b == val {
			return true
		}
	}
	return false
}

// 将key加入到v切片头部
func kpv(key string, values []string) []interface{} {
	a := append([]string{key}, values...)

	//converting a []string to a []interface{}
	x := make([]interface{}, len(a))
	for i, v := range a {
		x[i] = v
	}

	return x
}
