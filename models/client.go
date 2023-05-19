package models

import (
	"fmt"
	"time"
)

type Requests struct {
	PatID  string       `json:"pat_id"`
	UserID string       `json:"user_id"`
	Type   string       `json:"type"`
	Param  RequestParam `json:"param"`
}
type RequestParam struct {
	ModelType1          string   `json:"model_type-1"`
	ModelType2          string   `json:"model_type-2"`
	OutputFormat        []string `json:"outputFormat"`
	NumThreadsPreproc   string   `json:"num_threads_preprocessing"`
	NumThreadsNiftiSave string   `json:"num_threads_nifti_save"`
	Organs              []string `json:"Organs"`
}

func ClientAddDatabase(databaseTable struct {
	Requests Requests `json:"requests"`
}) bool {
	var count int
	var reID int
	db.Model(&Rec{}).Count(&count)
	reID = count + 1
	fmt.Println(reID)
	db.Create(&Rec{ //这个数据库后面的名字就是这么来的
		ReID:          reID,
		PatID:         databaseTable.Requests.PatID,
		UserID:        databaseTable.Requests.UserID,
		Type:          databaseTable.Requests.Type,
		Status:        0,
		AlgoServer_ip: "192.168.1.189",
		Create_time:   time.Now().String(),
		Done_time:     time.Now().String(),
	})
	return true
}
