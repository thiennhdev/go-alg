package main

import "fmt"

func getMinimumWindow(original string, check string) string {
	result := ""

	if len(original) < len(check) {
		return result
	}
	startPtr := 0
	pullCheck := make(map[string]int)
	for i := 0; i < len(check); i++ {
		pullCheck[string(check[i])] = pullCheck[string(check[i])] + 1
	}

	countCheck := 0
	pullOriginal := make(map[string]int)
	for i := 0; i < len(original); i++ {
		oString := string(original[i])
		if _, exist := pullCheck[oString]; exist && pullOriginal[oString] < pullCheck[oString] {
			countCheck++
		}
		pullOriginal[oString] = pullOriginal[oString] + 1

		for countCheck >= len(check) && startPtr <= i {
			oStringTemp := string(original[startPtr])
			if _, exist := pullCheck[oStringTemp]; exist {
				if pullOriginal[oStringTemp] <= pullCheck[oStringTemp] {
					countCheck--
				}
				if len(original[startPtr:i+1]) <= len(result) || result == "" {
					result = original[startPtr : i+1]
				}
			}
			pullOriginal[oStringTemp] = pullOriginal[oStringTemp] - 1
			startPtr++
		}
	}

	return result
}

func main() {
	fmt.Println(getMinimumWindow("cdbaebaecd", "abc"))
}
