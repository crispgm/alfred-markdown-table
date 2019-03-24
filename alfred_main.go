package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/deanishe/awgo"
	"github.com/olekukonko/tablewriter"
)

var (
	wf *aw.Workflow

	colNum int
	rowNum int
)

func initDefaults() {
	colNum = 2
	rowNum = 2
}

func run() {
	argc := len(os.Args)
	if argc == 1 {
		return
	}
	if argc > 1 {
		params := os.Args[1]
		realArgs := strings.Split(params, " ")
		count := 0
		for _, arg := range realArgs {
			if arg == "" || arg == " " {
				continue
			}
			val, err := strconv.Atoi(arg)
			if err != nil {
				break
			}
			if count == 0 {
				colNum = val
			} else if count == 1 {
				rowNum = val
			} else {
				break
			}
			count++
		}
	}
	wf.NewItem(fmt.Sprintf("Generate a %dx%d table", colNum, rowNum)).Arg(buildTable()).Valid(true)
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
