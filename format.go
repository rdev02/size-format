package sizeformat

import (
	"errors"
	"fmt"
	"strings"
)

const (
	kilo        = 1024
	printFormat = "%.2f%s"
	scanFormat  = "%f%2s"
)

/* Size constants */
const (
	B  = 1
	KB = kilo * B
	MB = kilo * KB
	GB = kilo * MB
	TB = kilo * GB
	PB = kilo * TB
)

const (
	sPB = "PB"
	sTB = "TB"
	sGB = "GB"
	sKB = "KB"
	sMB = "MB"
	sB  = "B"
)

/*ToString converts size to a string, using the highest size constant */
func ToString(size int64) string {
	sizeF, sizeS := getSizeStr(size)

	return fmt.Sprintf(printFormat, sizeF, sizeS)
}

/*ToNum parses string into num bytes */
func ToNum(formattedString *string) (int64, error) {
	var sizeF float64
	var sizeStr string
	tail := 2

	uppercaseStr := strings.Trim(strings.ToUpper(*formattedString), " ")
	if strings.HasSuffix(uppercaseStr, sPB) {
		sizeStr = sPB
	} else if strings.HasSuffix(uppercaseStr, sTB) {
		sizeStr = sTB
	} else if strings.HasSuffix(uppercaseStr, sGB) {
		sizeStr = sGB
	} else if strings.HasSuffix(uppercaseStr, sMB) {
		sizeStr = sMB
	} else if strings.HasSuffix(uppercaseStr, sKB) {
		sizeStr = sKB
	} else if strings.HasSuffix(uppercaseStr, sB) {
		sizeStr = sB
		tail = 1
	} else {
		return 0, errors.New("unknown unit: " + (*formattedString)[len(*formattedString)-2:])
	}

	uppercaseStr = strings.TrimRight(uppercaseStr[:len(uppercaseStr)-tail], " ")
	_, err := fmt.Sscanf(uppercaseStr, "%f", &sizeF)
	if err != nil || sizeF <= 0 {
		return 0, errors.New("could not parse size value. must be a positive float")
	}

	return getSizeLong(&sizeF, &sizeStr)
}

func getSizeStr(size int64) (float64, string) {
	fSize := float64(size)
	if size >= PB {
		return float64(fSize / PB), sPB
	} else if size >= TB {
		return float64(fSize / TB), sTB
	} else if size >= GB {
		return float64(fSize / GB), sGB
	} else if size >= MB {
		return float64(fSize / MB), sMB
	} else if size >= KB {
		return float64(fSize / KB), sKB
	}

	return fSize, sB
}

func getSizeLong(sizeF *float64, sizeStr *string) (int64, error) {
	if strings.Compare(*sizeStr, sB) == 0 {
		return int64(*sizeF), nil
	} else if strings.Compare(*sizeStr, sKB) == 0 {
		return int64(*sizeF * KB), nil
	} else if strings.Compare(*sizeStr, sMB) == 0 {
		return int64(*sizeF * MB), nil
	} else if strings.Compare(*sizeStr, sGB) == 0 {
		return int64(*sizeF * GB), nil
	} else if strings.Compare(*sizeStr, sTB) == 0 {
		return int64(*sizeF * TB), nil
	} else if strings.Compare(*sizeStr, sPB) == 0 {
		return int64(*sizeF * PB), nil
	}

	return 0, errors.New("unknown size: " + *sizeStr)
}
