package main

import (
	"fmt"
	"time"

	"github.com/nitetrik/Godataframe/dataframe"
)

func main() {
	// Import Excel file
	df, err := dataframe.ImportExcel("data.xlsx")
	if err != nil {
		fmt.Println("Error importing Excel file:", err)
		return
	}

	// Perform operations on DataFrame

	// Print header and data of the DataFrame
	fmt.Println("DataFrame Header:")
	df.PrintHeader()
	fmt.Println()

	fmt.Println("DataFrame Data:")
	df.PrintData()
	fmt.Println()

	// Perform count operation
	count, err := df.Count("Age")
	if err != nil {
		fmt.Println("Error counting values:", err)
		return
	}
	fmt.Println("Count of Ages:", count)
	fmt.Println()

	// Perform sum operation
	sum, err := df.Sum("Salary")
	if err != nil {
		fmt.Println("Error summing values:", err)
		return
	}
	fmt.Println("Sum of Salaries:", sum)
	fmt.Println()

	// Perform mean operation
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

	// Modify a column in the DataFrame
	newAges := []interface{}{30, 35, 40, 33}
	err = df.ModifyColumn("Age", newAges)
	if err != nil {
		fmt.Println("Error modifying column:", err)
		return
	}

	fmt.Println("Modified DataFrame:")
	df.PrintHeader()
	df.PrintData()
	fmt.Println()

	// Change the column order in the DataFrame
	newColumnOrder := []string{"City", "Name", "Age", "Salary"}
	err = df.ChangeColumnOrder(newColumnOrder)
	if err != nil {
		fmt.Println("Error changing column order:", err)
		return
	}

	fmt.Println("DataFrame with Changed Column Order:")
	df.PrintHeader()
	df.PrintData()
	fmt.Println()

	// Perform time series analysis
	dateColumn := []interface{}{
		time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 2, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC),
	}
	valueColumn := []interface{}{100.0, 150.0, 120.0, 200.0}

	err = df.AddColumn("Date", dateColumn)
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}

	err = df.AddColumn("Value", valueColumn)
	if err != nil {
		fmt.Println("Error adding column:", err)
		return
	}

	err = df.TimeSeriesAnalysis("Date", "Value")
	if err != nil {
		fmt.Println("Error performing time series analysis:", err)
		return
	}
}
