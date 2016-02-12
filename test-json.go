package main

import "fmt"
import "encoding/json"

type PublicKey struct {
	Id  int
	Key string
}

type KeysResponse struct {
	Collection []PublicKey
}

func main() {
	keysBody := []byte(`[{"id": 1,"key": "asdf"},{"id": 2,"key": "sdfg"},{"id": 3,"key": "wert"}]`)
	keys := make([]PublicKey, 0)
	json.Unmarshal(keysBody, &keys)
	fmt.Printf("%#v", keys)

}
