package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func prettify(v interface{}) interface{} {
	switch v := v.(type) {
	case []interface{}:
		nv := make([]interface{}, 0)
		for _, e := range v {
			if e = prettify(e); e != nil {
				nv = append(nv, e)
			}
		}
		if len(nv) == 0 {
			return nil
		}
		return nv
	case map[string]interface{}:
		nv := make(map[string]interface{})
		for k, e := range v {
			if e = prettify(e); e != nil {
				nv[k] = e
			}
		}
		if len(nv) == 0 {
			return nil
		}
		return nv
	default:
		return v
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := 0
	fmt.Sscanf(scanner.Text(), "%d", &t)
	results := make([]interface{}, t)
	for i := 0; i < t; i++ {
		scanner.Scan()
		n := 0
		fmt.Sscanf(scanner.Text(), "%d", &n)
		jsonStr := ""
		for j := 0; j < n; j++ {
			scanner.Scan()
			jsonStr += scanner.Text()
		}
		var jsonObj interface{}
		json.Unmarshal([]byte(jsonStr), &jsonObj)
		jsonObj = prettify(jsonObj)
		results[i] = jsonObj
	}
	jsonBytes, _ := json.Marshal(results)
	fmt.Println(string(jsonBytes))
}
