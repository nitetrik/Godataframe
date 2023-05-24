package dataframe

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// DataFrame represents a two-dimensional table of data.
type DataFrame[T any] struct {
	Headers []string
	Data    [][]T
}

// NewDataFrame creates a new DataFrame with the given columns.
func NewDataFrame[T any](columns []string) *DataFrame[T] {
	return &DataFrame[T]{
		Headers: columns,
		Data:    [][]T{},
	}
}

// AddRow adds a new row of data to the DataFrame.
func (df *DataFrame[T]) AddRow(rowData ...T) error {
	if len(rowData) != len(df.Headers) {
		return errors.New("number of values does not match the number of columns")
	}

	df.Data = append(df.Data, rowData)
	return nil
}

// Print displays the DataFrame in a tabular format.
func (df *DataFrame[T]) Print() {
	// Print column names
	for _, col := range df.Headers {
		fmt.Printf("%-15s", col)
	}
	fmt.Println()

	// Print data rows
	for _, row := range df.Data {
		for _, val := range row {
			fmt.Printf("%-15v", val)
		}
		fmt.Println()
	}
}

// GetColumn returns the values in a specific column of the DataFrame.
func (df *DataFrame[T]) GetColumn(columnName string) ([]T, error) {
	index := df.getColumnIndex(columnName)
	if index == -1 {
		return nil, errors.New("column not found")
	}

	column := []T{}
	for _, row := range df.Data {
		column = append(column, row[index])
	}

	return column, nil
}

// SumColumn calculates the sum of the values in a specific numeric column of the DataFrame.
func (df *DataFrame[T]) SumColumn(columnName string) (float64, error) {
	column, err := df.GetColumn(columnName)
	if err != nil {
		return 0, err
	}

	sum := 0.0
	for _, val := range column {
		if val, ok := val.(float64); ok {
			sum += val
		} else {
			return 0, errors.New("non-numeric value found in the column")
		}
	}

	return sum, nil
}

// Helper function to get the index of a column by name
func (df *DataFrame[T]) getColumnIndex(columnName string) int {
	for i, col := range df.Headers {
		if col == columnName {
			return i
		}
	}
	return -1
}

// Filter returns a new DataFrame containing only the rows that match the filter criteria.
func (df *DataFrame[T]) Filter(filter func(row []T) bool) *DataFrame[T] {
	newDF := NewDataFrame[T](df.Headers)

	for _, row := range df.Data {
		if filter(row) {
			newDF.AddRow(row...)
		}
	}

	return newDF
}

// MeanColumn returns the mean of the values in a specific numeric column of the DataFrame.
func (df *DataFrame[T]) MeanColumn(columnName string) (float64, error) {
	column, err := df.GetColumn(columnName)
	if err != nil {
		return 0, err
	}

	n := float64(len(column))
	if n == 0 {
		return 0, errors.New("no values in column")
	}

	sum := 0.0
	for _, val := range column {
		if val, ok := val.(float64); ok {
			sum += val
		} else {
			return 0, errors.New("non-numeric value found in the column")
		}
	}

	return sum / n, nil
}

// RenameColumn renames a column in the DataFrame.
func (df *DataFrame[T]) RenameColumn(oldName string, newName string) error {
	index := df.getColumnIndex(oldName)
	if index == -1 {
		return errors.New("column not found")
	}

	df.Headers[index] = newName
	return nil
}

// AddColumn adds a new column to the DataFrame.
func (df *DataFrame[T]) AddColumn(columnName string, values []T) error {
	if len(values) != len(df.Data) {
		return errors.New("number of values does not match the number of rows")
	}

	df.Headers = append(df.Headers, columnName)
	for i, val := range values {
		df.Data[i] = append(df.Data[i], val)
	}

	return nil
}

// RemoveColumn removes a column from the DataFrame.
func (df *DataFrame[T]) RemoveColumn(columnName string) error {
	index := df.getColumnIndex(columnName)
	if index == -1 {
		return errors.New("column not found")
	}

	df.Headers = append(df.Headers[:index], df.Headers[index+1:]...)
	for i := range df.Data {
		df.Data[i] = append(df.Data[i][:index], df.Data[i][index+1:]...)
	}

	return nil
}

