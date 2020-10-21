package middlewars

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

//authentication olahjdoah*/
func Authentication(username, password string, c echo.Context) (bool, error) {
	//check against DB
	fmt.Println(username)
	if password == "pass" {
		return true, nil
	}
	return false, nil
}
