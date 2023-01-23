package handlers

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h handler) GetDays(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		errorLog.Fatal(err)
		return
	}
	t2 := time.Date(id, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	t1 := time.Now()
	if id > t1.Year() {
		days := t2.Sub(t1).Hours() / 24
		c.String(http.StatusOK, "Days left: %v\n", int(days))
	} else {
		days := t1.Sub(t2).Hours() / 24
		c.String(http.StatusOK, "Days gone: %v\n", int(days))
	}
}
