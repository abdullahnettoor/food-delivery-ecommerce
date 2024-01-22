package repository

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/entities"
	e "github.com/abdullahnettoor/food-delivery-eCommerce/internal/domain/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var usrRow = []string{"id", "first_name", "last_name", "email", "password", "phone", "status"}

func TestUserRepository_FindAll(t *testing.T) {

	tests := []struct {
		name    string
		stub    func(sqlmock.Sqlmock)
		want    *[]entities.User
		wantErr error
	}{
		{
			name: "success",
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).
					WillReturnRows(sqlmock.NewRows(usrRow).AddRow(
						1, "Abdu", "Nettoor", "abdullahnettoor@gmail.com", "123456", "+919061904860", "Verified",
					))
			},

			want: &[]entities.User{
				entities.User{
					ID:        1,
					FirstName: "Abdu",
					LastName:  "Nettoor",
					Email:     "abdullahnettoor@gmail.com",
					Password:  "123456",
					Phone:     "+919061904860",
					Status:    "Verified",
				},
			},
			wantErr: nil,
		},
		{
			name: "failure",
			stub: func(mockSQL sqlmock.Sqlmock) {

				mockSQL.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users`)).
					WillReturnRows(sqlmock.NewRows([]string{
						// "id", "first_name", "last_name", "email", "phone", "status",
					}).AddRow(
					// 	1, "Abdu", "Nettoor", "abdullahnettoor@gmail.com", "+919061904860", "Verified",
					// ).AddRow(
					// 	1, "Abdu", "Nettoor", "abdullah@gmail.com", "+919061904870", "Verified",
					))
			},

			// want: &[]entities.User{
			// 	entities.User{
			// 		ID:        1,
			// 		FirstName: "Abdu",
			// 		LastName:  "Nettoor",
			// 		Email:     "abdullahnettoor@gmail.com",
			// 		Phone:     "+919061904860",
			// 		Status:    "Verified",
			// 	},
			// 	entities.User{
			// 		ID:        1,
			// 		FirstName: "Abdu",
			// 		LastName:  "Nettoor",
			// 		Email:     "abdullah@gmail.com",
			// 		Phone:     "+919061904870",
			// 		Status:    "Verified",
			// 	},
			// },
			want:    &[]entities.User{entities.User{}},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSQL)

			u := NewUserRepository(gormDB)

			got, err := u.FindAll()
			if err != tt.wantErr {
				t.Errorf("UserRepository.FindByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_FindByPhone(t *testing.T) {
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    *entities.User
		wantErr error
	}{
		{
			name: "success",
			args: args{phone: "+919061904860"},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE phone = $1`)).
					WithArgs("+919061904860").
					WillReturnRows(s.NewRows(usrRow).
						AddRow(
							1, "Abdu", "Nettoor", "abdu@mail.com", "123456", "+919061904860", "VERIFIED",
						))

			},
			want:    &entities.User{ID: 1, FirstName: "Abdu", LastName: "Nettoor", Email: "abdu@mail.com", Password: "123456", Phone: "+919061904860", Status: "VERIFIED"},
			wantErr: nil,
		},
		{
			name: "fail",
			args: args{phone: "+919061904860"},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE phone = $1`)).
					WithArgs("+919061904860").
					WillReturnRows(sqlmock.NewRows([]string{}).
						AddRow()).WillReturnError(e.ErrNotFound)

			},
			want:    nil,
			wantErr: e.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSql)

			repo := NewUserRepository(gormDB)

			got, err := repo.FindByPhone(tt.args.phone)
			if err != tt.wantErr {
				t.Errorf("UserRepository.FindByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_FindByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		args    args
		stub    func(sqlmock.Sqlmock)
		want    *entities.User
		wantErr error
	}{
		{
			name: "success",
			args: args{email:"abdu@mail.com"},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE email = $1`)).
					WithArgs("abdu@mail.com").
					WillReturnRows(s.NewRows(usrRow).
						AddRow(
							1, "Abdu", "Nettoor", "abdu@mail.com", "123456", "+919061904860", "VERIFIED",
						))

			},
			want:    &entities.User{ID: 1, FirstName: "Abdu", LastName: "Nettoor", Email: "abdu@mail.com", Password: "123456", Phone: "+919061904860", Status: "VERIFIED"},
			wantErr: nil,
		},
		{
			name: "error",
			args: args{email:"abdu@mail.com"},
			stub: func(s sqlmock.Sqlmock) {
				s.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE email = $1`)).
					WithArgs("abdu@mail.com").
					WillReturnRows(sqlmock.NewRows([]string{}).
						AddRow()).WillReturnError(e.ErrNotFound)

			},
			want:    nil,
			wantErr: e.ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{
				Conn: mockDB,
			}), &gorm.Config{})

			tt.stub(mockSql)

			repo := NewUserRepository(gormDB)

			got, err := repo.FindByEmail(tt.args.email)
			if err != tt.wantErr {
				t.Errorf("UserRepository.FindByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_Create(t *testing.T) {
	type args struct {
		usrModel *entities.User
	}
	tests := []struct {
		name    string
		args    args
		query   func(sqlmock.Sqlmock)
		want    *entities.User
		wantErr error
	}{
		{name: "success",
			args: args{usrModel: &entities.User{FirstName: "Test", Email: "test@gmail.com", Phone: "12345678901", Password: "12345678", Status: "PENDING"}},
			query: func(s sqlmock.Sqlmock) {
				query := `SELECT *
						FROM users
						WHERE email = $1
						OR phone = $2`
				s.ExpectQuery(regexp.QuoteMeta(query)).
					WithArgs("test@gmail.com", "12345678901").
					WillReturnRows(sqlmock.NewRows(usrRow).AddRow(
						10, "Test", "", "test@gmail.com", "12345678901", "12345678", "PENDING",
					))
			},
			want:    &entities.User{ID: 10, FirstName: "Test", Email: "test@gmail.com", Phone: "12345678901", Password: "12345678", Status: "PENDING"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSql, _ := sqlmock.New()
			defer mockDB.Close()

			gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: mockDB}), &gorm.Config{})

			tt.query(mockSql)

			repo := NewUserRepository(gormDB)

			got, err := repo.Create(tt.args.usrModel)
			if err != tt.wantErr {
				t.Errorf("UserRepository.FindByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.FindByPhone() = %v, want %v", got, tt.want)
			}
		})
	}
}
