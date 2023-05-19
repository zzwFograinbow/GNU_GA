package v1

import (
	"Gin_study_new/models"
	"Gin_study_new/pkg/e"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

//type Requests struct {
//	PatID  string       `json:"pat_id"`
//	UserID string       `json:"user_id"`
//	Type   string       `json:"type"`
//	Param  RequestParam `json:"param"`
//}
//type RequestParam struct {
//	ModelType1          string `json:"model_type-1"`
//	ModelType2          string `json:"model_type-2"`
//	OutputFormat        string `json:"outputFormat"`
//	NumThreadsPreproc   int    `json:"num_threads_preprocessing"`
//	NumThreadsNiftiSave int    `json:"num_threads_nifti_save"`
//	Organs              string `json:"Organs"`
//}

func ClientPost(c *gin.Context) {
	//file, _ := c.FormFile("C:\\Users\\Administrator\\Desktop\\demo.json")
	//以下方式来进行对多包涵的json来解析，这里定义一个已经包含的结构体来解析
	var req struct {
		Requests models.Requests `json:"requests"`
	}

	code := e.SUCCESS
	fmt.Println("ClientPost")
	//fmt.Println(models.RequestParam.)

	//if err := c.BindJSON(&req); err != nil {
	//	code = e.ERROR_EXIST_TAG
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": code,
	//		"msg":  e.GetMsg(code),
	//		"data": make(map[string]string),
	//	})
	//	return
	//}

	/*以下为读取json文件所用代码*/
	/*以下代码需要优化有错误代码形式*/
	jsonData, err0 := ioutil.ReadFile("C:\\Users\\Administrator\\Desktop\\demo.json")
	if err0 != nil {
		log.Println(err0)
		c.JSON(500, gin.H{"error": "Failed to read JSON file"})
		return
	}
	if err := json.Unmarshal(jsonData, &req); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "Invalid JSON"})
		return
	}

	fmt.Println("9000000000000000")
	models.ClientAddDatabase(req)
	fmt.Println(req.Requests.Param.Organs[0])
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
