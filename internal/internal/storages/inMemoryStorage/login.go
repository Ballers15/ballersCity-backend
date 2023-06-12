package inMemoryStorage

import (
	"fmt"
	"github.com/nostressdev/tvx/internal/storages"
)

var (
	dialerHost     = "smtp.eu.mailgun.org"
	dialerPort     = 587
	dialerUserName = "postmaster@nostress.dev"
	dialerPassword = "1089c9fa0b314508b595345c307da9e6-1831c31e-bcff488c"
)

func (s *IMStorage) Login(email string) (token string, err error) {

	accessToken, err := storages.GenerateOneTimePassword(32)
	if err != nil {
		return "", err
	}
	for {
		if _, ok := s.UserAccessTokens[accessToken]; !ok {
			s.UserAccessTokens[accessToken] = email
			if _, ok := s.UserSettings[email]; !ok {
				s.UserSettings[email] = &UserSettings{StringMap: make(map[string]string),
					IntMap:   make(map[string]int64),
					FloatMap: make(map[string]float64)}
			}
			fmt.Printf("Verification of user '%s' was completed! New token: '%s'\n", email, accessToken)
			return accessToken, nil
		}
		accessToken, err = storages.GenerateOneTimePassword(32)
		if err != nil {
			return "", err
		}
	}

	//msg := gomail.NewMessage()
	//msg.SetHeader("From", "noreply@nostress.dev")
	//msg.SetHeader("To", email)
	//msg.SetHeader("Subject", "TvX Login Verification")
	//verificationCode, err := storages.GenerateOneTimePassword(-1)
	//if err != nil {
	//	return err
	//}
	//msg.SetBody("text/html", verificationCode)
	//dialer := gomail.NewDialer(dialerHost, dialerPort, dialerUserName, dialerPassword)
	//if err := dialer.DialAndSend(msg); err != nil {
	//	return err
	//}
	//s.EmailVerificationCodes[email] = verificationCode
	//fmt.Printf("Email: '%s' | Verification Code: '%s'\n", email, s.EmailVerificationCodes[email])
	//var token string
	//for tk, em := range s.UserAccessTokens {
	//	if em == email {
	//		token = tk
	//		break
	//	}
	//}
	//if token != "" {
	//	delete(s.UserAccessTokens, token)
	//	fmt.Printf("user '%s' was already logged in with token '%s'. Previous token has been deleted.\n", email, token)
	//}
}
