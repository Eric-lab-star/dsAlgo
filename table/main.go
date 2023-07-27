package main

import "fmt"

type Table struct {
	Name        string
	ColumnNames []string
	Rows        []Row
}

type Row struct {
	id     string
	column []Column
}

type Column struct {
	id     string
	values string
}

func printTable(table Table) {
	fmt.Println(table.Name)
	rows := table.Rows
	for _, row := range rows {
		for i, column := range row.column {
			fmt.Println(table.ColumnNames[i], column.id, column.values)
		}
	}
}

func main() {
	table := Table{}
	table.Name = "customer"
	table.ColumnNames = []string{"Id", "Name", "PhoneNumner"}
	table.Rows = make([]Row, 2)

	row1 := Row{}
	row1.column = make([]Column, 3)
	row1.column[0] = Column{"1", "01"}
	row1.column[1] = Column{"1", "Kyungsub Kim"}
	row1.column[2] = Column{"1", "01062888587"}

	row2 := Row{}
	row2.column = make([]Column, 3)
	row2.column[0] = Column{"2", "02"}
	row2.column[1] = Column{"2", "James ho"}
	row2.column[2] = Column{"2", "1032220"}

	table.Rows[0] = row1
	table.Rows[1] = row2
	printTable(table)

}

