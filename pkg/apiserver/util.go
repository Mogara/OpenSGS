package apiserver

import (
	"reflect"
	"runtime"
)

func nameOfFunction(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func allowOriginFunc(allowedOrigins []string) func(origin string) bool {
	return func(origin string) bool {
		for _, allowed := range allowedOrigins {
			if wildcardMatch([]rune(allowed), []rune(origin)) {
				return true
			}
		}
		return false
	}
}

func wildcardMatch(pattern, str []rune) bool {
	for len(pattern) > 0 {
		switch pattern[0] {
		case '*':
			// not match '-'(ascii: 45) with *
			return wildcardMatch(pattern[1:], str) || (len(str) > 0 && wildcardMatch(pattern, str[1:]) && str[0] != 45)
		default:
			if len(str) == 0 || str[0] != pattern[0] {
				return false
			}
		}
		str = str[1:]
		pattern = pattern[1:]
	}

	return len(str) == 0 && len(pattern) == 0
}
