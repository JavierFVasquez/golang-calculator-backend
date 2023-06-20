package utils

func StringIntersection(arr1, arr2 []string) []string {
	inter := []string{}

	for _, i := range arr1 {
		for _, j := range arr2 {
			if i == j {
				inter = append(inter, i)
			}
		}
	}

	return inter
}
