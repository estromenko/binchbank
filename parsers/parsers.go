package parsers

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func Parse(file *os.File, dest interface{}) error {
	nameParts := strings.Split(file.Name(), ".")

	if len(nameParts) <= 1 {
		return fmt.Errorf("invalid file extension")
	}

	ext := nameParts[len(nameParts)-1]

	if ext == "json" {
		return ParseJSON(file, dest)
	}

	if ext == "csv" {
		return ParseCSV(file, dest)
	}

	return fmt.Errorf("unsupported file path")
}

func ParseJSON(file io.Reader, dest interface{}) error {
	return json.NewDecoder(file).Decode(&dest)
}

func ParseCSV(file io.Reader, dest interface{}) error {
	reader := csv.NewReader(file)
	content, _ := reader.ReadAll()

	if len(content) < 1 {
		return fmt.Errorf("invalid csv file")
	}

	headersArr := make([]string, 0)
	headersArr = append(headersArr, content[0]...)

	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headersArr[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}

			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}

		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}

	buffer.WriteString(`]`)
	rawMessage := json.RawMessage(buffer.String())
	data, _ := json.MarshalIndent(rawMessage, "", "  ")

	return ParseJSON(bytes.NewReader(data), &dest)
}