// ReorderColumns reorders the columns based on the specified order.
func (df *DataFrame[T]) ReorderColumns(columnOrder []string) error {
	if len(columnOrder) != len(df.Headers) {
		return errors.New("number of columns in the order list does not match the dataframe")
	}

	newHeaders := make([]string, len(df.Headers))
	copy(newHeaders, df.Headers)

	newData := make([][]T, len(df.Data))
	for i, row := range df.Data {
		newData[i] = make([]T, len(row))
		copy(newData[i], row)
	}

	for i, col := range columnOrder {
		index := df.getColumnIndex(col)
		if index == -1 {
			return errors.New("column not found in the order list")
		}

		newHeaders[i] = col
		for j, row := range df.Data {
			newData[j][i] = row[index]
		}
	}

	df.Headers = newHeaders
	df.Data = newData

	return nil
}

// Sort sorts the dataframe based on one or more columns.
func (df *DataFrame[T]) Sort(columns []string) error {
	sortSlice := make([][]T, len(df.Data))
	for i, row := range df.Data {
		sortSlice[i] = make([]T, len(row))
		copy(sortSlice[i], row)
	}

	sort.SliceStable(sortSlice, func(i, j int) bool {
		for _, col := range columns {
			index := df.getColumnIndex(col)
			if index == -1 {
				return false
			}

			v1 := sortSlice[i][index]
			v2 := sortSlice[j][index]
			if v1 != v2 {
				return fmt.Sprintf("%v", v1) < fmt.Sprintf("%v", v2)
			}
		}
		return false
	})

	df.Data = sortSlice

	return nil
}

// FilterRows filters the dataframe rows based on specific conditions.
func (df *DataFrame[T]) FilterRows(filter func(row []T) bool) {
	filteredData := make([][]T, 0)
	for _, row := range df.Data {
		if filter(row) {
			filteredData = append(filteredData, row)
		}
	}
	df.Data = filteredData
}

// Min computes the minimum value for the specified numeric column.
func (df *DataFrame[T]) Min(columnName string) (float64, error) {
	column, err := df.GetColumn(columnName)
	if err != nil {
		return 0, err
	}

	min := column[0].(float64)
	for _, val := range column {
		if val, ok := val.(float64); ok {
			if val < min {
				min = val
			}
		} else {
			return 0, errors.New("non-numeric value found in the column")
		}
	}

	return min, nil
}

// Max computes the maximum value for the specified numeric column.
func (df *DataFrame[T]) Max(columnName string) (float64, error) {
	column, err := df.GetColumn(columnName)
	if err != nil {
		return 0, err
	}

	max := column[0].(float64)
	for _, val := range column {
		if val, ok := val.(float64); ok {
			if val > max {
				max = val
			}
		} else {
			return 0, errors.New("non-numeric value found in the column")
		}
	}

	return max, nil
}

// Median computes the median value for the specified numeric column.
func (df *DataFrame[T]) Median(columnName string) (float64, error) {
	column, err := df.GetColumn(columnName)
	if err != nil {
		return 0, err
	}

	values := make([]float64, 0)
	for _, val := range column {
		if val, ok := val.(float64); ok {
			values = append(values, val)
		} else {
			return 0, errors.New("non-numeric value found in the column")
		}
	}

	sort.Float64s(values)
	n := len(values)
	if n%2 == 0 {
		return (values[n/2-1] + values[n/2]) / 2, nil
	}
	return values[n/2], nil
}

