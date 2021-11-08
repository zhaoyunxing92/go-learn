package factory

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func httpHandle(url string, cookie string) {

	payload := strings.NewReader(`{"deptId":-1,"statisticsMonth":""}`)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	//req.Header.Add("authority", "dailysalary.eapps.dingtalkcloud.com")
	//req.Header.Add("sec-ch-ua", "\"Google Chrome\";v=\"93\", \" Not;A Brand\";v=\"99\", \"Chromium\";v=\"93\"")
	//req.Header.Add("content-type", "text/plain;charset=UTF-8")
	//req.Header.Add("x-csrf-token", "3615e0265f6144c5a71898571026d847")
	//req.Header.Add("sec-ch-ua-mobile", "?0")
	//req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36")
	//req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	//req.Header.Add("accept", "*/*")
	//req.Header.Add("origin", "https://dailysalary.eapps.dingtalkcloud.com")
	//req.Header.Add("sec-fetch-site", "same-origin")
	//req.Header.Add("sec-fetch-mode", "same-origin")
	//req.Header.Add("sec-fetch-dest", "empty")
	//req.Header.Add("referer", "https://dailysalary.eapps.dingtalkcloud.com/index")
	//req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func initAppData(cookie string) {
	url := "https://daily-k8s-payslip.renlijia.com/rest/api/v1/paySlip/appInitData"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("cookie", cookie)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
}
