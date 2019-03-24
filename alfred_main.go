package main

import (
	"bytes"
	"os"
	"strconv"

	"github.com/deanishe/awgo"
	"github.com/olekukonko/tablewriter"
)

var (
	wf *aw.Workflow

	colNum   int
	rowNum   int
	rowWidth int
)

func initDefaults() {
	colNum = 2
	rowNum = 2
	rowWidth = 6
}

func run() {
	argc := len(os.Args)
	if argc == 1 {
		return
	}
	if argc > 1 {
		val, err := strconv.Atoi(os.Args[1])
		if err != nil {
			colNum = val
		}
	}
	if argc > 2 {
		val, err := strconv.Atoi(os.Args[2])
		if err != nil {
			rowNum = val
		}
	}
	if argc > 3 {
		val, err := strconv.Atoi(os.Args[3])
		if err != nil {
			rowWidth = val
		}
	}
	wf.NewItem("Generate a 2x2 table")
	wf.NewItem("Generate a 2x2 table with 3-width placeholder")
	wf.SendFeedback()
	return
}

func buildTable() string {
	buf := new(bytes.Buffer)

	table := tablewriter.NewWriter(buf)
	placeholder := make([]string, colNum)
	table.SetHeader(placeholder)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	for idx := 0; idx < rowNum; idx++ {
		table.Append(placeholder)
	}
	table.Render()

	return buf.String()
}

func main() {
	// Init variable with default values
	initDefaults()

	wf = aw.New()
	wf.Run(run)
}
