/**
		GDG GoLang Backend Daemon
        Copyright (C) 2014+  Gabriele Baldoni
 */

package dbhelper


//TODO far diventare dbhelper un oggetto

//nome package

import ( //package importati
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"utils"
	"strconv"
)

type STMTResult struct{
	ID           int64
	RowsAffected int64
}

func countRowsSelected(rows *sql.Rows, c chan int) {
	i := 0;
	for rows.Next() {
		i++
	}

	c<-i;
}

func connect() (*sql.DB, error) {
	conf := utils.LoadConfiguration()
	
	utils.Log(utils.DEBUG, "GDGBackend Service", "DB Readed Configuration " + conf.Address + " " +conf.Port)
	utils.Log(utils.DEBUG, "GDGBackend Service", "DB Readed Configuration " + conf.DBHost + " " + conf.DBPort )
	utils.Log(utils.DEBUG, "GDGBackend Service", "DB Readed Configuration " + conf.DBName + " " +conf.DBUser)


	if (conf.Debug){
		utils.Log(utils.DEBUG, "GDGBackend DBHelper Debug", "DB Connection String: " + conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)
		db, err := sql.Open("mysql", conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)
		if (err != nil) {
			fmt.Println(err.Error())
			utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
			return nil,err
		}
		return db, err
		
	} else {
		utils.Log(utils.DEBUG, "GDGBackend DBHelper PROD", "DB Connection String: " + conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)
		db, err := sql.Open("mysql", conf.DBUser+":"+conf.DBPassword+"@tcp("+conf.DBHost+":"+conf.DBPort+")/"+conf.DBName)
		if (err != nil) {
			fmt.Println(err.Error())
			utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
			return nil,err
		}
		return db, err
	}
	

	
	

	
}

func disconnect(database *sql.DB) {
	database.Close()
}



func RawQuery(query string) string {


	utils.Log(0, "dbhelper", "executing raw query: "+query)
	barr := make([]map[string]interface{}, 0, 0)

	db, err := connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer disconnect(db)


	rows, err := db.Query(query)
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}

	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))

	for i := range values {
		scanArgs[i] = &values[i]
	}
	
	i := 0
	count:=0
	for rows.Next() {
		err = rows.Scan(scanArgs...)

		if (err != nil) {
			fmt.Println(err.Error())
			utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
			return strconv.Quote("false")
		}

		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				record[columns[i]] = fmt.Sprintf("%s", string(col.([]byte)))
			}
		}


		barr = append(barr[i:], record)
		count++
	}


	if count == 0 {
		return strconv.Quote("null")
	}

	s, _ := json.Marshal(barr)
	return string(s)
}

// TODO correggere metodo Select
func Select(table string,whereArgs string,columns ...string) string {
	utils.Log(utils.ASSERT, "dbhelper", "executing Select on " + table)
	barr := make([]map[string]interface{}, 0, 0)

	db, err := connect()
	if (err != nil) {
		fmt.Println(err.Error())
		return strconv.Quote("false")
	}
	defer disconnect(db)


	var selColoumns string
	selColoumns=""

	for i := range columns {

		selColoumns+=columns[i]
		if (i != len(columns)-1) {
			selColoumns+=","
		}

	}

	var query string


	if(whereArgs!=""){
		query="Select " + selColoumns + " FROM "+table+" WHERE "+ whereArgs
	} else {
		query="Select " + selColoumns + " FROM "+table
	}


	stmt, err := db.Prepare(query)

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()


	rows, err := stmt.Query(query)
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer rows.Close()

	qcolumns, err := rows.Columns()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}

	scanArgs := make([]interface{}, len(qcolumns))
	values := make([]interface{}, len(qcolumns))

	for i := range values {
		scanArgs[i] = &values[i]
	}


	i := 0
	count:=0
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if (err != nil) {
			fmt.Println(err.Error())
			utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
			return strconv.Quote("false")
		}

		record := make(map[string]interface{})

		for i, col := range values {
			if col != nil {
				record[columns[i]] = fmt.Sprintf("%s", string(col.([]byte)))
			}
		}


		barr = append(barr[i:], record)
		count++
	}



	if count == 0 {
		return strconv.Quote("null")
	}

	s, _ := json.Marshal(barr)
	return string(s)
}




func Insert(table string , columns []string, values []string) string {

	db, err := connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer disconnect(db)

	insColoumn := ""
	insValues := ""

	for i := range columns {

		insColoumn+=columns[i]
		if (i != len(columns)-1) {
			insColoumn+=","
		}

	}


	for i := range values {

		insValues+=strconv.Quote(values[i])
		if (i != len(values)-1) {
			insValues+=","
		}

	}

	stmt, err := db.Prepare("INSERT INTO " + table + " (" + insColoumn + ") VALUES ( " + insValues + " )")

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}

	i, _ := res.LastInsertId()
	n, _ := res.RowsAffected()
	id := strconv.FormatInt(i, 10)
	num := strconv.FormatInt(n, 10)

	utils.Log(2, "dbhelper", "insert into "+table+" rowsaffected:"+num+" last id: "+id)

	g := STMTResult{ID:i, RowsAffected:n}


	b, _ := json.Marshal(g)
	return string(b)

}


func Update(table string , columns []string, values []string,whereArgs string) string {

	db, err := connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer disconnect(db)

	if len(values)!=len(columns){
		return strconv.Quote("false")
	}

	updateValues:=""

	for i := range columns {

		updateValues=updateValues+columns[i]+"="+strconv.Quote(values[i])
		if (i != len(columns)-1) {
			updateValues+=","
		}

	}




	stmt, err := db.Prepare("UPDATE " + table + " SET  " +updateValues+ "  WHERE "+ whereArgs)

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}

	i, _ := res.LastInsertId()
	n, _ := res.RowsAffected()
	id := strconv.FormatInt(i, 10)
	num := strconv.FormatInt(n, 10)

	utils.Log(2, "dbhelper", "update table "+table+" rowsaffected:"+num+" last id: "+id)

	g := STMTResult{ID:i, RowsAffected:n}


	b, _ := json.Marshal(g)
	return string(b)

}


func Delete(table string , whereArgs string) string {

	db, err := connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer disconnect(db)



	stmt, err := db.Prepare("DELETE FROM " + table +  " WHERE "+ whereArgs)

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "GDGBackend DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	i, _ := res.LastInsertId()
	n, _ := res.RowsAffected()
	id := strconv.FormatInt(i, 10)
	num := strconv.FormatInt(n, 10)

	utils.Log(2, "dbhelper", "delete table "+table+" rowsaffected:"+num+" last id: "+id)

	g := STMTResult{ID:i, RowsAffected:n}


	b, _ := json.Marshal(g)
	return string(b)

}