// Variance computes the variance for the specified numeric column.
func (df *DataFrame[T]) Variance(columnName string) (float64, error) {
	column, err := df.GetColumn(columnName)
	if err != nil {
		return 0, err
	}

	values := make([]float64, 0)
	for _, val := range column {
		if val, ok := val.(float64); ok {
			values = append(values, val)
		} else {
			return 0, errors.New("non-numeric value found in the column")
		}
	}

	n := float64(len(values))
	if n == 0 {
		return 0, errors.New("no values in column")
	}

	sum := 0.0
	for _, val := range values {
		sum += val
	}
	mean := sum / n

	varianceSum := 0.0
	for _, val := range values {
		varianceSum += (val - mean) * (val - mean)
	}
	variance := varianceSum / n

	return variance, nil
}

// Merge merges two dataframes based on common columns.
func (df *DataFrame[T]) Merge(otherDF *DataFrame[T]) (*DataFrame[T], error) {
	commonColumns := make([]string, 0)
	for _, col := range df.Headers {
		if otherDF.getColumnIndex(col) != -1 {
			commonColumns = append(commonColumns, col)
		}
	}

	if len(commonColumns) == 0 {
		return nil, errors.New("no common columns found for merge")
	}

	newHeaders := make([]string, 0)
	newHeaders = append(newHeaders, df.Headers...)
	for _, col := range otherDF.Headers {
		if otherDF.getColumnIndex(col) == -1 {
			newHeaders = append(newHeaders, col)
		}
	}

	newDF := NewDataFrame[T](newHeaders)

	for _, row1 := range df.Data {
		matchingRows := otherDF.Filter(func(row2 []T) bool {
			for _, col := range commonColumns {
				index1 := df.getColumnIndex(col)
				index2 := otherDF.getColumnIndex(col)
				if !reflect.DeepEqual(row1[index1], row2[index2]) {
					return false
				}
			}
			return true
		})

		for _, row2 := range matchingRows.Data {
			newRow := make([]T, len(newHeaders))
			copy(newRow, row1)
			for _, col := range otherDF.Headers {
				if otherDF.getColumnIndex(col) == -1 {
					continue
				}
				index := otherDF.getColumnIndex(col)
				newRow = append(newRow, row2[index])
			}
			newDF.AddRow(newRow...)
		}
	}

	return newDF, nil
}

// ConvertColumnType converts the data type of a column.
func (df *DataFrame[T]) ConvertColumnType(columnName string, newType reflect.Kind) error {
	index := df.getColumnIndex(columnName)
	if index == -1 {
		return errors.New("column not found")
	}

	newData := make([]T, len(df.Data))
	for i, row := range df.Data {
		val, err := convertValue(row[index], newType)
		if err != nil {
			return err
		}
		newData[i] = val.(T)
	}

	for i := range df.Data {
		df.Data[i][index] = newData[i]
	}

	return nil
}

// InferDataType infers the data types of columns based on the values.
func (df *DataFrame[T]) InferDataType() map[string]reflect.Kind {
	dataTypes := make(map[string]reflect.Kind)

	for _, col := range df.Headers {
		column, _ := df.GetColumn(col)
		dataType := inferDataType(column)
		dataTypes[col] = dataType
	}

	return dataTypes
}

// FillMissingValues fills missing values in the dataframe with the specified default value.
func (df *DataFrame[T]) FillMissingValues(defaultValue T) {
	for i, row := range df.Data {
		for j, val := range row {
			if val == nil {
				df.Data[i][j] = defaultValue
			}
		}
	}
}

// GroupBy groups the dataframe by a specific column and aggregates the data based on the aggregation function.
func (df *DataFrame[T]) GroupBy(groupByColumn string, aggregationFn func(data []T) T) (*DataFrame[T], error) {
	index := df.getColumnIndex(groupByColumn)
	if index == -1 {
		return nil, errors.New("group by column not found")
	}

	groups := make(map[T][]T)
	for _, row := range df.Data {
		groupValue := row[index]
		if _, ok := groups[groupValue]; !ok {
			groups[groupValue] = make([]T, 0)
		}
		groups[groupValue] = append(groups[groupValue], row)
	}

	groupedData := make([][]T, 0)
	for _, group := range groups {
		aggregatedRow := make([]T, len(df.Headers))
		for i, col := range df.Headers {
			if i == index {
				aggregatedRow[i] = group[0][i]
			} else {
				columnData := make([]T, len(group))
				for j, row := range group {
					columnData[j] = row[i]
				}
				aggregatedRow[i] = aggregationFn(columnData)
			}
		}
		groupedData = append(groupedData, aggregatedRow)
	}

	groupedHeaders := make([]string, len(df.Headers))
	copy(groupedHeaders, df.Headers)

	return &DataFrame[T]{
		Headers: groupedHeaders,
		Data:    groupedData,
	}, nil
}

