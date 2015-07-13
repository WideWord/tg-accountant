package accountant

import(
	"../db"
	"strconv"
	"math"
	"regexp"
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

func UserDebtors(user string) (debtors map[string]float64, err error) {
	keys, err := db.Client().Keys(user + ":debts:*").Result()
	if err != nil {
		return
	}

	debtors = make(map[string]float64)
	for _, key := range keys {
		debtStr, _err := db.Client().Get(key).Result()
		if _err != nil {
			err = _err
			return
		}

		debt, _err := strconv.ParseFloat(debtStr, 64)
		if _err == nil && math.Abs(debt) > 0.01 {
			re := regexp.MustCompile("[0-9]+:debts:(.*)")
			debtor := re.ReplaceAllString(key, "$1")
			debtors[debtor] = debt
		}
	}
	return
}
