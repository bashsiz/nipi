package main

import (
	"strings"
	"fmt"
)

func (database *Database) cmdCheck(cmd string) {
	cmdSplit := strings.Split(cmd," ")
	cmdType := ""
	argOne := ""
    argTwo := ""
	for _, split := range cmdSplit {
        if split == "set" {
            cmdType = "set"
            continue
        }
        if split == "get" {
            cmdType = "get"
            continue
        }
        if split != "set" && cmdType == "set" && argOne == "" && argTwo == "" {
            argOne = split
            continue
        }
        if split != "set" && cmdType == "set" && argOne != "" && argTwo == "" {
            argTwo = split
            continue
        }
        if split != "get" && cmdType == "get" && argOne == "" {
            argOne = split
            continue
        }
    }
    if cmdType == "set" && argOne != "" && argTwo != "" {
        database.Set(argOne,argTwo)
    }
    if cmdType == "get" && argOne != "" {
		value,ok := database.Get(argOne)
		if ok {
			fmt.Println(value)	
		}
    }
}