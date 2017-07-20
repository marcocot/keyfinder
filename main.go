package main

import (
	"fmt"
	"os"

	"sort"

	dbf "github.com/LindsayBradford/go-dbf/godbf"
)

type keyValue struct {
	Key   string
	Value int
}

func help() {
	message := fmt.Sprintf("Usage: %s [sbscripts path]\n", os.Args[0])
	os.Stderr.WriteString(message)
}

func findKeys(path string) {
	db, err := dbf.NewFromFile(path, "UTF-8")
	keyCount := make(map[string]int)

	if err != nil {
		panic(err)
	}

	for i := 0; i < db.NumberOfRecords(); i++ {
		code := db.FieldValue(i, 1)

		if val, check := keyCount[code]; check {
			keyCount[code] = val + 1
		} else {
			keyCount[code] = 1
		}
	}

	kv := make([]keyValue, len(keyCount))

	i := 0
	for k, v := range keyCount {
		kv[i] = keyValue{Key: k, Value: v}
		i++
	}

	sort.Slice(kv, func(i, j int) bool {
		return kv[i].Value > kv[j].Value
	})

	for _, k := range kv {
		fmt.Printf("%s: %d\n", k.Key, k.Value)
	}

	fmt.Println("File processed")
}

func main() {
	if len(os.Args) != 2 {
		help()
		return
	}

	findKeys(os.Args[1])
}
