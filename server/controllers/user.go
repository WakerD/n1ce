package controllers

import (
//	"net/http"
//"n1ce/models"
//kkk	"github.com/gin-gonic/gin"
)

type MsgResult struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Obj  string `json:"obj"`
}

/*
func SmsSendCode(phone string) (MsgResult, error) {
	client := &http.Client{}
	parama := "phone=" + phone
	req, err := http.NewRequest("POST", url, strings.NewReader(parama))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var res MsgResult
	err := json.Unmarshal(body, &res)
	return res, err
}*/
