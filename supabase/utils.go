package supabase

import "strings"

func extractProjectIdFromUrl(url string) string {
	splitStr := strings.Split(url, ".")
	return strings.Split(splitStr[0], "//")[1]
}
