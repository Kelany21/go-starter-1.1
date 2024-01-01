package helpers

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ArrayToString(a []int) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", ",", -1), "[]")
}

func GetIDs(arr interface{}) []int {
	var ids []int
	type ID struct {
		ID uint `json:"id"`
	}
	var idsArr []ID
	arrBytes, _ := json.Marshal(arr)
	_ = json.Unmarshal(arrBytes, &idsArr)
	for i := 0; i < len(idsArr); i++ {
		ids = append(ids, int(idsArr[i].ID))
	}
	return ids
}

func ConvertArrayIntToArrayInterfaces(arr []int) []interface{} {
	var ids []interface{}
	for i := 0; i < len(arr); i++ {
		ids = append(ids, arr[i])
	}
	return ids
}

func ConvertArrayInterfacesToArrayStrings(arr []interface{}) []string {
	var ids []string

	for i := 0; i < len(arr); i++ {
		ids = append(ids, arr[i].(string))
	}
	return ids
}
