package main

type Column struct {
	Name string
	Type string
}

type LoadReq struct {
	TableName   string            `json:"table_name"`
	ColumnMap   map[string]Column `json:"column_map"`   //csv 列名对应table列
	FileColumns []string          `json:"file_columns"` //csv 按序列名
	FilePath    string            `json:"file_path"`
	Format      string            `json:"format"`
	Delimiter   string            `json:"delimiter"`
	NullValue   string            `json:"null_value"`
	Escape      string            `json:"escape"`
	Newline     string            `json:"newline"`
	Quote       string            `json:"quote"`
	Header      bool              `json:"header"`
	MissingFill bool              `json:"missing_fill"`
}

TempStorage := "/tmp/holmes"

func main() {

}

// Get UUID
func GetUUID() string {
	guid := uuid.Must(uuid.NewV4())
	return strings.Replace(guid.String(), "-", "", -1)
}

// 1. create final table & tmp table
// 2. copy csv to tmp table
// 3. defer drop tmp table
// 4. insert into final table
// 5. if err drop final table

func load(req *LoadReq) error {
	tmpTableName := fmt.Sprintf("%s_%s", req.TableName, GetUUID())
	colList := []string{}
	for _, colCsv := range req.FileColumns {
		if col, ok := req.ColumnMap[colCsv]; ok {
			colList = append(colList, fmt.Sprintf(`%s %s`, colCsv, col.Type))
		} else {
			colList = append(colList, fmt.Sprintf(`%s text`, colCsv))
		}
	}

	tmpCreateSql := fmt.Sprintf(`CREATE TABLE %s(%s);`, tmpTableName, strings.Join(colList, ","))
	fmt.Println("tmpCreateSql: ", tmpCreateSql)

	finalColList := []string{}
	for _, v := range req.FileColumns {
		finalColList = append(finalColList, fmt.Sprintf(`%s %s`, v.Name, v.Type))
	}

	finalCreateSql := fmt.Sprintf(`CREATE TABLE %s(%s)`, req.TableName, strings.Join(finalColList, ","))
	fmt.Println("finalCreateSql: ", finalCreateSql)

	copySql := fmt.Sprintf(
		`COPY %s FROM '%s' `,
		tmpTableName,
		path.Join(TempStorage, req.FilePath))
	if req.Delimiter != "" {
		copySql = copySql + " delimiter '" + req.Delimiter + "'"
	}
	if req.NullValue != "" {
		copySql = copySql + " null '" + req.NullValue + "'"
	}
	if req.Escape != "" {
		copySql = copySql + " escape '" + req.Escape + "'"
	}
	if req.Newline != "" {
		copySql = copySql + " newline '" + req.Newline + "'"
	}
	if req.Format == "csv" {
		copySql = copySql + " csv"
	}
	if req.Quote != "" {
		copySql = copySql + " quote '" + req.Quote + "'"
	}
	if req.Header == true {
		copySql = copySql + " header"
	}
	if req.MissingFill == true {
		copySql = copySql + " fill missing fields"
	}
	lavalog.Debug.Println("copySql is : ", copySql)
	res, err := requestHandler.Transaction.Exec(copySql)
	if err != nil {
		lavalog.Error.Println(fmt.Sprintf("Failed to copy data from file=%s to table=%s :", table, req.FilePath), err)
		return -1, err
	}
	rowsAffected, _ := res.RowsAffected()
	lavalog.Debug.Println("res.La : ", rowsAffected)

	return rowsAffected, nil
}
