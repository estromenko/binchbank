package credits

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

func parseJSON(file *os.File) ([]Credit, error) {
	var operations []Credit

	err := json.NewDecoder(file).Decode(&operations)
	if err != nil {
		return nil, err
	}

	return operations, nil
}

func parseCSV(file *os.File) ([]Credit, error) {
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()

	if err != nil {
		return nil, err
	}

	var operations []Credit

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

		operation := Credit{
			ID: uint(id),
		}

		operations = append(operations, operation)
	}

	return operations, nil
}
