package dingtalk

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func SendJSON(accessToken string, data string) error {
	resp, err := http.Post("https://oapi.dingtalk.com/robot/send?access_token="+accessToken,
		"application/json",
		strings.NewReader(`
{
     "msgtype": "markdown",
     "markdown": {
         "title":"代码质量分析",
         "text": "## 代码质量分析
	 - 项目：cloud-hub-backend
	 - 分支：publish/szgzw_release
	 - Bugs: 1.6K
	 - 漏洞: 3
	 - 重复率: 44.4%
	 - 查看>>[详情](http://10.14.31.161:31007/dashboard?id=cloud-hub-backend_origin-publish-szgzw_release)<<
	 "
     },
      "at": {
          "atMobiles": [],
          "isAtAll": false
      }
 }
			`))
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println(string(body))
	return nil
}
