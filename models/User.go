package models

import (
	"errors"
	"strings"
	"time"
	"tsbank/infra/database"
	"tsbank/security"
	"tsbank/tools"
)

type User struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"-"`
	Name      string    `json:"name" gorm:"not null"`
	Email     string    `json:"email" gorm:"unique;not null"`
	CpfCnpj   string    `json:"cpf_cnpj" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	Balance   int64     `json:"balance" gorm:"not null;default:0"`
	Type      string    `json:"type" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null;default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:now()"`
	// Apenas para logar e registrar usuarios
	Pass string `json:"password,omitempty" gorm:"-"`
}

// Metódos de User

// Metódos de Bancos de Dados

// Create cria um usuário no banco de dados
func (user *User) create() error {
	return database.DB.Create(user).Error
}

// FindByEmail busca um usuário pelo email
func (user *User) FindByEmail() error {
	return database.DB.Where("email = ?", user.Email).First(user).Error
}

// FindByCpfCnpj busca um usuário pelo cpf ou cnpj
func (user *User) FindByCpfCnpj() error {
	return database.DB.Where("cpf_cnpj = ?", user.CpfCnpj).First(user).Error
}

// FindById busca um usuário pelo id
func (user *User) FindById() error {
	return database.DB.Where("id = ?", user.ID).First(user).Error
}

// Update atualiza um usuário no banco de dados
func (user *User) update() error {
	return database.DB.Model(user).Updates(user).Error
}

// Metódos de Negócio

func (user *User) CanTransfer() bool {
	return user.Type == "pf"
}

// Deposit realiza um depósito na conta do usuário
func (user *User) Deposit(value int64) error {
	if value <= 0 {
		return errors.New("Valor inválido")
	}

	user.Balance += value
	return user.update()
}

// Withdraw realiza um saque na conta do usuário
func (user *User) Withdraw(value int64) error {
	if value <= 0 {
		return errors.New("Valor inválido")
	}

	if user.Balance < value {
		return errors.New("Saldo insuficiente")
	}

	user.Balance -= value
	return user.update()
}

func (user *User) Create() error {
	// Validações
	// Nome
	user.Name = strings.TrimSpace(user.Name)
	if user.Name == "" {
		return errors.New("Nome inválido")
	}
	if len(user.Name) < 3 {
		return errors.New("Nome muito curto")
	}

	// Email
	user.Email = strings.TrimSpace(user.Email)
	user.Email = strings.ToLower(user.Email)
	if user.Email == "" {
		return errors.New("Email inválido")
	}
	if len(user.Email) < 3 {
		return errors.New("Email muito curto")
	}

	user.FindByEmail()
	if user.ID != 0 {
		return errors.New("Email já cadastrado")
	}

	// CpfCnpj
	user.CpfCnpj = tools.RemoveNonNumbers(user.CpfCnpj)
	if len(user.CpfCnpj) != 11 && len(user.CpfCnpj) != 14 {
		return errors.New("CpfCnpj inválido")
	}

	user.FindByCpfCnpj()
	if user.ID != 0 {
		return errors.New("Cpf ou Cnpj já cadastrado")
	}

	//Valida CPF ou CNPJ
	if len(user.CpfCnpj) == 11 {
		if !tools.ValidateCPF(user.CpfCnpj) {
			return errors.New("CPF inválido")
		}
	} else {
		if !tools.ValidateCNPJ(user.CpfCnpj) {
			return errors.New("CNPJ inválido")
		}
	}

	// Tipo
	if len(user.CpfCnpj) == 11 {
		user.Type = "pf"
	} else {
		user.Type = "pj"
	}

	// Saldo
	user.Balance = 0

	// Created At e Updated At
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Senha
	if user.Pass == "" {
		return errors.New("Senha inválida")
	}
	if len(user.Pass) < 4 {
		return errors.New("Senha muito curta, mínimo 4 caracteres")
	}
	user.SetPassword(user.Pass)
	user.Pass = ""

	return user.create()
}

// Metódos de Segurança

// CheckPassword verifica se a senha é igual a senha armazenada no banco de dados
func (user *User) CheckPassword(password string) bool {
	return security.PassCompare(password, user.Password)
}

// SetPassword gera um hash da senha e armazena no banco de dados
func (user *User) SetPassword(password string) {
	var err error
	user.Password, err = security.PassHash(password)
	if err != nil {
		panic(err)
	}
}
