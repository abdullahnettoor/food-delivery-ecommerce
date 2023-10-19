package helpers

import (
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	accountSid string
	authToken  string
	serviceSid string
)

func SendOtp(phoneNo string) error {
	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("ACCOUNT_AUTH_TOKEN")
	serviceSid = os.Getenv("SERVICE_SID")

	// Find your Account SID and Auth Token at twilio.com/console
	// and set the environment variables. See http://twil.io/secure
	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: accountSid, Password: authToken})

	params := &verify.CreateVerificationParams{}
	params.SetTo(phoneNo)
	params.SetChannel("sms")

	resp, err := client.VerifyV2.CreateVerification(serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		if resp.Status != nil {
			fmt.Println(*resp.Status)
		} else {
			fmt.Println(resp.Status)
		}
	}
	return err
}

func VerifyOtp(phoneNo, otp string) (string, error) {
	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("ACCOUNT_AUTH_TOKEN")
	serviceSid = os.Getenv("SERVICE_SID")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: accountSid, Password: authToken})

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNo)
	params.SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		return *resp.Status, err
	}

	if resp.Status != nil {
		fmt.Println(*resp.Status)
		return *resp.Status, err
	}

	fmt.Println(resp.Status)
	return *resp.Status, err
}
