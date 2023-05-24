# DataFrame Module Documentation

The DataFrame module provides functionality for working with two-dimensional tabular data in Go. It includes operations for creating, manipulating, and analyzing dataframes.

## main.go

The `main.go` file contains an example usage of the DataFrame module. It demonstrates various operations on a dataframe, such as column manipulation, sorting, filtering, descriptive statistics, merging, data type conversion, missing value handling, grouping, and pivoting.

To run the example code, follow these steps:

1. Ensure you have Go installed on your system.
2. Install the required dependencies by running `go get -u github.com/nitetrik/Godataframe/dataframe`.
3. Replace `"github.com/nitetrik/Godataframe/dataframe"` with the correct import path to your dataframe module in the `main.go` file.
4. Run the command `go run main.go` in the terminal.

The output will show the results of the dataframe operations performed in the example code.

## dataframe/dataframe.go

The `dataframe.go` file contains the implementation of the DataFrame module. It defines a `DataFrame[T]` struct and provides methods for creating, manipulating, and analyzing dataframes.

### DataFrame[T] Struct

The `DataFrame[T]` struct represents a two-dimensional table of data. It has the following fields:

- `Headers`: A slice of strings that contains the column names of the dataframe.
- `Data`: A slice of slices of type `T` that stores the data values of the dataframe.

### Functions and Methods

- `NewDataFrame[T](columns []string) *DataFrame[T]`: Creates a new DataFrame with the given column names.
- `AddRow(rowData ...T) error`: Adds a new row of data to the DataFrame.
- `Print()`: Displays the DataFrame in a tabular format.
- `GetColumn(columnName string) ([]T, error)`: Returns the values in a specific column of the DataFrame.
- `SumColumn(columnName string) (float64, error)`: Calculates the sum of the values in a specific numeric column of the DataFrame.
- `RenameColumn(oldName string, newName string) error`: Renames a column in the DataFrame.
- `AddColumn(columnName string, values []T) error`: Adds a new column to the DataFrame.
- `RemoveColumn(columnName string) error`: Removes a column from the DataFrame.
- `ReorderColumns(columnOrder []string) error`: Reorders the columns based on the specified order.
- `Sort(columns []string) error`: Sorts the dataframe based on one or more columns.
- `Filter(filter func(row []T) bool) *DataFrame[T]`: Returns a new DataFrame containing only the rows that match the filter criteria.
- `MeanColumn(columnName string) (float64, error)`: Returns the mean of the values in a specific numeric column of the DataFrame.
- `Min(columnName string) (float64, error)`: Computes the minimum value for the specified numeric column.
- `Max(columnName string) (float64, error)`: Computes the maximum value for the specified numeric column.
- `Median(columnName string) (float64, error)`: Computes the median value for the specified numeric column.
- `Variance(columnName string) (float64, error)`: Computes the variance for the specified numeric column.
- `Merge(otherDF *DataFrame[T]) (*DataFrame[T], error)`: Merges two dataframes based on common columns.
- `ConvertColumnType(columnName string, newType reflect.Kind) error`: Converts the data type of a column.
- `InferColumnTypes()`: Infers the data types of columns based on the values.
- `FillMissingValues(defaultValue T)`: Fills missing values in the dataframe with the specified default value.
- `GroupBy(groupByColumn string, aggregationFn func(data []T) T) (*DataFrame[T], error)`: Groups the dataframe by a specific column and aggregates the data based on the aggregation function.
- `Pivot(rowColumn string, columnColumn string, valueColumn string) (*DataFrame[T], error)`: Performs a pivot operation on the dataframe based on the specified columns.

---

This is just an overview of the DataFrame module and its usage. For detailed documentation and additional examples, refer to the comments and code in the `dataframe.go` file.

###Examples:
## Original DataFrame:
| Name | Age | City     |
|------|-----|----------|
| John | 30  | New York |
| Jane | 25  | London   |
| Mike | 35  | Paris    |

## Ages:
[30 25 35]

## Sum of ages:
90

## Filtered DataFrame:
| Name | Age | City     |
|------|-----|----------|
| John | 30  | New York |
| Mike | 35  | Paris    |

## Renamed DataFrame:
| Name | Age | Location |
|------|-----|----------|
| John | 30  | New York |
| Jane | 25  | London   |
| Mike | 35  | Paris    |

## DataFrame with new column:
| Name | Age | Location | Population (in millions) |
|------|-----|----------|-------------------------|
| John | 30  | New York | 8.5                     |
| Jane | 25  | London   | 9                       |
| Mike | 35  | Paris    | 10.2                    |

## DataFrame with column removed:
| Age | Location | Population (in millions) |
|-----|----------|-------------------------|
| 30  | New York | 8.5                     |
| 25  | London   | 9                       |
| 35  | Paris    | 10.2                    |

## DataFrame with reordered columns:
| Location | Population (in millions) | Age |
|----------|-------------------------|-----|
| New York | 8.5                     | 30  |
| London   | 9                       | 25  |
| Paris    | 10.2                    | 35  |

## DataFrame sorted by Age:
| Location | Population (in millions) | Age |
|----------|-------------------------|-----|
| London   | 9                       | 25  |
| New York | 8.5                     | 30  |
| Paris    | 10.2                    | 35  |

## Minimum age:
25

## Maximum age:
35

## Median age:
30

## Variance of age:
16.333333333333332

## Merged DataFrame:
| Location | Population (in millions) | Age | Name |
|----------|-------------------------|-----|------|
| New York | 8.5                     | 30  | John |
| London   | 9                       | 25  | Jane |
| Paris    | 10.2                    | 35  | Mike |

## DataFrame with converted column type:
| Location | Population (in millions) | Age |
|----------|-------------------------|-----|
| New York | 8.5                     | 30  |
| London   | 9                       | 25  |
| Paris    | 10.2                    | 35  |

## Inferred data types:
Location: string
Population (in millions): float64
Age: float64

## DataFrame with filled missing values:
| Location | Population (in millions) | Age |
|----------|-------------------------|-----|
| New York | 8.5                     | 30  |
| London   | 9                       | 25  |
| Paris    | 10.2                    | 35  |

## Grouped DataFrame:
| Location | Count |
|----------|-------|
| New York | 1     |
| London   | 1     |
| Paris    | 1     |

## Pivoted DataFrame:
| Age | New York | London | Paris |
|-----|----------|--------|-------|
| 30  | 8.5      | 0      | 0     |
| 25  | 0        | 9      | 0     |
| 35  | 0        | 0      | 10.2  |
