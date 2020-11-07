package strs

func Strs2Interface(s []string) []interface{} {
	var interfaces = make([]interface{}, len(s))
	for i, v := range s {
		interfaces[i] = v
	}
	return interfaces
}

func Interface2Strs(interfaces []interface{}) []string {
	var strs = make([]interface{}, len(interfaces))
	for i, v := range interfaces {
		strs[i] = v
	}
	return strs
}
