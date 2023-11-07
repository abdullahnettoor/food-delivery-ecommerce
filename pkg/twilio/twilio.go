package otphelper

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/twilio/twilio-go"
	verify "github.com/twilio/twilio-go/rest/verify/v2"
)

var (
	accountSid string
	authToken  string
	serviceSid string
)

func SendOtp(phoneNo string) error {
	accountSid = viper.GetString("ACCOUNT_SID")
	authToken = viper.GetString("ACCOUNT_AUTH_TOKEN")
	serviceSid = viper.GetString("SERVICE_SID")

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

func VerifyOtp(phoneNo, otp string) (bool, error) {
	accountSid = viper.GetString("ACCOUNT_SID")
	authToken = viper.GetString("ACCOUNT_AUTH_TOKEN")
	serviceSid = viper.GetString("SERVICE_SID")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{Username: accountSid, Password: authToken})

	params := &verify.CreateVerificationCheckParams{}
	params.SetTo(phoneNo).SetCode(otp)

	resp, err := client.VerifyV2.CreateVerificationCheck(serviceSid, params)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	if resp.Status != nil {
		fmt.Println(*resp.Status)
		fmt.Println("Status is", (*resp.Status == "approved"))
		return *resp.Status == "approved", err
	}
	return false, err

}
