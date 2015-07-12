package accountant

import(
	"../db"
	"strconv"
)

func AddDebt(user string, debtor string, change float64) (newDebt float64, err error) {
	res := db.Client().IncrByFloat(user + ":debts:" + debtor, change)
	newDebt = res.Val()
	err = res.Err()
	return
}

func SetDebt(user string, debtor string, newDebt float64) error {
	err := db.Client().Set(user + ":debts:" + debtor, strconv.FormatFloat(newDebt, 'f', 0, 64), 0).Err()
	return err
}

