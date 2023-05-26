package dataframe

import (
	"math"
)

// Function to transform data, such as scaling, normalization, encoding categorical variables, etc.
func (df *DataFrame) TransformData() {
	// Loop through each column in the DataFrame
	for _, columnName := range df.header {
		columnData := df.columns[columnName]

		// Perform data transformation based on column type
		switch columnData[0].(type) {
		case int, int64, float64:
			// Perform scaling or normalization on numeric columns
			minVal, maxVal := df.getMinMaxValues(columnData)
			df.scaleColumn(columnName, minVal, maxVal)
		case string:
			// Perform encoding on categorical columns
			df.encodeColumn(columnName)
		}
	}
}

// Function to calculate the minimum and maximum values in a numeric column
func (df *DataFrame) getMinMaxValues(column []interface{}) (float64, float64) {
	minVal := math.Inf(1)
	maxVal := math.Inf(-1)

	for _, value := range column {
		if val, ok := value.(float64); ok {
			if val < minVal {
				minVal = val
			}
			if val > maxVal {
				maxVal = val
			}
		}
	}

	return minVal, maxVal
}

// Function to scale a numeric column to a range of [0, 1]
func (df *DataFrame) scaleColumn(columnName string, minVal, maxVal float64) {
	column := df.columns[columnName]

	for i := 0; i < len(column); i++ {
		if val, ok := column[i].(float64); ok {
			scaledVal := (val - minVal) / (maxVal - minVal)
			df.columns[columnName][i] = scaledVal
		}
	}
}

// Function to encode categorical column using one-hot encoding
func (df *DataFrame) encodeColumn(columnName string) {
	column := df.columns[columnName]
	uniqueValues := make(map[interface{}]bool)

	// Get unique values in the column
	for _, value := range column {
		uniqueValues[value] = true
	}

	// Create new columns for each unique value
	for value := range uniqueValues {
		newColumnName := columnName + "_" + value.(string)
		encodedValues := make([]interface{}, len(column))

		// Encode the column values based on unique value presence
		for i, val := range column {
			if val == value {
				encodedValues[i] = 1
			} else {
				encodedValues[i] = 0
			}
		}

		// Add the new encoded column to the DataFrame
		df.AddColumn(newColumnName, encodedValues)
	}

	// Remove the original categorical column from the DataFrame
	df.RemoveColumn(columnName)
}
