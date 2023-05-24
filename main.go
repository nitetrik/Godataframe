package main

import (
	"fmt"
	"log"

	"github.com/your-username/your-module/dataframe"
)

func main() {
	// Create a new dataframe
	df := dataframe.NewDataFrame([]string{"Name", "Age", "City"})

	// Add rows to the dataframe
	df.AddRow("John", 25, "New York")
	df.AddRow("Alice", 30, "London")
	df.AddRow("Bob", 28, "Paris")
	df.AddRow("Emily", 22, "Tokyo")

	// Print the dataframe
	fmt.Println("Original DataFrame:")
	df.Print()
	fmt.Println()

	// Rename a column
	err := df.RenameColumn("City", "Location")
	if err != nil {
		log.Fatal(err)
	}

	// Add a new column
	err = df.AddColumn("Gender", []string{"Male", "Female", "Male", "Female"})
	if err != nil {
		log.Fatal(err)
	}

	// Remove a column
	err = df.RemoveColumn("Age")
	if err != nil {
		log.Fatal(err)
	}

	// Reorder columns
	err = df.ReorderColumns([]string{"Name", "Location", "Gender"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DataFrame after column manipulation:")
	df.Print()
	fmt.Println()

	// Sort the dataframe by name in ascending order
	err = df.Sort([]string{"Name"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DataFrame after sorting:")
	df.Print()
	fmt.Println()

	// Filter rows based on a condition
	filteredDF := df.Filter(func(row []interface{}) bool {
		location := row[1].(string)
		return location == "London"
	})

	fmt.Println("Filtered DataFrame:")
	filteredDF.Print()
	fmt.Println()

	// Compute descriptive statistics
	min, err := df.Min("Age")
	if err != nil {
		log.Fatal(err)
	}

	max, err := df.Max("Age")
	if err != nil {
		log.Fatal(err)
	}

	median, err := df.Median("Age")
	if err != nil {
		log.Fatal(err)
	}

	variance, err := df.Variance("Age")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Descriptive Statistics:\nMin: %v\nMax: %v\nMedian: %v\nVariance: %v\n\n", min, max, median, variance)

	// Merge two dataframes
	df2 := dataframe.NewDataFrame([]string{"Name", "Location", "Salary"})
	df2.AddRow("John", "New York", 5000)
	df2.AddRow("Alice", "London", 6000)
	df2.AddRow("Bob", "Paris", 5500)

	mergedDF, err := df.Merge(df2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Merged DataFrame:")
	mergedDF.Print()
	fmt.Println()

	// Convert column data type
	err = df.ConvertColumnType("Salary", dataframe.KindFloat64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DataFrame after column data type conversion:")
	df.Print()
	fmt.Println()

	// Infer column data types
	df.InferColumnTypes()

	fmt.Println("DataFrame after inferring column data types:")
	df.Print()
	fmt.Println()

	// Fill missing values
	df.FillMissingValues("N/A")

	fmt.Println("DataFrame after filling missing values:")
	df.Print()
	fmt.Println()

	// Group by location and compute the average salary
	groupedDF, err := df.GroupBy("Location", func(data []interface{}) interface{} {
		salaries := make([]float64, 0)
		for _, val := range data {
			if salary, ok := val.(float64); ok {
				salaries = append(salaries, salary)
			}
		}
		if len(salaries) > 0 {
			sum := 0.0
			for _, salary := range salaries {
				sum += salary
			}
			return sum / float64(len(salaries))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grouped DataFrame:")
	groupedDF.Print()
	fmt.Println()

	// Pivot the dataframe
	pivotedDF, err := df.Pivot("Name", "Location", "Salary")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pivoted DataFrame:")
	pivotedDF.Print()
}
