# Alfred Markdown Table

Generate Markdown Table with Alfred Workflow.

![screenshot](/screenshot.png)

## Installation

### Install via Packal

[Alfred Markdown Table](http://www.packal.org/workflow/alfred-markdown-table)

### Install manually

Download and import [alfred-markdown-table.alfredworkflow](https://github.com/crispgm/alfred-markdown-table/raw/master/alfred-markdown-table.alfredworkflow).

## Dependencies

- [Alfred](https://www.alfredapp.com/)
- [Alfred Powerpack](https://www.alfredapp.com/powerpack/)

## Usage

Generate empty table:

```shell
table 3 3

Output:
|  |  |  |
|--|--|--|
|  |  |  |
|  |  |  |
|  |  |  |
```

Generate table with test data:

```shell
table 3 3 test
# `table 3 3 data` is an alternative

Output:
| COMPASSIONATE ALBATTANI | INSPIRING RITCHIE  |   UPBEAT BABBAGE    |
|-------------------------|--------------------|---------------------|
| determined_haibt        | goofy_darwin       | hungry_wilson       |
| elegant_rosalind        | suspicious_swirles | flamboyant_dijkstra |
| compassionate_blackwell | elastic_franklin   | silly_bardeen       |
```

## Credit

- [deanishe/awgo](https://github.com/deanishe/awgo)
- [olekukonko/tablewriter](https://github.com/olekukonko/tablewriter)

## License

MIT
