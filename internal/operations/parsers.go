package operations

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

func parseJSON(file *os.File) ([]Operation, error) {
	var operations []Operation

	err := json.NewDecoder(file).Decode(&operations)
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func parseCSV(file *os.File) ([]Operation, error) {
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var operations []Operation

	headers := data[0]

	for _, v := range data[1:] {
		if len(v) < 1 {
			continue
		}

		line := make(map[string]interface{}, len(v))

		for ii, vv := range v {
			line[headers[ii]] = vv
		}

		idstr := line["id"].(string)
		id, err := strconv.Atoi(idstr)
		if err != nil {
			return nil, err
		}

		operation := Operation{
<<<<<<< HEAD
			ID: uint(id),
=======
			ID:       uint(id),
			Username: line["username"].(string),
>>>>>>> 03b2f2fd2bc76b58d6ae3aea5c3694b959a8bac3
		}

		operations = append(operations, operation)
	}

	return operations, nil
}
