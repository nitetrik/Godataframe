# DataFrame Module

The DataFrame module provides a data structure for storing and manipulating tabular data in Go. It offers functionalities similar to popular DataFrame libraries in other programming languages, allowing you to work with structured data efficiently.

## Features

- Create a DataFrame with column names
- Add columns to the DataFrame
- Modify existing columns in the DataFrame
- Change the order of columns in the DataFrame
- Count non-nil values in a column
- Sum values in a numeric column
- Calculate the mean (average) of values in a numeric column
- Filter the DataFrame based on conditions
- Perform basic time series analysis
- Import data from Excel files into a DataFrame
- Export DataFrame data to Excel files

## Installation

To use the DataFrame module in your Go project, you need to include it as a dependency:

```bash
go get github.com/nitetrik/Godataframe/dataframe
```
## Usage

```go
package main

import (
	"fmt"
	"github.com/nitetrik/Godataframe/dataframe"
)

func main() {
	// Create a new DataFrame
	df, err := dataframe.NewDataFrame([]string{"Name", "Age", "City", "Salary"})
	if err != nil {
		fmt.Println("Error creating DataFrame:", err)
		return
	}

	// Add columns to the DataFrame
	names := []interface{}{"John", "Alice", "Bob", "Jane"}
	df.AddColumn("Name", names)

	ages := []interface{}{25, 30, 35, 28}
	df.AddColumn("Age", ages)

	cities := []interface{}{"New York", "London", "Tokyo", "Paris"}
	df.AddColumn("City", cities)

	salaries := []interface{}{50000, 60000, 75000, 55000}
	df.AddColumn("Salary", salaries)

	// Perform operations on the DataFrame

	// Print header and data of the DataFrame
	fmt.Println("DataFrame Header:")
	df.PrintHeader()
	fmt.Println()

	fmt.Println("DataFrame Data:")
	df.PrintData()
	fmt.Println()

	// Count the number of ages
	count, err := df.Count("Age")
	if err != nil {
		fmt.Println("Error counting values:", err)
		return
	}
	fmt.Println("Count of Ages:", count)
	fmt.Println()

	// Sum the salaries
	sum, err := df.Sum("Salary")
	if err != nil {
		fmt.Println("Error summing values:", err)
		return
	}
	fmt.Println("Sum of Salaries:", sum)
	fmt.Println()

	// Calculate the mean age
	mean, err := df.Mean("Age")
	if err != nil {
		fmt.Println("Error calculating mean:", err)
		return
	}
	fmt.Println("Mean Age:", mean)
	fmt.Println()

	// Filter the DataFrame
	filteredDF, err := df.Filter(func(row int) bool {
		return df.columns["Age"][row].(int) > 28
	})
	if err != nil {
		fmt.Println("Error filtering DataFrame:", err)
		return
	}

	fmt.Println("Filtered DataFrame:")
	filteredDF.PrintHeader()
	filteredDF.PrintData()
	fmt.Println()

	// Export DataFrame to Excel file
	err = df.ExportExcel("data.xlsx")
	if err != nil {
		fmt.Println("Error exporting DataFrame to Excel file:", err)
		return
	}

	fmt.Println("Exported DataFrame to Excel file successfully.")
}
```
## Contributing
Contributions to the DataFrame module are welcome! If you encounter any issues or have suggestions for improvement, please create an issue on the GitHub repository.

## License
This DataFrame module is open source and available under the MIT License.
