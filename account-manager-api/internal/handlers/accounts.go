package handlers

import (
	"account-manager-api/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var bankService *service.BankService = service.NewBankService()

func CreateAccount(c *gin.Context) {
	var data struct {
		AccountNumber int    `json:"account_number"`
		Owner         string `json:"owner"`
	}
	c.BindJSON(&data)
	bankService.CreateAccount(data.AccountNumber, data.Owner)
	c.JSON(http.StatusOK, gin.H{"message": "Conta criada com sucesso"})
}

func DeleteAccount(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	bankService.DeleteAccount(id)
	c.JSON(http.StatusOK, gin.H{"message": "Conta deletada com sucesso"})
}

func Deposit(c *gin.Context) {
	var data struct {
		Amount float64 `json:"amount"`
	}
	id, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&data)
	err := bankService.Deposit(id, data.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Depósito realizado com sucesso"})
	}
}

func Withdraw(c *gin.Context) {
	var data struct {
		Amount float64 `json:"amount"`
	}
	id, _ := strconv.Atoi(c.Param("id"))
	c.BindJSON(&data)
	err := bankService.Withdraw(id, data.Amount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Saque realizado com sucesso"})
	}
}

func Transfer(c *gin.Context) {
	var data struct {
		FromAccount int     `json:"from_account"`
		ToAccount   int     `json:"to_account"`
		Amount      float64 `json:"amount"`
	}
	c.BindJSON(&data)
	err := bankService.Transfer(data.Amount, data.FromAccount, data.ToAccount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Transferência realizada com sucesso"})
	}
}

func GetAccounts(c *gin.Context) {
	accounts := bankService.GetAccounts()
	c.JSON(http.StatusOK, accounts)
}
