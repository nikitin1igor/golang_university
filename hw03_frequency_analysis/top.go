package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

type KeyValue struct {
	Key   string
	Value int
}

var (
	RegexToSplit   = regexp.MustCompile(`[0-9\n\t.,!?;: «»()—\"']+`)
	OnlySeparators = regexp.MustCompile(`^[0-9\n\t.,!?;:\- «»()—\"']+$`)
)

func Top10(inputtedString string) []string {
	wordsMap := make(map[string]int)
	result := make([]string, 0, 10)

	keyValueSlice := make([]KeyValue, 0, len(wordsMap))
	if inputtedString == "" {
		return result
	}

	split := RegexToSplit.Split(inputtedString, -1)

	for i := range split {
		if !OnlySeparators.MatchString(split[i]) {
			loweredString := strings.ToLower(split[i])
			wordsMap[loweredString]++
		}
	}

	for k, v := range wordsMap {
		keyValueSlice = append(keyValueSlice, KeyValue{k, v})
	}

	sort.Slice(keyValueSlice, func(i, j int) bool {
		if keyValueSlice[i].Value > keyValueSlice[j].Value {
			return true
		} else if keyValueSlice[i].Value < keyValueSlice[j].Value {
			return false
		}
		switch strings.Compare(keyValueSlice[i].Key, keyValueSlice[j].Key) {
		case -1:
			return true
		case 1:
			return false
		}
		return false
	})

	for i := 0; i < 10; i++ {
		result = append(result, keyValueSlice[i].Key)
	}

	return result
}
