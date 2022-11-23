package db

import (
	"appointments-system/models"
)

func (pr *postgresRepo) GetPaymentMethods() ([]models.PaymentMethod, error) {
	var paymentMethods []models.PaymentMethod

	err := pr.DB.Select(
		&paymentMethods,
		`SELECT id, name FROM payment_methods`,
	)
	if err != nil {
		return paymentMethods, err
	}

	return paymentMethods, nil
}

func (pr *postgresRepo) GetPaymentMethod(id int) (models.PaymentMethod, error) {
	var paymentMethod models.PaymentMethod

	err := pr.DB.Get(
		&paymentMethod,
		`SELECT id, name 
		FROM payment_methods
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return paymentMethod, err
	}

	return paymentMethod, nil
}
