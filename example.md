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
