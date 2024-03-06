package transformer

import (
	"bytes"
	"fmt"

	"github.com/baimiyishu13/wechatrobot-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to wechat markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.WeChatMarkdown, robotURL string, err error) {

	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["wechatRobot"]

	var buffer bytes.Buffer

	for _, alert := range notification.Alerts {
		labels := alert.Labels
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n # 告警: <font color='warning'> %s </font>\n", annotations["summary"]))
		// datacenter 为 victoriametrics 获取 prometheus时区分环境的 label
		buffer.WriteString(fmt.Sprintf("\n>【环境】 %s\n", labels["datacenter"]))
		buffer.WriteString(fmt.Sprintf("\n>【级别】 %s\n", labels["severity"]))
		buffer.WriteString(fmt.Sprintf("\n>【类型】 %s\n", labels["alertname"]))
		buffer.WriteString(fmt.Sprintf("\n>【主机】 %s\n", labels["instance"]))
		buffer.WriteString(fmt.Sprintf("\n>【内容】 %s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n>【当前状态】%s \n", status))
		buffer.WriteString(fmt.Sprintf("\n>【触发时间】 %s\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		buffer.WriteString("\n [跳转Grafana看板](http://addr:3000/dashboards)")
		buffer.WriteString("\n [屏蔽告警](http://addr:9093/#/alerts)")
		buffer.WriteString("\n @运维1 @运维2")
	}

	markdown = &model.WeChatMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Content: buffer.String(),
		},
	}

	return
}
