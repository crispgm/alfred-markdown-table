package main

import "testing"

func TestBuildTable(t *testing.T) {
	tableWithoutTestData := buildTable(2, 2, false)
	t.Log(tableWithoutTestData)
	tableWithTestData := buildTable(2, 2, true)
	t.Log(tableWithTestData)
}
