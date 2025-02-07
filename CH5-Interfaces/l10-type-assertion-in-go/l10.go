package l10typeassertioningo

func getExpenseReport(e expense) (string, float64) {
	// check if the interface is an email struct or no
	// if it is, we can then call the properties of email
	// same with sms below
	email, isEmail := e.(email)

	if isEmail {
		return email.toAddress, email.cost()
	}

	sms, isSms := e.(sms)

	if isSms {
		return sms.toPhoneNumber, sms.cost()
	}

	return "", 0.0
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
