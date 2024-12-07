package task_4

func StartTask4(slice1 []string, slice2 []string) []string {
	m := make(map[string]struct{}, len(slice2))
	res := make([]string, 0, len(slice1))
	for _, v := range slice2 {
		m[v] = struct{}{}
	}

	for _, v := range slice1 {
		if _, ok := m[v]; !ok {
			res = append(res, v)
		}
	}
	return res
}
