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
			ID: uint(id),
		}

		operations = append(operations, operation)
	}

	return operations, nil
}
