package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

func merge(a, b interface{}) interface{} { //interface is used for varying types of values in json file
	switch a := a.(type) {
	case map[string]interface{}:
		b, ok := b.(map[string]interface{})
		if !ok {
			return a
		}
		for k, v := range b {
			a[k] = merge(a[k], v)
		}
		return a
	case []interface{}:
		b, ok := b.([]interface{})
		if !ok {
			return a
		}
		return append(a, b...)
	default:
		return b
	}
}

func printHead(jsondata interface{}) {
	array, ok := jsondata.([]interface{})
	if !ok {
		fmt.Println("input JSON is not an array")
		return
	}

	length := len(array)
	if length == 0 {
		fmt.Println("JSON array is empty")
		return
	}

	end := 5
	if length < 5 {
		end = length
	}

	for i := 0; i < end; i++ {
		printTable(array[i])
	}
}

func printTail(jsondata interface{}) {
	array, ok := jsondata.([]interface{})
	if !ok {
		fmt.Println("input JSON is not an array")
		return
	}

	length := len(array)
	if length == 0 {
		fmt.Println("JSON array is empty")
		return
	}

	start := length - 5
	if start < 0 {
		start = 0
	}

	for i := start; i < length; i++ {
		printTable(array[i])
	}
}

func printTable(row interface{}) {
	rowMap, ok := row.(map[string]interface{})
	if !ok {
		fmt.Println("row is not a JSON object")
		return
	}

	keys := make([]string, 0, len(rowMap))
	values := make([]string, 0, len(rowMap))

	for k, v := range rowMap {
		keys = append(keys, k)
		values = append(values, fmt.Sprintf("%v", v))
	}

	printRow(keys)
	printRow(values)
	fmt.Println(strings.Repeat("-", 50))
}

func printRow(fields []string) {
	for _, field := range fields {
		fmt.Printf("%-15s", field)
	}
	fmt.Println()
}

func main() {
	//command-line flags
	inputfile := flag.String("input", "", "-> path of input JSON file")
	mergefile := flag.String("string", "", "-> path of JSON file to be merged with input")
	outputfile := flag.String("output", "", "-> path of output formatter")
	validate := flag.Bool("validate", false, "-> validate the JSON file")
	format := flag.Bool("format", false, "-> format the JSON file")
	head := flag.Bool("head", false, "-> display first 5 rows of JSON array with header")
	tail := flag.Bool("tail", false, "-> display last 5 rows of JSON array with header")

	flag.Parse() //arguments from command-line

	//check for empty input argument
	if *inputfile == "" {
		fmt.Println("Provide input using -input flag")
		return
	}

	//read file contents using os package
	data, err := os.ReadFile(*inputfile)
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
		return
	}

	//unmarshal - lo level to high level
	var jsondata interface{} //jsondata will hold unmarshalled data
	err = json.Unmarshal(data, &jsondata)
	if err != nil {
		fmt.Printf("error unmarshalling file %v\n", err)
		return
	}
	//check for validity
	if *validate {
		fmt.Println("json file is valid")
	}

	if *format {
		formatted, err := json.MarshalIndent(jsondata, "", " ") //creates pretty json strings in formatted
		if err != nil {
			fmt.Printf("error marshalling file %v\n", err)
			return
		}
		//check if outputfile exists and then write on console if false
		if *outputfile == "" {
			fmt.Println(string(formatted))
		} else { //outputfile exists -> write formatted strings to file
			err = os.WriteFile(*outputfile, formatted, 0644)
			if err != nil {
				fmt.Printf("error writing file %v\n", err)
			}
			fmt.Println("formatted json saved to ", *outputfile)
		}
	}
	//check if mergefile exists
	if *mergefile != "" {
		mergedata, err := os.ReadFile(*mergefile) //read mergefile
		if err != nil {
			fmt.Printf("error reading merge file %v\n", err)
			return
		}
		var mergejsondata interface{} //mergejsondata will hold unmarshalled data
		err = json.Unmarshal(mergedata, &mergejsondata)
		if err != nil {
			fmt.Printf("error unmarshalling merge file %v\n", err)
			return
		}

		jsondata = merge(jsondata, mergejsondata) //implement merge function
	}
	//print first 5 rows
	if *head {
		printHead(jsondata)
	}
	//print last 5 rows
	if *tail {
		printTail(jsondata)
	}
}
