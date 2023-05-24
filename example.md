Original DataFrame:

| Name  | Age | City     |
|-------|-----|----------|
| John  | 25  | New York |
| Alice | 30  | London   |
| Bob   | 28  | Paris    |
| Emily | 22  | Tokyo    |

DataFrame after column manipulation:

| Name  | Location | Gender |
|-------|----------|--------|
| John  | New York | Male   |
| Alice | London   | Female |
| Bob   | Paris    | Male   |
| Emily | Tokyo    | Female |

DataFrame after sorting:

| Name  | Location | Gender |
|-------|----------|--------|
| Alice | London   | Female |
| Bob   | Paris    | Male   |
| Emily | Tokyo    | Female |
| John  | New York | Male   |

Filtered DataFrame:

| Name  | Location | Gender |
|-------|----------|--------|
| Alice | London   | Female |

Descriptive Statistics:
Min: 22
Max: 30
Median: 28
Variance: 7

Merged DataFrame:

| Name  | Location | Gender | Salary |
|-------|----------|--------|--------|
| John  | New York | Male   | 5000   |
| Alice | London   | Female | 6000   |
| Bob   | Paris    | Male   | 5500   |

DataFrame after column data type conversion:

| Name  | Location | Gender | Salary |
|-------|----------|--------|--------|
| Alice | London   | Female | 6000   |
| Bob   | Paris    | Male   | 5500   |
| John  | New York | Male   | 5000   |

DataFrame after inferring column data types:

| Name  | Location | Gender | Salary |
|-------|----------|--------|--------|
| Alice | London   | Female | 6000   |
| Bob   | Paris    | Male   | 5500   |
| John  | New York | Male   | 5000   |

DataFrame after filling missing values:

| Name  | Location | Gender | Salary |
|-------|----------|--------|--------|
| Alice | London   | Female | 6000   |
| Bob   | Paris    | Male   | 5500   |
| John  | New York | Male   | 5000   |

Grouped DataFrame:

| Location | Salary |
|----------|--------|
| London   | 6000   |
| Paris    | 5500   |
| New York | 5000   |

Pivoted DataFrame:

|       | London | Paris | New York |
|-------|--------|-------|----------|
| Alice | 6000   |       |          |
| Bob   |        | 5500  |          |
| John  |        |       | 5000     |
