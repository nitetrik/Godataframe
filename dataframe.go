package dataframe

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
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

// Sort sorts the DataFrame based on one or more columns in ascending or descending order
func (df *DataFrame) Sort(columns []string, ascending bool) error {
	sort.SliceStable(df.columns[df.header[0]], func(i, j int) bool {
		for _, col := range columns {
			if df.columns[col][i] != df.columns[col][j] {
				switch df.columns[col][i].(type) {
				case int:
					if ascending {
						return df.columns[col][i].(int) < df.columns[col][j].(int)
					} else {
						return df.columns[col][i].(int) > df.columns[col][j].(int)
					}
				case float64:
					if ascending {
						return df.columns[col][i].(float64) < df.columns[col][j].(float64)
					} else {
						return df.columns[col][i].(float64) > df.columns[col][j].(float64)
					}
				case string:
					if ascending {
						return strings.Compare(df.columns[col][i].(string), df.columns[col][j].(string)) < 0
					} else {
						return strings.Compare(df.columns[col][i].(string), df.columns[col][j].(string)) > 0
					}
				}
			}
		}
		return i < j
	})

	return nil
}

// GroupBy groups the DataFrame by one or more columns
func (df *DataFrame) GroupBy(columns []string) (*DataFrame, error) {
	groupedColumns := make(map[string][]interface{})
	groupedRowCount := make(map[string]int)

	for _, col := range columns {
		if _, ok := df.columns[col]; !ok {
			return nil, fmt.Errorf("column '%s' does not exist", col)
		}
		groupedColumns[col] = []interface{}{}
		groupedRowCount[col] = 0
	}

	for i := 0; i < df.RowCount(); i++ {
		groupKey := ""
		for _, col := range columns {
			groupKey += fmt.Sprintf("%v-%v", col, df.columns[col][i])
		}

		if _, ok := groupedColumns[groupKey]; !ok {
			for _, col := range columns {
				groupedColumns[col] = append(groupedColumns[col], df.columns[col][i])
			}
		}
		groupedRowCount[groupKey]++
	}

	for col, rowCount := range groupedRowCount {
		for i := 0; i < df.RowCount(); i++ {
			if groupedRowCount[fmt.Sprintf("%v-%v", col, df.columns[col][i])] > 0 {
				groupedColumns[col] = append(groupedColumns[col], rowCount)
				groupedRowCount[fmt.Sprintf("%v-%v", col, df.columns[col][i])] = 0
			} else {
				groupedColumns[col] = append(groupedColumns[col], nil)
			}
		}
	}

	return &DataFrame{
		header:  append(columns, "Count"),
		columns: groupedColumns,
	}, nil
}

// Join joins multiple DataFrames based on common columns
func Join(dataFrames []*DataFrame, joinColumns []string) (*DataFrame, error) {
	if len(dataFrames) == 0 {
		return nil, errors.New("no DataFrames provided for join")
	}

	joinedColumns := make(map[string][]interface{})
	for _, df := range dataFrames {
		for _, col := range df.header {
			if _, ok := joinedColumns[col]; ok {
				return nil, fmt.Errorf("column '%s' already exists in the join result", col)
			}
			joinedColumns[col] = df.columns[col]
		}
	}

	for _, df := range dataFrames[1:] {
		for i := 0; i < df.RowCount(); i++ {
			rowMatch := make([]bool, len(dataFrames[0].header))
			for j, col := range joinColumns {
				if _, ok := df.columns[col]; !ok {
					return nil, fmt.Errorf("column '%s' does not exist in DataFrame for join", col)
				}

				if _, ok := joinedColumns[col]; !ok {
					return nil, fmt.Errorf("column '%s' does not exist in the join result", col)
				}

				if df.columns[col][i] != nil {
					if idx := findIndex(joinedColumns[col], df.columns[col][i]); idx != -1 {
						rowMatch[idx] = true
					}
				}
			}

			if allTrue(rowMatch) {
				for _, col := range df.header {
					if _, ok := joinedColumns[col]; !ok {
						return nil, fmt.Errorf("column '%s' does not exist in the join result", col)
					}

					joinedColumns[col] = append(joinedColumns[col], df.columns[col][i])
				}
			}
		}
	}

	return &DataFrame{
		header:  dataFrames[0].header,
		columns: joinedColumns,
	}, nil
}

