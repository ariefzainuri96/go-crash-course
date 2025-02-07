package l1theerrorinterface

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costSendToCustomer, errorSendToCustomer := sendSMS(msgToCustomer)

	if errorSendToCustomer != nil {
		return 0, errorSendToCustomer
	}

	costSendToSpouse, errorSendToSpouse := sendSMS(msgToSpouse)

	if errorSendToSpouse != nil {
		return 0, errorSendToSpouse
	}

	return (costSendToCustomer + costSendToSpouse), nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
