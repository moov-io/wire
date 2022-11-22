package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"syscall/js"

	"github.com/moov-io/wire"
)

func isJSON(input string) bool {
	var dummy json.RawMessage
	return json.Unmarshal([]byte(input), &dummy) == nil
}

func parseContents(input string) (*wire.File, error) {

	var file wire.File
	var err error

	if isJSON(input) {
		if err = json.Unmarshal([]byte(input), &file); err != nil {
			return nil, fmt.Errorf("unable to parse with json foramt")
		}
	} else {
		r := strings.NewReader(input)
		if file, err = wire.NewReader(r).Read(); err != nil {
			return nil, err
		}
	}

	return &file, nil
}

func prettyJson(file *wire.File) (string, error) {
	pretty, err := json.MarshalIndent(file, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func printWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}

		inputJSON := args[0].String()
		outFormat := args[1].String()

		file, err := parseContents(inputJSON)
		if err != nil {
			msg := fmt.Sprintf("unable to parse wire file - %v", err)
			fmt.Print(msg)
			return msg
		}

		if outFormat == "wire" {
			var buf bytes.Buffer

			w := wire.NewWriter(bufio.NewWriter(&buf))
			if err := w.Write(file); err != nil {
				fmt.Printf("unable to convert wire file to wire %s\n", err)
				return "There was an error converting the wire"
			}
			w.Flush()

			return buf.String()
		} else {
			pretty, err := prettyJson(file)
			if err != nil {
				fmt.Printf("unable to convert wire file to json %s\n", err)
				return "There was an error converting the json"
			}
			return pretty
		}
	})
	return jsonFunc
}

func main() {
	js.Global().Set("parseContents", printWrapper())
	<-make(chan bool)
}
