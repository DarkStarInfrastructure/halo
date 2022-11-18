package services

import (
	"fmt"
	"halo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Disk struct {
	// The name of the disk.
	Name string `json:"name"`
}

type Disks struct {
	Disks []Disk `json:"disks"`
}

func ListAllDisks(c *gin.Context) {

	result, err := utils.ExecCommand("lsblk", "-d", "--json")
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
	}

	fmt.Fprintln(gin.DefaultWriter, result)

	c.JSON(http.StatusOK, result)
}

func GetDiskDetails(c *gin.Context) {
	c.JSON(http.StatusOK, "")

}
