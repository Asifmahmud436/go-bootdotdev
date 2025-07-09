package main

func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	for _,v := range names{
		firstChar := rune(v[0])
		if _,ok := result[firstChar]; !ok{
			result[firstChar] = make(map[string]int) 
		}
		result[firstChar][v] +=1
		
	}
	return result
}
