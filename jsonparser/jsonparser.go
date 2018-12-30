package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	data := []byte(`
		{
    		"dev": {
        		"host":"172.24.112.111",
        		"port": "27017",
        		"username": "qars1db3adm1",
        		"password": "wUch3BaV",
        		"param": "authSource=qars1db3"
    		},
    		"local": {
        		"host": "localhost",
        		"port": "27017"
    		}
		}
	`)
	res := map[string]interface{}{}

	json.Unmarshal(data, &res)
	parsejson(res)

}

func parsejson(j map[string]interface{}) {
	for k, v := range j {
		switch v.(type) {
		case map[string]interface{}:
			fmt.Printf("key : %s value: %s\n", k, v)
			parsejson(v.(map[string]interface{}))
		case string:
			fmt.Printf("key : %s value: %s\n", k, v)
		case int:
			fmt.Println("Int")
		case interface{}:
			fmt.Println("Interface")
		default:
			fmt.Println("Default")
		}
	}
}
