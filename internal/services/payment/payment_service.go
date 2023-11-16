package payment

import (
	"errors"
	"fmt"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/razorpay/razorpay-go"
	"github.com/spf13/viper"
)

type PaymentService struct {
	client *razorpay.Client
}

func NewPaymentService(client *razorpay.Client) *PaymentService {
	return &PaymentService{client}
}

func (s *PaymentService) InitRazorpayClient() *razorpay.Client {
	keyId := viper.GetString("PAYMENT_KEY_ID")
	keySecret := viper.GetString("PAYMENT_KEY_SECRET")
	client := razorpay.NewClient(keyId, keySecret)
	return client
}

func (s *PaymentService) CreatePaymentOrder(amt float64) (map[string]any, error) {
	client := s.InitRazorpayClient()
	options := map[string]any{
		"amount":   uint(amt * 100),
		"currency": "INR",
	}
	mapp, err := client.Order.Create(options, nil)
	fmt.Println("Client Order", mapp)
	return mapp, err
}

func (s *PaymentService) VerifyPayment(orderId, paymentId, receivedSignature string) error {
	client := s.InitRazorpayClient()

	// Extract orderID and paymentID from payment callback
	orderID := orderId
	paymentID := paymentId

	// Get secret key from when you initialized Razorpay
	// secret := viper.GetString("PAYMENT_KEY_SECRET")
	secret := client.Account.Request.Auth.Secret

	// Concatenate orderID and paymentID with '|'
	data := orderID + "|" + paymentID

	// Generate the HMAC hash using sha256 algorithm
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	generatedSignature := hex.EncodeToString(h.Sum(nil))

	// Compare signatures
	if generatedSignature != receivedSignature {
		println("Signature mismatch - payment failed")
		return errors.New("payment is not legit / is failed")
	}
	// Payment is legit!
	println("Payment verified successfully")
	return nil
}
