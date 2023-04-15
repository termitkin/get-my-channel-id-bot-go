package utils

import "strconv"

func FloatToStr(f float64) string {
	int64ChatId := int64(f)
	s := strconv.FormatInt(int64ChatId, 10)

	return s
}
