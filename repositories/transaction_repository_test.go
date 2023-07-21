package repositories

import (
	"kredit-plus/models"
	"net/http"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// assert mockDB
func assertMockDB() (mockgormDB *gorm.DB, mockSqlmock sqlmock.Sqlmock) {
	mockDB, mock, _ := sqlmock.New()
	defer mockDB.Close()

	// Create a new GORM database connection using the mock database
	db, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})

	return db, mock
}

func Test_Repositories_CreateTransaction_ShouldReturnError500DBClosed(t *testing.T) {
	// Arrage
	db, mock := assertMockDB()
	repoAdmin := transaksiRepository{
		db: db,
	}
	newTransaction := &models.Transaksi{
		IDKonsumen:   1,
		NomorKontrak: "Test",
	}
	transactionRequest := models.TransaksiRequest{
		IDKonsumen: 1,
		OTR:        1000,
	}
	expectedQuery := "INSERT INTO `Admin` (`name`,`email`) VALUES (?,?)"
	mockQueryResult := sqlmock.NewResult(1, 1)

	mock.ExpectExec(expectedQuery).WithArgs(newTransaction.IDKonsumen, newTransaction.NomorKontrak).WillReturnResult(mockQueryResult)

	// call the create method on the repository
	resultChannel := repoAdmin.CreateTransaction("token_dummy", transactionRequest)
	result := <-resultChannel
	assert.Equal(t, result.StatusCode, http.StatusInternalServerError)
}
