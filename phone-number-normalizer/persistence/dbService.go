package persistence

import (
	"github.com/pkg/errors"
	"log"
)

const (
	startTransactionError = "Transaction creation error"
	createTableError = "Create table query error"
	insertError = "Insert query error"
	commitError = "Transaction commit error"
	rollbackError = "Transaction rollback error"
	selectError = "Select query error"
	scanError = "Scan record error"
)

func CreatePhonesTable() error{
	tx , err := DbInstance.Begin()
	if err != nil {
		return errors.New(startTransactionError)
	}
	_, err = tx.Exec("CREATE TABLE IF NOT EXISTS phones (id int auto_increment primary key, number text not null)")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return errors.New(rollbackError)
		}
		return errors.New(createTableError)
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return errors.New(rollbackError)
		}
		return errors.New(commitError)
	}
	return nil
}

func InsertSomePhones() error{
	tx , err := DbInstance.Begin()
	if err != nil {
		return errors.New(startTransactionError)
	}
	_, err = tx.Exec("INSERT INTO phones(number) VALUES ('1234567890'), ('123 456 7891'), ('(123) 456 7892'), ('(123) 456-7893'), ('123-456-7894'), ('123-456-7890'), ('1234567892'), ('(123)456-7892')")
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return errors.New(rollbackError)
		}
		return errors.New(insertError)
	}
	err = tx.Commit()
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return errors.New(rollbackError)
		}
		return errors.New(commitError)
	}
	return nil
}

func GetNumbers()[]string {
	rs, err := DbInstance.Query("SELECT * FROM phones")
	if err != nil {
		log.Println(selectError, err)
		return nil
	}

	var numbers []string
	for rs.Next() {
		var id int
		var nr string
		err = rs.Scan(&id, &nr)
		if err != nil {
			log.Println(scanError)
			continue
		}
		numbers = append(numbers, nr)
	}
	return numbers
}
