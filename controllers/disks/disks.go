package disks

import (
	"encoding/json"
	"fmt"
	"halo/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DiskController struct {
	Blockdevices []Blockdevices `json:"blockdevices"`
}

type Blockdevices struct {
	Name       string      `json:"name"`
	MajMin     string      `json:"maj:min"`
	Rm         string      `json:"rm"`
	Size       string      `json:"size"`
	Ro         string      `json:"ro"`
	Type       string      `json:"type"`
	Mountpoint interface{} `json:"mountpoint"`
}

func (d DiskController) ListAllDisks(c *gin.Context) {
	// -d, --nodeps         don't print slaves or holders
	// -O, --output-all     output all columns
	// -b, --bytes          print SIZE in bytes rather than in human readable format
	// -J, --json           use JSON output format
	result, err := utils.ExecCommand("lsblk", "-d", "-O", "-b", "--json")
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Fprintln(gin.DefaultWriter, result)

	c.JSON(http.StatusOK, gin.H{"result": json.RawMessage(result)})
}
