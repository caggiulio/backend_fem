/**
		Progetto Architetture 
        Copyright (C) 2015+  Gabriele Baldoni
 */

package dbhelper

/* oggetto DBHelper
*	Si occuppa della connessione al DB e della esecuzione delle query
*/
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"utils"
	"strconv"
)





type DBHelper struct {
	Config utils.Configuration
	db *sql.DB
}



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


func (dbh *DBHelper)LoadConfiguration() {
	dbh.Config=utils.LoadConfiguration()
}

func (dbh *DBHelper)SetConfiguration(mConf utils.Configuration) {
	dbh.Config=mConf
}





func (dbh *DBHelper) connect() ( error) {
	//conf := utils.LoadConfiguration()

	if (dbh.Config.Debug){
		utils.Log(utils.DEBUG, "ProgettoFEM DBHelper Debug", "DB Connection String: " + dbh.Config.DBUser+":"+dbh.Config.DBPassword+"@tcp("+dbh.Config.DBHost+":"+dbh.Config.DBPort+")/"+dbh.Config.DBName)
		db, err := sql.Open("mysql", dbh.Config.DBUser+":"+dbh.Config.DBPassword+"@tcp("+dbh.Config.DBHost+":"+dbh.Config.DBPort+")/"+dbh.Config.DBName)
		if (err != nil) {
			fmt.Println(err.Error())
			utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
			return err
		}
		dbh.db=db
		return  err

	} else {
		utils.Log(utils.DEBUG, "ProgettoFEM DBHelper PROD", "DB Connection String: " + dbh.Config.DBUser+":"+dbh.Config.DBPassword+"@tcp("+dbh.Config.DBHost+":"+dbh.Config.DBPort+")/"+dbh.Config.DBName)
		db, err := sql.Open("mysql", dbh.Config.DBUser+":"+dbh.Config.DBPassword+"@tcp("+dbh.Config.DBHost+":"+dbh.Config.DBPort+")/"+dbh.Config.DBName)
		if (err != nil) {
			fmt.Println(err.Error())
			utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
			return err
		}
		dbh.db=db
		return  err
	}






}

func (dbh *DBHelper)disconnect() {
	dbh.db.Close()
	dbh.db=nil
}



func (dbh DBHelper)RawQuery(query string) string {


	utils.Log(0, "dbhelper", "executing raw query: "+query)
	barr := make([]map[string]interface{}, 0, 0)

	err := dbh.connect()
	if dbh.db==nil {
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "DB is nill!!!")
		return strconv.Quote("false")
	}

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer dbh.disconnect()


	rows, err := dbh.db.Query(query)
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
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
			utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
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
/*func Select(table string,whereArgs string,columns ...string) string {
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
}*/




func (dbh DBHelper)Insert(table string , columns []string, values []string) string {

	err := dbh.connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer dbh.disconnect()

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

	stmt, err := dbh.db.Prepare("INSERT INTO " + table + " (" + insColoumn + ") VALUES ( " + insValues + " )")

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
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


func (dbh DBHelper)Update(table string , columns []string, values []string,whereArgs string) string {

	err := dbh.connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer dbh.disconnect()

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




	stmt, err := dbh.db.Prepare("UPDATE " + table + " SET  " +updateValues+ "  WHERE "+ whereArgs)

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
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


func (dbh DBHelper)Delete(table string , whereArgs string) string {

	err := dbh.connect()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer dbh.disconnect()



	stmt, err := dbh.db.Prepare("DELETE FROM " + table +  " WHERE "+ whereArgs)

	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
		return strconv.Quote("false")
	}
	defer stmt.Close()

	res, err := stmt.Exec()
	if (err != nil) {
		fmt.Println(err.Error())
		utils.Log(utils.ERROR, "ProgettoFEM DBHelper", "Error: " + err.Error())
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

func (dbh DBHelper)CountRowsSelected(rows *sql.Rows, c chan int) {
	i := 0;
	for rows.Next() {
		i++
	}

	c<-i;
}

