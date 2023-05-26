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

	// Sort the DataFrame
	err = df.Sort([]string{"Age", "Salary"}, true)
	if err != nil {
		fmt.Println("Error sorting DataFrame:", err)
		return
	}

	fmt.Println("Sorted DataFrame:")
	df.PrintHeader()
	df.PrintData()
	fmt.Println()

	// Group and aggregate the DataFrame
	groupedDF, err := df.GroupBy([]string{"City"})
	if err != nil {
		fmt.Println("Error grouping DataFrame:", err)
		return
	}

	fmt.Println("Grouped DataFrame:")
	groupedDF.PrintHeader()
	groupedDF.PrintData()
	fmt.Println()

	// Join DataFrames
	joinDF1, err := dataframe.NewDataFrame([]string{"Name", "Age", "City"})
	if err != nil {
		fmt.Println("Error creating join DataFrame 1:", err)
		return
	}

	joinNames := []interface{}{"John", "Alice"}
	joinDF1.AddColumn("Name", joinNames)

	joinAges := []interface{}{25, 30}
	joinDF1.AddColumn("Age", joinAges)

	joinCities := []interface{}{"New York", "London"}
	joinDF1.AddColumn("City", joinCities)

	joinDF2, err := dataframe.NewDataFrame([]string{"City", "Population"})
	if err != nil {
		fmt.Println("Error creating join DataFrame 2:", err)
		return
	}

	joinCities2 := []interface{}{"New York", "London"}
	joinDF2.AddColumn("City", joinCities2)

	joinPopulations := []interface{}{8500000, 9000000}
	joinDF2.AddColumn("Population", joinPopulations)

	joinedDF, err := dataframe.Join([]*dataframe.DataFrame{joinDF1, joinDF2}, []string{"City"})
	if err != nil {
		fmt.Println("Error joining DataFrames:", err)
		return
	}

	fmt.Println("Joined DataFrame:")
	joinedDF.PrintHeader()
	joinedDF.PrintData()
	fmt.Println()

	// Serialize DataFrame to JSON
	jsonData, err := df.SerializeToJSON()
	if err != nil {
		fmt.Println("Error serializing DataFrame to JSON:", err)
		return
	}

	fmt.Println("Serialized DataFrame (JSON):")
	fmt.Println(jsonData)
	fmt.Println()

	// Serialize DataFrame to CSV
	csvData, err := df.SerializeToCSV()
	if err != nil {
		fmt.Println("Error serializing DataFrame to CSV:", err)
		return
	}

	fmt.Println("Serialized DataFrame (CSV):")
	fmt.Println(csvData)
	fmt.Println()
}
