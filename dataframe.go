package dataframe

import (
	"errors"
	"fmt"
	"time"
)

// DataFrame represents a data structure for storing tabular data
type DataFrame struct {
	header  []string
	columns map[string][]interface{}
}

// NewDataFrame creates a new DataFrame with the given column names
func NewDataFrame(columnNames []string) (*DataFrame, error) {
	if len(columnNames) == 0 {
		return nil, errors.New("column names are required")
	}

	columns := make(map[string][]interface{})
	for _, columnName := range columnNames {
		if _, ok := columns[columnName]; ok {
			return nil, fmt.Errorf("column name '%s' already exists", columnName)
		}
		columns[columnName] = []interface{}{}
	}

	return &DataFrame{
		header:  columnNames,
		columns: columns,
	}, nil
}

// AddColumn adds a new column to the DataFrame
func (df *DataFrame) AddColumn(name string, data []interface{}) error {
	if len(data) != df.RowCount() {
		return errors.New("data length does not match row count")
	}

	if _, ok := df.columns[name]; ok {
		return fmt.Errorf("column name '%s' already exists", name)
	}

	df.columns[name] = data
	return nil
}

// ModifyColumn modifies an existing column in the DataFrame
func (df *DataFrame) ModifyColumn(name string, data []interface{}) error {
	if len(data) != df.RowCount() {
		return errors.New("data length does not match row count")
	}

	if _, ok := df.columns[name]; !ok {
		return fmt.Errorf("column '%s' does not exist", name)
	}

	df.columns[name] = data
	return nil
}

// ChangeColumnOrder changes the order of columns in the DataFrame
func (df *DataFrame) ChangeColumnOrder(newOrder []string) error {
	if len(newOrder) != len(df.header) {
		return errors.New("invalid column order")
	}

	// Check if all new column names exist in the DataFrame
	columnExists := make(map[string]bool)
	for _, columnName := range df.header {
		columnExists[columnName] = true
	}
	for _, columnName := range newOrder {
		if _, ok := columnExists[columnName]; !ok {
			return fmt.Errorf("column '%s' does not exist", columnName)
		}
	}

	// Rearrange the columns based on the new order
	newColumns := make(map[string][]interface{})
	for _, columnName := range newOrder {
		newColumns[columnName] = df.columns[columnName]
	}

	df.header = newOrder
	df.columns = newColumns

	return nil
}

// RowCount returns the number of rows in the DataFrame
func (df *DataFrame) RowCount() int {
	if len(df.columns) == 0 {
		return 0
	}

	return len(df.columns[df.header[0]])
}

// ColumnNames returns the names of the columns in the DataFrame
func (df *DataFrame) ColumnNames() []string {
	return df.header
}

// PrintHeader prints the header of the DataFrame
func (df *DataFrame) PrintHeader() {
	for _, columnName := range df.header {
		fmt.Printf("%v\t", columnName)
	}
	fmt.Println()
}

// PrintData prints the data in the DataFrame
func (df *DataFrame) PrintData() {
	for i := 0; i < df.RowCount(); i++ {
		for _, columnName := range df.header {
			fmt.Printf("%v\t", df.columns[columnName][i])
		}
		fmt.Println()
	}
}

// Filter applies a filter to the DataFrame based on a given condition
func (df *DataFrame) Filter(condition func(row int) bool) (*DataFrame, error) {
	filteredColumns := make(map[string][]interface{})
	for columnName, columnData := range df.columns {
		filteredColumns[columnName] = make([]interface{}, 0)
		for i := 0; i < df.RowCount(); i++ {
			if condition(i) {
				filteredColumns[columnName] = append(filteredColumns[columnName], columnData[i])
			}
		}
	}

	if len(filteredColumns) == 0 {
		return nil, errors.New("no columns matched the filter condition")
	}

	return &DataFrame{
		header:  df.header,
		columns: filteredColumns,
	}, nil
}

// Count returns the number of non-nil values in a column
func (df *DataFrame) Count(columnName string) (int, error) {
	columnData, ok := df.columns[columnName]
	if !ok {
		return 0, fmt.Errorf("column '%s' does not exist", columnName)
	}

	count := 0
	for _, value := range columnData {
		if value != nil {
			count++
		}
	}
	return count, nil
}

// Sum returns the sum of values in a numeric column
func (df *DataFrame) Sum(columnName string) (float64, error) {
	columnData, ok := df.columns[columnName]
	if !ok {
		return 0, fmt.Errorf("column '%s' does not exist", columnName)
	}

	sum := 0.0
	for _, value := range columnData {
		if numericValue, ok := value.(float64); ok {
			sum += numericValue
		} else {
			return 0, fmt.Errorf("column '%s' is not numeric", columnName)
		}
	}
	return sum, nil
}

// Mean returns the mean (average) of values in a numeric column
func (df *DataFrame) Mean(columnName string) (float64, error) {
	count, err := df.Count(columnName)
	if err != nil {
		return 0, err
	}

	sum, err := df.Sum(columnName)
	if err != nil {
		return 0, err
	}

	if count == 0 {
		return 0, errors.New("no values in column")
	}

	mean := sum / float64(count)
	return mean, nil
}

// TimeSeriesAnalysis performs a basic time series analysis on a date-based column
func (df *DataFrame) TimeSeriesAnalysis(dateColumnName, valueColumnName string) error {
	dateData, ok := df.columns[dateColumnName]
	if !ok {
		return fmt.Errorf("date column '%s' does not exist", dateColumnName)
	}

	valueData, ok := df.columns[valueColumnName]
	if !ok {
		return fmt.Errorf("value column '%s' does not exist", valueColumnName)
	}

	if len(dateData) != len(valueData) {
		return errors.New("date and value columns have different lengths")
	}

	if len(dateData) == 0 {
		return errors.New("no data in the DataFrame")
	}

	fmt.Println("Time Series Analysis:")
	fmt.Printf("Date Range: %v - %v\n", dateData[0], dateData[len(dateData)-1])

	// Perform analysis on value column (e.g., min, max, average)
	minValue, maxValue, err := df.MinValueMaxValue(valueColumnName)
	if err != nil {
		return err
	}

	fmt.Printf("Min %v: %v\n", valueColumnName, minValue)
	fmt.Printf("Max %v: %v\n", valueColumnName, maxValue)

	// Additional analysis can be performed as per requirements

	return nil
}

// MinValueMaxValue returns the minimum and maximum values in a numeric column
func (df *DataFrame) MinValueMaxValue(columnName string) (float64, float64, error) {
	columnData, ok := df.columns[columnName]
	if !ok {
		return 0, 0, fmt.Errorf("column '%s' does not exist", columnName)
	}

	minValue := columnData[0].(float64)
	maxValue := columnData[0].(float64)
	for _, value := range columnData {
		if numericValue, ok := value.(float64); ok {
			if numericValue < minValue {
				minValue = numericValue
			}
			if numericValue > maxValue {
				maxValue = numericValue
			}
		} else {
			return 0, 0, fmt.Errorf("column '%s' is not numeric", columnName)
		}
	}

	return minValue, maxValue, nil
}
