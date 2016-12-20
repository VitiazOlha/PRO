package main

import (
	"testing"
	"strings"
)

func getTab(tables *TablesInMemory, name string) *Table {
	for i := range *tables {
		if (*tables)[i].name == name {
			return (*tables)[i]
		}
	}
	table := DecodeJSON(name)

	if table != nil {
		*tables = append(*tables, table)
	}
	return table
}

func set(tables *TablesInMemory, query_split []string) string {
	if len(query_split) >= 4 {
		table := getTab(tables, query_split[0])
		if (table == nil) {
			table = NewTable(query_split[0], make(map[string]string))
		} 
		
		table.data[query_split[2]] = strings.Join(query_split[3: ], " ")
		return "OK"
	} else {
		return "Error"
	}
}

func TestSet(t *testing.T) {
	var tables TablesInMemory
	query := "name1 set 11 28" 
	query_split := strings.Fields(query)
	if set(&tables, query_split) == "OK"{
		t.Log("Set operation in existing file successful.")
	} else {
		t.Error("FAILED: Set operation in existing file.")
	}

	query1 := "name1012 set user boris" 
	query_split1 := strings.Fields(query1)
	if set(&tables, query_split1) == "OK"{
		t.Log("Set operation with creating of new file passed successful.")
	} else {
		t.Error("FAILED: Set operation with creating of new file.")
	}
	
}

// func TestGet(t *testing.T) {

// }

// func TestDel(t *testing.T) {

// }

// func TestKeys(t *testing.T) {
	
// }
