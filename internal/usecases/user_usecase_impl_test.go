package usecases

// import (
// 	"errors"
// 	"testing"

// 	mock_interfaces "github.com/abdullahnettoor/food-delivery-eCommerce/mock"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/mock/gomock"
// )

// func TestVerifyOtp(t *testing.T) {
// 	tests := []struct {
// 		name       string
// 		prepare    func(mock *mock_interfaces.MockIUserUseCase)
// 		verifyFunc func(t *testing.T, err error)
// 	}{
// 		{
// 			name: "Success",
// 			prepare: func(mock *mock_interfaces.MockIUserUseCase) {
// 				mock.EXPECT().VerifyOtp(gomock.Any(), gomock.Any()).Return(nil)
// 			},
// 			verifyFunc: func(t *testing.T, err error) {
// 				assert.Nil(t, err)
// 			},
// 		},
// 		{
// 			name: "InvalidOtp",
// 			prepare: func(mock *mock_interfaces.MockIUserUseCase) {
// 				mock.EXPECT().VerifyOtp(gomock.Any(), gomock.Any()).Return(errors.New("invalid otp"))
// 			},
// 			verifyFunc: func(t *testing.T, err error) {
// 				assert.EqualError(t, err, "invalid otp")
// 			},
// 		},
// 		{
// 			name: "VerifyUserError",
// 			prepare: func(mock *mock_interfaces.MockIUserUseCase) {
// 				mock.EXPECT().VerifyOtp(gomock.Any(), gomock.Any()).Return(nil)
// 				// mock.EXPECT().Verify(gomock.Any()).Return(errors.New("user verification error"))
// 			},
// 			verifyFunc: func(t *testing.T, err error) {
// 				assert.EqualError(t, err, "user verification error")
// 			},
// 		},
// 	}

// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			mockUserUseCase := mock_interfaces.NewMockIUserUseCase(ctrl)
// 			tt.prepare(mockUserUseCase)
// 			err := mockUserUseCase.VerifyOtp("123456789", &mock_interfaces.AnyType{})
// 			tt.verifyFunc(t, err)
// 		})
// 	}
// }
