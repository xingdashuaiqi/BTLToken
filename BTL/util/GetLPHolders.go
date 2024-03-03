package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetLPHolders(contractAddress, baseAPIUrl, apiKey string) []string {
	apiURL := fmt.Sprintf("%s?module=account&action=txlist&address=%s&apikey=%s", baseAPIUrl, contractAddress, apiKey)

	response, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}

	status, ok := result["status"].(string)
	if !ok || status != "1" {
		log.Println("API request failed")
		return nil
	}

	transactions, ok := result["result"].([]interface{})
	if !ok {
		log.Println("Failed to parse transactions")
		return nil
	}

	lpHolders := make(map[string]struct{})

	for _, tx := range transactions {
		txInfo := tx.(map[string]interface{})
		from := txInfo["from"].(string)
		to := txInfo["to"].(string)

		// 如果交易的输入地址或输出地址与合约地址相同，说明涉及到该合约
		if strings.EqualFold(from, contractAddress) || strings.EqualFold(to, contractAddress) {
			lpHolders[from] = struct{}{}
			lpHolders[to] = struct{}{}
		}
	}

	// 将 LP 持有人地址放入切片
	var resultHolders []string
	for holder := range lpHolders {

		resultHolders = append(resultHolders, holder)

	}
	resultHolders = removeDuplicates(resultHolders)
	return resultHolders
}

func removeDuplicates(slice []string) []string {
	seen := make(map[string]struct{}) // 使用空结构体作为 map 的值，只关心 key

	var result []string

	for _, value := range slice {
		if _, ok := seen[value]; !ok {
			// 如果字符串还没有出现过，则将其添加到结果切片中，并将其记录到 map 中
			seen[value] = struct{}{}
			result = append(result, value)
		}
	}

	return result
}
