package maps

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
		"Azure":    "https://www.azure.com",
	}
	mapWithNumbersAsKeys := map[int]string{
		1: "One",
		2: "Two",
	}

	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(mapWithNumbersAsKeys)
	fmt.Println(mapWithNumbersAsKeys[1])

	websites["LinkedIn"] = "https://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Facebook")
	fmt.Println(websites)
}
