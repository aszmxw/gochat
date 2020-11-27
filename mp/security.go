package mp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/shenghui0779/gochat/helpers"
	"github.com/tidwall/gjson"
)

// SecCheckMediaType 检测的素材类型
type SecCheckMediaType int

var (
	// SecCheckMediaAudio 音频
	SecCheckMediaAudio = 1
	// SecCheckMediaImage 图片
	SecCheckMediaImage = 2
)

// ImgSecCheck 校验一张图片是否含有违法违规内容
func ImgSecCheck(filename string) Action {
	return &WechatAPI{
		body: helpers.NewUploadBody("media", filename, func() ([]byte, error) {
			return ioutil.ReadFile(filename)
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("UPLOAD|%s?access_token=%s", ImgSecCheckURL, accessToken)
		},
	}
}

// MediaCheckAsyncInfo 任务id，用于匹配异步推送结果
type MediaCheckAsyncInfo struct {
	TraceID string
}

// MediaCheckAsync 异步校验图片/音频是否含有违法违规内容
func MediaCheckAsync(mediaType SecCheckMediaType, mediaURL string, dest *MediaCheckAsyncInfo) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"media_type": mediaType,
				"media_url":  mediaURL,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", MediaCheckAsyncURL, accessToken)
		},
		decode: func(resp []byte) error {
			dest.TraceID = gjson.GetBytes(resp, "trace_id").String()

			return nil
		},
	}
}

// MsgSecCheck 检查一段文本是否含有违法违规内容
func MsgSecCheck(content string) Action {
	return &WechatAPI{
		body: helpers.NewPostBody(func() ([]byte, error) {
			return json.Marshal(helpers.X{
				"content": content,
			})
		}),
		url: func(accessToken string) string {
			return fmt.Sprintf("POST|%s?access_token=%s", MsgSecCheckURL, accessToken)
		},
	}
}
