package sqltogo

import (
	"bytes"
	"fmt"
	"github.com/liangyaopei/sqltogo/internal"
	"github.com/xwb1989/sqlparser"
)

// Parse converts a sql create statement to Go struct
// sqlStmt for sql create statement, pkgName for output directory, createFile for create file of Go struct
func Parse(sqlStmt string, pkgName string, createFile bool) (string, error) {
	statement, err := sqlparser.ParseStrictDDL(sqlStmt)
	if err != nil {
		return "", err
	}
	stmt, ok := statement.(*sqlparser.DDL)
	if !ok {
		return "", fmt.Errorf("input sql is not a create statment")
	}
	tableName := stmt.NewName.Name.String()
	lines, err := genGoStruct(stmt, tableName, pkgName)
	if err != nil {
		return "", err
	}
	if !createFile {
		return lines, nil
	}
	return lines, internal.GenGoFile(lines, tableName)
}

func genGoStruct(stmt *sqlparser.DDL, tableName string, pkgName string) (string, error) {
	var builder bytes.Buffer

	header := fmt.Sprintf("package %s\n", pkgName)
	headerPkg := "\nimport (\n" +
		"\t\"time\"\n" +
		")\n"
	importTime := false

	structName := internal.SnakeCaseToCamel(tableName)
	builder.WriteString(fmt.Sprintf("\n// %s%s", structName, stmt.TableSpec.Options))
	builder.WriteString(fmt.Sprintf("\ntype %s struct { \n", structName))

	var indexMap = make(map[string][]*sqlparser.IndexInfo)
	for _, index := range stmt.TableSpec.Indexes {
		for _, column := range index.Columns {
			if _, ok := indexMap[column.Column.String()]; !ok {
				indexMap[column.Column.String()] = []*sqlparser.IndexInfo{index.Info}
			} else {
				indexMap[column.Column.String()] = append(indexMap[column.Column.String()], index.Info)
			}
		}
	}

	for _, col := range stmt.TableSpec.Columns {
		colName, colType := col.Name.String(), col.Type

		columnType := colType.Type
		if colType.Unsigned {
			columnType += " unsigned"
		}

		goType := sqlTypeMap[columnType]
		if goType == "time.Time" {
			importTime = true
		}

		builder.WriteString(fmt.Sprintf("\t%s\t%s\t", internal.SnakeCaseToCamel(colName), goType))

		builder.WriteString(fmt.Sprintf("`gorm:\"column:%s;type:%s", colName, colType.Type))
		if colType.Length != nil {
			builder.WriteString(fmt.Sprintf("(%s)", internal.BytesToString(colType.Length.Val)))
		}
		if colType.NotNull {
			builder.WriteString(";not null")
		}
		if colType.Autoincrement {
			builder.WriteString(";autoIncrement")
		}
		if colType.Default != nil {
			var tmpl = ";default:%s"
			if colType.Default.Type == sqlparser.StrVal {
				tmpl = ";default:'%s'"
			}
			builder.WriteString(fmt.Sprintf(tmpl, internal.BytesToString(colType.Default.Val)))
		}

		if indexes, ok := indexMap[colName]; ok {
			for _, index := range indexes {
				indexName := "index"
				if index.Primary {
					builder.WriteString(";primaryKey")
					continue
				} else if index.Unique {
					indexName = "uniqueIndex"
				}
				builder.WriteString(fmt.Sprintf(";%s:%s", indexName, index.Name.String()))
			}
		}

		if colType.Comment != nil {
			builder.WriteString(fmt.Sprintf(";comment:'%s'", internal.BytesToString(colType.Comment.Val)))
		}
		builder.WriteString("\"`\n")
	}
	builder.WriteString("}\n")

	if importTime {
		return header + headerPkg + builder.String(), nil
	}
	return header + builder.String(), nil
}
