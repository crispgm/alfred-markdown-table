#!/bin/bash

go build
zip alfred-markdown-table.alfredworkflow ./alfred-markdown-table ./info.plist # ./icon.png
rm ./alfred-markdown-table
