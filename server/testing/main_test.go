package test

import (
	"os"
	"testing"

	"github.com/skirill430/Quick-Shop/server/router"
	"github.com/skirill430/Quick-Shop/server/utils"
)

var Router = router.Router()

// needed to connect to DB before each test is run
func TestMain(m *testing.M) {
	utils.ConnectDB("users_test")
	utils.ConnectDB("products_test")
	code := m.Run()

	// clear databases after so future tests execute the same
	utils.ClearUsersDB()
	utils.ClearUserProductsDB()
	os.Exit(code)
}
