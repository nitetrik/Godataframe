**DataFrame Header:**
| Name  | Age | City      | Salary |
|-------|-----|-----------|--------|
| John  | 25  | New York  | 50000  |
| Alice | 30  | London    | 60000  |
| Bob   | 35  | Tokyo     | 75000  |
| Jane  | 28  | Paris     | 55000  |

**Count of Ages:** 4

**Sum of Salaries:** 240000

**Mean Age:** 29.5

**Filtered DataFrame:**
| Name | Age | City  | Salary |
|------|-----|-------|--------|
| Bob  | 35  | Tokyo | 75000  |

**Sorted DataFrame:**
| Name  | Age | City      | Salary |
|-------|-----|-----------|--------|
| John  | 25  | New York  | 50000  |
| Jane  | 28  | Paris     | 55000  |
| Alice | 30  | London    | 60000  |
| Bob   | 35  | Tokyo     | 75000  |

**Grouped DataFrame:**
| City     | Count |
|----------|-------|
| New York | 1     |
| London   | 1     |
| Tokyo    | 1     |
| Paris    | 1     |

**Joined DataFrame:**
| Name  | Age | City      | Population |
|-------|-----|-----------|------------|
| John  | 25  | New York  | 8500000    |
| Alice | 30  | London    | 9000000    |

**Serialized DataFrame (JSON):**
```json
[
	{"Name":"John","Age":25,"City":"New York","Salary":50000},
	{"Name":"Alice","Age":30,"City":"London","Salary":60000},
	{"Name":"Bob","Age":35,"City":"Tokyo","Salary":75000},
	{"Name":"Jane","Age":28,"City":"Paris","Salary":55000}
]

Name,Age,City,Salary
John,25,New York,50000
Alice,30,London,60000
Bob,35,Tokyo,75000
Jane,28,Paris,55000
