package spass

import (
	"strconv"
)

type Password struct {
	ID                  uint
	Origin_URL          string
	Action_URL          string
	Username_Element    string
	Username_Value      string
	ID_TZ_Enc           string
	Password_Element    string
	Password_Value      string
	PW_TZ_Enc           string
	Host_URL            string
	SSL_Valid           string
	Preferred           string
	Blacklisted_By_User string
	Use_Additional_Auth string
	CM_API_Support      string
	Created_Time        uint
	Modified_Time       uint
	Title               string
	Favicon             []byte
	Source_Type         string
	App_Name            string
	Package_Name        string
	Package_Signature   string
	Reserved_1          string
	Reserved_2          string
	Reserved_3          string
	Reserved_4          string
	Reserved_5          string
	Reserved_6          string
	Reserved_7          string
	Reserved_8          string
	Credential_Memo     string
	OTP                 string
}

func parsePassword(data []string) (Password, error) {
	var pass Password

	id, err := strconv.ParseUint(data[0], 10, 0)
	if err != nil {
		return Password{}, err
	}

	created_time, err := strconv.ParseUint(data[15], 10, 0)
	if err != nil {
		return Password{}, err
	}

	modified_time, err := strconv.ParseUint(data[16], 10, 0)
	if err != nil {
		return Password{}, err
	}

	pass.ID = uint(id)
	pass.Origin_URL = data[1]
	pass.Action_URL = data[2]
	pass.Username_Element = data[3]
	pass.Username_Value = data[4]
	pass.ID_TZ_Enc = data[5]
	pass.Password_Element = data[6]
	pass.Password_Value = data[7]
	pass.PW_TZ_Enc = data[8]
	pass.Host_URL = data[9]
	pass.SSL_Valid = data[10]
	pass.Preferred = data[11]
	pass.Blacklisted_By_User = data[12]
	pass.Use_Additional_Auth = data[13]
	pass.CM_API_Support = data[14]
	pass.Created_Time = uint(created_time)
	pass.Modified_Time = uint(modified_time)
	pass.Title = data[17]
	pass.Favicon = []byte(data[18])
	pass.Source_Type = data[19]
	pass.App_Name = data[20]
	pass.Package_Name = data[21]
	pass.Package_Signature = data[22]
	pass.Reserved_1 = data[23]
	pass.Reserved_2 = data[24]
	pass.Reserved_3 = data[25]
	pass.Reserved_4 = data[26]
	pass.Reserved_5 = data[27]
	pass.Reserved_6 = data[28]
	pass.Reserved_7 = data[29]
	pass.Reserved_8 = data[30]
	pass.Credential_Memo = data[31]
	pass.OTP = data[32]

	return pass, nil
}

type Card struct {
	ID                    uint
	Card_Number_Encrypted string
	First_Six_Digit       string
	Last_Four_Digit       string
	Name_On_Card          string
	Expiration_Month      string
	Expiration_Year       string
	Billing_Address_ID    uint
	Vault_Status          string
	Date_Modified         uint
	Reserved_4            string
	Reserved_5            string
	Reserved_6            string
}

func parseCard(data []string) (Card, error) {
	var card Card

	id, err := strconv.ParseUint(data[0], 10, 0)
	if err != nil {
		return Card{}, err
	}

	billing_address_id, err := strconv.ParseUint(data[7], 10, 0)
	if err != nil {
		return Card{}, err
	}

	modified, err := strconv.ParseUint(data[9], 10, 0)
	if err != nil {
		return Card{}, err
	}

	card.ID = uint(id)
	card.Card_Number_Encrypted = data[1]
	card.First_Six_Digit = data[2]
	card.Last_Four_Digit = data[3]
	card.Name_On_Card = data[4]
	card.Expiration_Month = data[5]
	card.Expiration_Year = data[6]
	card.Billing_Address_ID = uint(billing_address_id)
	card.Vault_Status = data[8]
	card.Date_Modified = uint(modified)
	card.Reserved_4 = data[10]
	card.Reserved_5 = data[11]
	card.Reserved_6 = data[12]

	return card, nil
}

type Address struct {
	ID             uint
	Full_Name      string
	Company_Name   string
	Street_Address string
	City           string
	State          string
	Zipcode        string
	Country_Code   string
	Phone_Number   string
	Email          string
	Date_Modified  uint
	Reserved_4     string
	Reserved_5     string
	Reserved_6     string
}

func parseAddress(data []string) (Address, error) {
	var address Address

	id, err := strconv.ParseUint(data[0], 10, 0)
	if err != nil {
		return Address{}, err
	}

	date_modified, err := strconv.ParseUint(data[10], 10, 0)
	if err != nil {
		return Address{}, err
	}

	address.ID = uint(id)
	address.Full_Name = data[1]
	address.Company_Name = data[2]
	address.Street_Address = data[3]
	address.City = data[4]
	address.State = data[5]
	address.Zipcode = data[6]
	address.Country_Code = data[7]
	address.Phone_Number = data[8]
	address.Email = data[9]
	address.Date_Modified = uint(date_modified)
	address.Reserved_4 = data[11]
	address.Reserved_5 = data[12]
	address.Reserved_6 = data[13]

	return address, nil
}

type Note struct {
	ID            uint
	Note_Title    string
	Note_Details  string
	Date_Modified uint
}

func parseNote(data []string) (Note, error) {
	var note Note

	id, err := strconv.ParseUint(data[0], 10, 0)
	if err != nil {
		return Note{}, err
	}

	date_modified, err := strconv.ParseUint(data[3], 10, 0)
	if err != nil {
		return Note{}, err
	}

	note.ID = uint(id)
	note.Note_Title = data[1]
	note.Note_Details = data[2]
	note.Date_Modified = uint(date_modified)

	return note, nil
}

// Struct that represents the .spass data.
type SPASS struct {
	Version   uint
	Passwords []Password
	Cards     []Card
	Addresses []Address
	Notes     []Note
}
