package main

func CheckBilangan(bils ...int) []string {
	result := []string{}
	for _, bilangan := range bils {
		if bilangan%2 == 0 {
			result = append(result, "Genap")
		} else {
			result = append(result, "Ganjil")
		}
	}
	return result
}
