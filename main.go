package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	// Create a new DataFrame
	df := NewDataFrame([]string{"Name", "Age", "City"})

	// Add rows to the DataFrame
	df.AddRow("John", 30, "New York")
	df.AddRow("Jane", 25, "London")
	df.AddRow("Mike", 35, "Paris")

	// Print the original DataFrame
	fmt.Println("Original DataFrame:")
	df.Print()
	fmt.Println()

	// Get a column from the DataFrame
	ages, err := df.GetColumn("Age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Ages:", ages)
	fmt.Println()

	// Calculate the sum of values in a column
	sum, err := df.SumColumn("Age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Sum of ages:", sum)
	fmt.Println()

	// Filter the DataFrame based on a condition
	filteredDF := df.Filter(func(row []interface{}) bool {
		age := row[1].(int)
		return age >= 30
	})
	fmt.Println("Filtered DataFrame:")
	filteredDF.Print()
	fmt.Println()

	// Rename a column in the DataFrame
	err = df.RenameColumn("City", "Location")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Renamed DataFrame:")
	df.Print()
	fmt.Println()

	// Add a new column to the DataFrame
	population := []float64{8.5, 9, 10.2}
	err = df.AddColumn("Population (in millions)", population)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DataFrame with new column:")
	df.Print()
	fmt.Println()

	// Remove a column from the DataFrame
	err = df.RemoveColumn("Name")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DataFrame with column removed:")
	df.Print()
	fmt.Println()

	// Reorder columns in the DataFrame
	err = df.ReorderColumns([]string{"Location", "Population (in millions)", "Age"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DataFrame with reordered columns:")
	df.Print()
	fmt.Println()

	// Sort the DataFrame based on one or more columns
	err = df.Sort([]string{"Age"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DataFrame sorted by Age:")
	df.Print()
	fmt.Println()

	// Perform statistical computations on a numeric column
	min, err := df.Min("Age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Minimum age:", min)

	max, err := df.Max("Age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Maximum age:", max)

	median, err := df.Median("Age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Median age:", median)

	variance, err := df.Variance("Age")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Variance of age:", variance)
	fmt.Println()

	// Merge two DataFrames based on common columns
	otherDF := NewDataFrame([]string{"Name", "Location"})
	otherDF.AddRow("John", "New York")
	otherDF.AddRow("Jane", "London")
	otherDF.AddRow("Mike", "Paris")

	mergedDF, err := df.Merge(otherDF)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Merged DataFrame:")
	mergedDF.Print()
	fmt.Println()

	// Convert the data type of a column
	err = df.ConvertColumnType("Age", reflect.Float64)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DataFrame with converted column type:")
	df.Print()
	fmt.Println()

	// Infer the data types of columns
	dataTypes := df.InferDataType()
	fmt.Println("Inferred data types:")
	for col, dataType := range dataTypes {
		fmt.Printf("%s: %s\n", col, dataType)
	}
	fmt.Println()

	// Fill missing values in the DataFrame
	df.FillMissingValues(0)

	fmt.Println("DataFrame with filled missing values:")
	df.Print()
	fmt.Println()

	// Group the DataFrame by a column and perform aggregation
	groupedDF, err := df.GroupBy("Location", func(data []interface{}) interface{} {
		count := len(data)
		return count
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Grouped DataFrame:")
	groupedDF.Print()
	fmt.Println()

	// Pivot the DataFrame based on specified columns
	pivotedDF, err := df.Pivot("Age", "Location", "Population (in millions)")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pivoted DataFrame:")
	pivotedDF.Print()
	fmt.Println()
}
