package internal

import (
	"errors"
	"fmt"
	"github.com/dariubs/percent"
)

type Request struct {
	ProductName string  `json:"product_name"`
	Amount      float64 `json:"amount"`
	Phone       string  `json:"phone"`
	Period      int     `json:"period"`
}

type Response struct {
	ProductName          string  `json:"product_name"`
	Amount               float64 `json:"amount"`
	Phone                string  `json:"phone"`
	Period               int     `json:"period"`
	Commission           int     `json:"commission"`
	AmountWithCommission float64 `json:"amount_with_commission"`
}

const (
	phone    = "phone"
	computer = "computer"
	tv       = "tv"
)

func InstallmentPayments(req Request) (Response, error) {
	if !(req.ProductName == phone || req.ProductName == computer || req.ProductName == tv) {
		return Response{}, fmt.Errorf("[ERROR] Invalid product name: " + req.ProductName)
	}

	cal, commission, err := calculatePercentage(req.ProductName, req.Period, req.Amount)
	if err != nil {
		return Response{}, err
	}

	resp := Response{
		ProductName:          req.ProductName,
		Amount:               req.Amount,
		Phone:                req.Phone,
		Period:               req.Period,
		Commission:           commission,
		AmountWithCommission: cal,
	}

	err = sendMessage(resp)
	if err != nil {
		return Response{}, err
	}

	return resp, nil
}

func sendMessage(res Response) error {
	return nil
}

func calculatePercentage(proName string, period int, amount float64) (res float64, commission int, err error) {

	fee := map[string]int{phone: 3, computer: 4, tv: 5}

	switch proName {
	case phone:
		commission = fee[phone]
		switch period {
		case 3, 6, 9:
			res = amount
			return res, commission, err
		case 12:
			res = amount + percent.Percent(commission, int(amount))
			return res, commission, err
		case 18:
			res = amount + percent.Percent(commission*2, int(amount))
			return res, commission, err
		default:
			err = errors.New("invalid period")
			return res, commission, err
		}
	case computer:
		commission = fee[computer]
		switch period {
		case 3, 6, 9, 12:
			res = amount
			return res, commission, err
		case 18:
			res = amount + percent.Percent(commission, int(amount))
			return res, commission, err
		case 24:
			res = amount + percent.Percent(commission*2, int(amount))
			return res, commission, err
		default:
			err = errors.New("invalid period")
			return res, commission, err
		}
	case tv:
		commission = fee[tv]
		switch period {
		case 3, 6, 9, 12, 18:
			res = amount
			return res, commission, err
		case 24:
			res = amount + percent.Percent(commission, int(amount))
			return res, commission, err
		default:
			err = errors.New("invalid period")
			return res, commission, err
		}
	default:
		err = errors.New("[ERROR] invalid product name")
		return res, commission, err
	}

	return 0, 0, nil
}
