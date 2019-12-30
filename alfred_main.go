package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crispgm/alfred-markdown-table/pkg/namesgenerator"
	aw "github.com/deanishe/awgo"
	"github.com/olekukonko/tablewriter"
)

var (
	wf *aw.Workflow
)

func run() {
	var (
		colNum   int = 2
		rowNum   int = 2
		testData bool
	)
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
			if err == nil {
				if count == 0 {
					colNum = val
				} else if count == 1 {
					rowNum = val
				}
			} else if count == 2 && (arg == "data" || arg == "test") {
				testData = true
			} else {
				break
			}
			count++
		}
	}
	wf.NewItem(getSubtitle(colNum, rowNum, testData)).
		Arg(buildTable(colNum, rowNum, testData)).
		Valid(true)
	wf.SendFeedback()
	return
}

func getSubtitle(colNum, rowNum int, testData bool) string {
	subtitle := fmt.Sprintf("Generate a %dx%d table", colNum, rowNum)
	if testData {
		subtitle += " with test data"
	}
	return subtitle
}

func buildTable(colNum, rowNum int, testData bool) string {
	buf := new(bytes.Buffer)

	table := tablewriter.NewWriter(buf)

	table.SetHeader(generateRow(colNum, testData))
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	for idx := 0; idx < rowNum; idx++ {
		table.Append(generateRow(colNum, testData))
	}
	table.Render()

	return buf.String()
}

func generateRow(colNum int, testData bool) []string {
	row := make([]string, colNum)
	for index := range row {
		if testData {
			row[index] = namesgenerator.GetRandomName(0)
		} else {
			row[index] = " "
		}
	}
	return row
}

func main() {
	wf = aw.New()
	wf.Run(run)
}
