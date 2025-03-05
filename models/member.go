package models

import (
	"time"
)

type Passport struct {
	Ser      int8   `bson:"seria" json:"ser,omitempty"`       //серия
	Number   int16  `bson:"number" json:"number,omitempty"`   //номер паспорта
	Registry int8   `bson:"registr" json:"registr,omitempty"` //тип регистрации
	Address  string `bson:"address" json:"address,omitempty"` //адресс регистрации
}

// Таблица для видов регистрации по месту жительства
type Registry struct {
	znach int8   `bson:"znach" json:"znach,omitempty"` //код регистрации
	Name  string `bson:"name" json:"name,omitempty"`   // наименование
}

// Данные о человеке в системе
type Member struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	FullName   string             `bson:"fullName" validate:"required"`
	Passport   Passport           `bson:"passport"`
	ShareSize  float64            `bson:"shareSize"`
	Apartments []string           `bson:"apartments"` // ObjectIDs как строки
	Payments   []Payment          `bson:"payments"`
	Status     string             `bson:"status" validate:"oneof=active expelled candidate"`
}
type Apartment struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Number     int8               `bson:"number" json:"number,omitempty"`          //Номер квартиры
	Area       float32            `bson:"area" json:"area,omitempty"`              //Площадь квартиры
	IdWm       primitive.ObjectID `bson:"_id,omitempty" json:"id___cwm,omitempty"` //Ссылка на список счетчиков
	Payments   []Payment          `bson:"payments"`                                // Сведения по оплате по квартире
	Ap_account string             `bson:"Ap_account" json:"Ap_Account,omitempty"`  // Лицевой счет привязанный к квартире
}
type Payment struct {
	Date        time.Time `bson:"date"`
	Amount      float64   `bson:"amount"`
	PaymentType string    `bson:"paymentType"`
}
type User struct {
	ChatID          int `bson:"chat_id"`
	ApartmentNumber int `bson:"apartment_number"`
}
type Reading struct {
	ApartmentNumber int       `bson:"apartment_number"`
	Date            time.Time `bson:"date"`
	Cold            float64   `bson:"cold"`
	Hot             float64   `bson:"hot"`
}
type News struct {
	Title   string    `bson:"title"`
	Content string    `bson:"content"`
	Date    time.Time `bson:"date"`
}
