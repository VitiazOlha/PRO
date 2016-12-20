package main

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"fmt"
	"net"
	"os"
	"strings"
)

var (
	gtm sync.Mutex
)

type TablesInMemory []*Table

type Table struct {
	name string
	data map[string]string
	m sync.RWMutex
}

func NewTable(file_name string, data map[string]string) *Table {
	return &Table{name: file_name, data: data,}
}

func DecodeJSON(file_name string) *Table{
	key_val, err := ioutil.ReadFile("data/" + file_name)
	if err != nil {
		return nil 
	}
	var f map[string]string
    err = json.Unmarshal(key_val, &f)
    if err != nil {
		return nil 
	}
    t := NewTable(file_name, f)
	return t
}

func EncodeJSON(tablechan <-chan Table) {
	for {
		table := <-tablechan
		jsonData, _ := json.Marshal(table.data)
		f, err := os.Create("data/" + table.name)
	    checkErr(err)
	    defer f.Close()
	    _, err = f.Write(jsonData)
	    checkErr(err)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getTable(tables *TablesInMemory, name string) *Table {
	gtm.Lock()
	for i := range *tables {
		if (*tables)[i].name == name {
			gtm.Unlock()
			return (*tables)[i]
		}
	}
	
	table := DecodeJSON(name)

	if table != nil {
		*tables = append(*tables, table)
	}
	gtm.Unlock()
	return table
}

func exit(c net.Conn) {
	c.Write([]byte(string("Bye\n")))
	c.Close()
}

func help(c net.Conn) {
	c.Write([]byte(string("[table name] get [key]\n")))
	c.Write([]byte(string("[table name] set [key] [value]\n")))
	c.Write([]byte(string("[table name] del [key]\n")))
	c.Write([]byte(string("[table name] keys\n")))
	c.Write([]byte(string("exit\n")))
}

func getKeys(c net.Conn, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 2  {
		table := getTable(tables, query_split[0])
		if (table == nil) {
			c.Write([]byte(string("Unknown table\n")))
		} else {
			table.m.RLock()
			keys := make([]string, 0, len(table.data))
		    for k := range table.data {
		        keys = append(keys, k)
		    }
		    table.m.RUnlock()
    		c.Write([]byte("[" + strings.Join(keys, ", ") + "]" + "\n"))
		}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func getVal(c net.Conn, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 3  {
		table := getTable(tables, query_split[0])
		if (table == nil) {
			c.Write([]byte(string("Unknown table\n")))
		} else {
			table.m.RLock()
			value, ok := table.data[query_split[2]]
			table.m.RUnlock()
			if ok {
				c.Write([]byte(string(value + "\n")))
			} else {
				c.Write([]byte(string("key does not exist\n")))
			}
		}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func setVal(c net.Conn, tablechan chan<- Table, tables *TablesInMemory, query_split []string) {
	if len(query_split) >= 4 {
		table := getTable(tables, query_split[0])
		if (table == nil) {
			table = NewTable(query_split[0], make(map[string]string))
		} 
		table.m.Lock()
		table.data[query_split[2]] = strings.Join(query_split[3: ], " ")
		table.m.Unlock()
		c.Write([]byte(string("OK\n")))
		tablechan <- *table
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func delKey(c net.Conn, tablechan chan<- Table, tables *TablesInMemory, query_split []string) {
	if len(query_split) == 3  {
		table := getTable(tables, query_split[0])
		if (table == nil) {
			c.Write([]byte(string("Unknown table\n")))
		} else {
			table.m.RLock()
			_, ok := table.data[query_split[2]]
			table.m.RUnlock()
			if ok {
				table.m.Lock()
				delete(table.data, query_split[2])
				table.m.Unlock()
				c.Write([]byte(string("OK\n")))
				tablechan <- *table
			} else {
				c.Write([]byte(string("key does not exist\n")))
			}
		}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func handleRequest(c net.Conn, tablechan chan<- Table, tables *TablesInMemory, query string) {
	query_split := strings.Fields(query)

	if len(query_split) >= 2 {
		switch strings.ToLower(query_split[1]) {
			case "set":
				setVal(c, tablechan, tables, query_split)
			case "get":
				getVal(c, tables, query_split)
			case "del":
				delKey(c, tablechan, tables, query_split)
			case "keys":
				getKeys(c, tables, query_split)
			default:
				c.Write([]byte(string("Unknown command\n")))
		}
	} else if len(query_split) == 1 {
		switch strings.ToLower(query_split[0]) {
			case "exit":
				exit(c)
			case "help":
				help(c)
			default:
				c.Write([]byte(string("Unknown command\n")))
			}
	} else {
		c.Write([]byte(string("Unknown command\n")))
	}
}

func handleConnection(c net.Conn, tablechan chan<- Table, tables *TablesInMemory) {
	buf := make([]byte, 4096)
	for {
		n, err := c.Read(buf)
		if (err != nil) || (n == 0) {
			break
		} else {
			go handleRequest(c, tablechan, tables, string(buf[0:n]))
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":2222")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer ln.Close()
	var tables TablesInMemory

	tablechan := make(chan Table)
	go EncodeJSON(tablechan)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()
		go handleConnection(conn, tablechan, &tables)
	}
}