// Pivot performs a pivot operation on the dataframe based on the specified column names.
func (df *DataFrame[T]) Pivot(rowColumn string, columnColumn string, valueColumn string) (*DataFrame[T], error) {
	rowIndex := df.getColumnIndex(rowColumn)
	columnIndex := df.getColumnIndex(columnColumn)
	valueIndex := df.getColumnIndex(valueColumn)

	if rowIndex == -1 || columnIndex == -1 || valueIndex == -1 {
		return nil, errors.New("row, column, or value column not found")
	}

	rowNames := make([]T, 0)
	columnNames := make([]T, 0)
	values := make([]T, 0)

	for _, row := range df.Data {
		rowName := row[rowIndex]
		columnName := row[columnIndex]
		value := row[valueIndex]

		if !contains(rowNames, rowName) {
			rowNames = append(rowNames, rowName)
		}

		if !contains(columnNames, columnName) {
			columnNames = append(columnNames, columnName)
		}

		values = append(values, value)
	}

	pivotedData := make([][]T, len(rowNames))
	for i, rowName := range rowNames {
		pivotedData[i] = make([]T, len(columnNames))
		for j, columnName := range columnNames {
			for k, value := range values {
				if df.Data[k][rowIndex] == rowName && df.Data[k][columnIndex] == columnName {
					pivotedData[i][j] = value.(T)
					break
				}
			}
		}
	}

	pivotedHeaders := make([]string, len(columnNames))
	copy(pivotedHeaders, columnNames)

	return &DataFrame[T]{
		Headers: pivotedHeaders,
		Data:    pivotedData,
	}, nil
}

// Helper function to check if a value is in a slice.
func contains(slice []T, value T) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}

// Helper function to convert a value to a specific data type.
func convertValue(value T, newType reflect.Kind) (T, error) {
	switch newType {
	case reflect.String:
		return fmt.Sprintf("%v", value), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		val, err := strconv.ParseInt(fmt.Sprintf("%v", value), 10, 64)
		if err != nil {
			return nil, err
		}
		return reflect.ValueOf(val).Convert(reflect.TypeOf(value)).Interface(), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		val, err := strconv.ParseUint(fmt.Sprintf("%v", value), 10, 64)
		if err != nil {
			return nil, err
		}
		return reflect.ValueOf(val).Convert(reflect.TypeOf(value)).Interface(), nil
	case reflect.Float32, reflect.Float64:
		val, err := strconv.ParseFloat(fmt.Sprintf("%v", value), 64)
		if err != nil {
			return nil, err
		}
		return reflect.ValueOf(val).Convert(reflect.TypeOf(value)).Interface(), nil
	case reflect.Bool:
		val, err := strconv.ParseBool(fmt.Sprintf("%v", value))
		if err != nil {
			return nil, err
		}
		return reflect.ValueOf(val).Convert(reflect.TypeOf(value)).Interface(), nil
	default:
		return nil, errors.New("unsupported data type conversion")
	}
}

// Helper function to infer the data type of a column.
func inferDataType(column []T) reflect.Kind {
	dataType := reflect.String

	for _, val := range column {
		switch val.(type) {
		case int, int8, int16, int32, int64:
			dataType = reflect.Int64
		case uint, uint8, uint16, uint32, uint64:
			dataType = reflect.Uint64
		case float32, float64:
			dataType = reflect.Float64
		case bool:
			dataType = reflect.Bool
		default:
			dataType = reflect.String
			return dataType
		}
	}

	return dataType
}
