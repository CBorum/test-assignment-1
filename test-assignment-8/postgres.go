package postgres

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type dbi interface {
	Query(string, ...interface{}) (*sql.Rows, error)
	Ping() error
	Close() error
}

type Track struct {
	trackid      int64
	name         string
	albumid      int64
	mediatypeid  int64
	genreid      int64
	composer     sql.NullString
	milliseconds int64
	bytes        int64
	unitprice    float64
}

type Album struct {
	albumid  int64
	title    string
	artistid int64
	composer sql.NullString
}

type Invoice struct {
	invoiceid   int64
	customerid  int64
	invoicedate time.Time
	total       float64
}

type Customer struct {
	customerid int64
	firstname  string
	lastname   string
	company    string
	email      string
}

func NewDB(dbName, dbUser string) (dbi, error) {
	dataSourceName := fmt.Sprintf("user=%s dbname=%s sslmode=disable", dbUser, dbName)
	return sql.Open("postgres", dataSourceName)
}

func Genre18And20(db dbi) (res []*Track, err error) {
	rows, err := db.Query(`
		SELECT * FROM chinook.track WHERE genreid = $1
		UNION
		SELECT * FROM chinook.track WHERE genreid = $2;
		`, 18, 20)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		t := &Track{}
		err = rows.Scan(&t.trackid, &t.name, &t.albumid, &t.mediatypeid, &t.genreid, &t.composer, &t.milliseconds, &t.bytes, &t.unitprice)
		if err != nil {
			return
		}
		res = append(res, t)
	}
	return
}

func Invoices(db dbi) (res []*Invoice, err error) {
	rows, err := db.Query(`
		SELECT invoiceid,customerid,invoicedate,total FROM chinook.invoice WHERE total < 10
		INTERSECT
		SELECT invoiceid,customerid,invoicedate,total FROM chinook.invoice WHERE total > 5;
		`)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		i := &Invoice{}
		err = rows.Scan(&i.invoiceid, &i.customerid, &i.invoicedate, &i.total)
		if err != nil {
			return
		}
		res = append(res, i)
	}
	return
}

func USACustomers(db dbi) (res []*Customer, err error) {
	rows, err := db.Query(`
		SELECT customerid,firstname,lastname,country,email FROM chinook.customer WHERE country = 'USA'
		EXCEPT
		SELECT customerid,firstname,lastname,country,email FROM chinook.customer WHERE email LIKE '%yahoo.com'
		`)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		c := &Customer{}
		err = rows.Scan(&c.customerid, &c.firstname, &c.lastname, &c.company, &c.email)
		if err != nil {
			return
		}
		res = append(res, c)
	}
	return
}

func MozartAndBach(db dbi) (res []*Album, err error) {
	rows, err := db.Query(`
		SELECT albumid, title, artistid, composer FROM chinook.album JOIN chinook.track USING(albumid) WHERE composer LIKE '%Mozart'
		UNION
		SELECT albumid, title, artistid, composer FROM chinook.album JOIN chinook.track USING(albumid) WHERE composer LIKE '%Bach';
	`)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		a := &Album{}
		err = rows.Scan(&a.albumid, &a.title, &a.artistid, &a.composer)
		if err != nil {
			return
		}
		res = append(res, a)
	}
	return
}
