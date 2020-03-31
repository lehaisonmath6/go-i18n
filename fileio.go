package i18n

import "os"

func parsePath(path string) (lang, format string) {
	formatStartIdx := -1
	for i := len(path) - 1; i >= 0; i-- {
		c := path[i]
		if os.IsPathSeparator(c) {
			if formatStartIdx != -1 {
				lang = path[i+1 : formatStartIdx]
			}
			return
		}
		if path[i] == '.' {
			if formatStartIdx != -1 {
				lang = path[i+1 : formatStartIdx]
				return
			}
			if formatStartIdx == -1 {
				format = path[i+1:]
				formatStartIdx = i
			}
		}
	}
	if formatStartIdx != -1 {
		lang = path[:formatStartIdx]
	}
	return
}