// CleanData cleans the DataFrame by handling missing values, duplicates, and data type conversion
func (df *DataFrame) CleanData() error {
	// Handle missing values
	for columnName, columnData := range df.columns {
		for i := 0; i < len(columnData); i++ {
			if columnData[i] == nil {
				switch df.columns[columnName][0].(type) {
				case int:
					df.columns[columnName][i] = 0
				case float64:
					df.columns[columnName][i] = 0.0
				case string:
					df.columns[columnName][i] = ""
				default:
					return fmt.Errorf("unknown data type in column '%s'", columnName)
				}
			}
		}
	}

	// Handle duplicates
	duplicateIndexes := make(map[int]bool)
	for i := 0; i < df.RowCount(); i++ {
		if duplicateIndexes[i] {
			continue
		}
		for j := i + 1; j < df.RowCount(); j++ {
			if duplicateIndexes[j] {
				continue
			}
			isDuplicate := true
			for _, columnName := range df.header {
				if df.columns[columnName][i] != df.columns[columnName][j] {
					isDuplicate = false
					break
				}
			}
			if isDuplicate {
				duplicateIndexes[j] = true
			}
		}
	}

	for columnName, columnData := range df.columns {
		cleanedColumn := make([]interface{}, 0)
		for i := 0; i < len(columnData); i++ {
			if !duplicateIndexes[i] {
				cleanedColumn = append(cleanedColumn, columnData[i])
			}
		}
		df.columns[columnName] = cleanedColumn
	}

	return nil
}

// Variance calculates the variance of values in a numeric column
func (df *DataFrame) Variance(columnName string) (float64, error) {
	columnData, ok := df.columns[columnName]
	if !ok {
		return 0, fmt.Errorf("column '%s' does not exist", columnName)
	}

	count, err := df.Count(columnName)
	if err != nil {
		return 0, err
	}

	mean, err := df.Mean(columnName)
	if err != nil {
		return 0, err
	}

	if count <= 1 {
		return 0, errors.New("insufficient data points for variance calculation")
	}

	variance := 0.0
	for _, value := range columnData {
		if numericValue, ok := value.(float64); ok {
			variance += (numericValue - mean) * (numericValue - mean)
		} else {
			return 0, fmt.Errorf("column '%s' is not numeric", columnName)
		}
	}
	variance /= float64(count - 1)

	return variance, nil
}

// StandardDeviation calculates the standard deviation of values in a numeric column
func (df *DataFrame) StandardDeviation(columnName string) (float64, error) {
	variance, err := df.Variance(columnName)
	if err != nil {
		return 0, err
	}

	standardDeviation := 0.0
	if variance > 0 {
		standardDeviation = �math.Sqrt(variance)
	}

	return standardDeviation, nil
}

// Correlation calculates the correlation coefficient between two numeric columns
func (df *DataFrame) Correlation(column1, column2 string) (float64, error) {
	column1Data, ok1 := df.columns[column1]
	column2Data, ok2 := df.columns[column2]

	if !ok1 {
		return 0, fmt.Errorf("column '%s' does not exist", column1)
	}
	if !ok2 {
		return 0, fmt.Errorf("column '%s' does not exist", column2)
	}

	count, err := df.Count(column1)
	if err != nil {
		return 0, err
	}

	if count <= 1 {
		return 0, errors.New("insufficient data points for correlation calculation")
	}

	var (
		sumXY    float64
		sumX     float64
		sumY     float64
		sumXSquare float64
		sumYSquare float64
	)

	for i := 0; i < df.RowCount(); i++ {
		if value1, ok := column1Data[i].(float64); ok {
			if value2, ok := column2Data[i].(float64); ok {
				sumXY += value1 * value2
				sumX += value1
				sumY += value2
				sumXSquare += value1 * value1
				sumYSquare += value2 * value2
			} else {
				return 0, fmt.Errorf("column '%s' is not numeric", column2)
			}
		} else {
			return 0, fmt.Errorf("column '%s' is not numeric", column1)
		}
	}

	numerator := count*sumXY - sumX*sumY
	denominator := math.Sqrt((count*sumXSquare - sumX*sumX) * (count*sumYSquare - sumY*sumY))

	correlation := 0.0
	if denominator != 0 {
		correlation = numerator / denominator
	}

	return correlation, nil
}

