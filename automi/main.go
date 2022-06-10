package main

import (
	"encoding/json"
	"fmt"

	"github.com/vladimirvivien/automi/collectors"
	"github.com/vladimirvivien/automi/stream"
)

type ProcessInstance struct {
	ProcessInstanceId          string        `json:"processInstanceId"`
	FinishTime                 int64         `json:"finishTime"`
	AttachedProcessInstanceIds []interface{} `json:"attachedProcessInstanceIds"`
	SyncAction                 string        `json:"syncAction"`
	BusinessId                 string        `json:"businessId"`
	Title                      string        `json:"title"`
	OriginatorDeptId           string        `json:"originatorDeptId"`
	OperationRecords           []Operation   `json:"operationRecords"`
	Result                     string        `json:"result"`
	BizAction                  string        `json:"bizAction"`
	CreateTime                 int64         `json:"createTime"`
	OriginatorUserid           string        `json:"originatorUserid"`
	ProcessCode                string        `json:"processCode"`
	FormValueVOS               []FormValue   `json:"formValueVOS"`
	Tasks                      []Task        `json:"tasks"`
	OriginatorDeptName         string        `json:"originatorDeptName"`
	Status                     string        `json:"status"`
	SyncSeq                    string        `json:"syncSeq"`
}

type Operation struct {
	Date   int64  `json:"date"`
	Result string `json:"result"`
	Type   string `json:"type"`
	UserId string `json:"userId"`
	Remark string `json:"remark,omitempty"`
}

type FormValue struct {
	ComponentType string `json:"componentType"`
	Name          string `json:"name"`
	BizAlias      string `json:"bizAlias,omitempty"`
	Id            string `json:"id"`
	Value         string `json:"value,omitempty"`
	ExtValue      string `json:"extValue,omitempty"`
}

type Task struct {
	Result     string `json:"result"`
	ActivityId string `json:"activityId"`
	FinishTime int64  `json:"finishTime"`
	PcUrl      string `json:"pcUrl"`
	CreateTime int64  `json:"createTime"`
	MobileUrl  string `json:"mobileUrl"`
	UserId     string `json:"userId"`
	TaskId     int64  `json:"taskId"`
	Status     string `json:"status"`
}

func main() {

	js := `{
  "processInstanceId" : "dfdfddjggjkghh",
  "attachedProcessInstanceIds" : [ ],
  "syncAction" : "org_bpms",
  "businessId" : "20220610105723",
  "title" : "张三提交的审批",
  "originatorDeptId" : "234",
  "operationRecords" : [ {
    "date" : 1654681563046,
    "result" : "NONE",
    "type" : "START_PROCESS_INSTANCE",
    "userId" : "454d"
  }],
  "result" : "agree",
  "bizAction" : "NONE",
  "createTime" : 1654681563000,
  "originatorUserid" : "tetuserid",
  "processCode" : "PROC-dfdf",
  "formValueVOS" : [ {
    "componentType" : "DDPhotoField",
    "name" : "图片",
    "id" : "图片"
  } ],
  "tasks" : [ {
    "result" : "NONE",
    "activityId" : "f295_1e01",
    "pcUrl" : "aflow.dingtalk.com",
    "createTime" : 1654681907000,
    "mobileUrl" : "aflow.dingtalk.com",
    "userId" : "dbgdfd",
    "taskId" : 74298840600,
    "status" : "RUNNING"
  }],
  "originatorDeptName" : "测试",
  "status" : "RUNNING",
  "syncSeq" : "3456fdgfigjfkshd"
}`

	ps := &ProcessInstance{}
	_ = json.Unmarshal([]byte(js), ps)

	tasks := make([]Task, 0, 5)
	if err := <-stream.
		New(ps.Tasks).
		Filter(func(item Task) bool {
			return item.Result == "AGREE"
		}).
		Into(collectors.Func(func(data interface{}) error {
			task := data.(Task)
			tasks = append(tasks, task)
			return nil
		})).
		Open(); err != nil {
		fmt.Println(err)
	}

	for _, task := range tasks {
		fmt.Println(task)
	}
}
