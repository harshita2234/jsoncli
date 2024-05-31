# jsoncli

![Go](https://img.shields.io/badge/Go-100%25-blue)

`jsoncli` is a versatile command-line interface (CLI) tool designed to simplify the handling and manipulation of JSON files. With a range of features including validation, formatting, merging, and the ability to view the first and last rows of JSON arrays, `jsoncli` is an essential utility for developers and data analysts working with JSON data.

## Features

- **Input JSON File**: Specify the path of the input JSON file using the `-input` flag.
- **Output JSON File**: Define the path for the output formatted JSON file with the `-output` flag.
- **Validate JSON**: Validate the JSON file for correctness using the `-validate` flag.
- **Format JSON**: Format the JSON file with indentation for better readability using the `-format` flag.
- **Merge JSON Files**: Merge another JSON file with the input file using the `-merge` flag.
- **Display First 5 Rows**: Display the first 5 rows of a JSON array with the `-head` flag.
- **Display Last 5 Rows**: Display the last 5 rows of a JSON array with the `-tail` flag.

## Usage

1. **Validate a JSON File**:
    ```sh
    ./jsoncli -input=path/to/input.json -validate
    ```

2. **Format a JSON File**:
    ```sh
    ./jsoncli -input=path/to/input.json -format
    ```

3. **Merge Two JSON Files**:
    ```sh
    ./jsoncli -input=path/to/input.json -merge=path/to/merge.json -output=merged.json
    ```

4. **Display the First 5 Rows of a JSON Array**:
    ```sh
    ./jsoncli -input=path/to/input.json -head
    ```

5. **Display the Last 5 Rows of a JSON Array**:
    ```sh
    ./jsoncli -input=path/to/input.json -tail
    ```

6. **Format and Validate a JSON File**:
    ```sh
    ./jsoncli -input=path/to/input.json -format -validate -output=formatted.json
    ```

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/harshita2234/jsoncli.git
    ```

2. Build the CLI tool:
    ```sh
    cd jsoncli
    go build -o jsoncli
    ```

## Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request with your changes.

## License

This project is licensed under the MIT License.