// Covariance calculates the covariance between two numeric columns
func (df *DataFrame) Covariance(column1, column2 string) (float64, error) {
	column1Data, ok1 := df.columns[column1]
	column2Data, ok2 := df.columns[column2]

	if !ok1 {
		return 0, fmt.Errorf("column '%s' does not exist", column1)
	}
	if !ok2 {
		return 0, fmt.Errorf("column '%s' does not exist", column2)
	}

	count, err := df.Count(column1)
	if err != nil {
		return 0, err
	}

	if count <= 1 {
		return 0, errors.New("insufficient data points for covariance calculation")
	}

	var (
		sumXY float64
		sumX  float64
		sumY  float64
	)

	for i := 0; i < df.RowCount(); i++ {
		if value1, ok := column1Data[i].(float64); ok {
			if value2, ok := column2Data[i].(float64); ok {
				sumXY += value1 * value2
				sumX += value1
				sumY += value2
			} else {
				return 0, fmt.Errorf("column '%s' is not numeric", column2)
			}
		} else {
			return 0, fmt.Errorf("column '%s' is not numeric", column1)
		}
	}

	covariance := 0.0
	if count > 1 {
		meanX := sumX / float64(count)
		meanY := sumY / float64(count)
		covariance = (sumXY - float64(count)*meanX*meanY) / float64(count-1)
	}

	return covariance, nil
}

// SerializeToJSON serializes the DataFrame to a JSON string
func (df *DataFrame) SerializeToJSON() (string, error) {
	jsonData := "[\n"

	for i := 0; i < df.RowCount(); i++ {
		jsonData += "\t{"
		for j, columnName := range df.header {
			jsonData += fmt.Sprintf("\"%s\":", columnName)
			if value, ok := df.columns[columnName][i].(string); ok {
				jsonData += "\"" + value + "\""
			} else {
				jsonData += fmt.Sprintf("%v", df.columns[columnName][i])
			}
			if j < len(df.header)-1 {
				jsonData += ","
			}
		}
		jsonData += "}"
		if i < df.RowCount()-1 {
			jsonData += ","
		}
		jsonData += "\n"
	}

	jsonData += "]\n"

	return jsonData, nil
}

// SerializeToCSV serializes the DataFrame to a CSV string
func (df *DataFrame) SerializeToCSV() (string, error) {
	csvData := strings.Join(df.header, ",") + "\n"

	for i := 0; i < df.RowCount(); i++ {
		for j, columnName := range df.header {
			if value, ok := df.columns[columnName][i].(string); ok {
				csvData += "\"" + value + "\""
			} else {
				csvData += fmt.Sprintf("%v", df.columns[columnName][i])
			}
			if j < len(df.header)-1 {
				csvData += ","
			}
		}
		csvData += "\n"
	}

	return csvData, nil
}

// findIndex returns the index of an element in a slice, or -1 if not found
func findIndex(slice []interface{}, element interface{}) int {
	for i, value := range slice {
		if value == element {
			return i
		}
	}
	return -1
}

// allTrue checks if all elements in a bool slice are true
func allTrue(slice []bool) bool {
	for _, value := range slice {
		if !value {
			return false
		}
	}
	return true
}
