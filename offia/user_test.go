package offia

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/shenghui0779/gochat/wx"
)

func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/user/info?access_token=ACCESS_TOKEN&lang=zh_CN&openid=OPENID").Return([]byte(`{
		"subscribe": 1,
		"openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
		"nickname": "Band",
		"sex": 1,
		"language": "zh_CN",
		"city": "广州",
		"province": "广东",
		"country": "中国",
		"headimgurl": "http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
		"subscribe_time": 1382694957,
		"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL",
		"remark": "",
		"groupid": 0,
		"tagid_list": [128, 2],
		"subscribe_scene": "ADD_SCENE_QR_CODE",
		"qr_scene": 98765,
		"qr_scene_str": ""
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsUserGet{
		OpenID: "OPENID",
	}
	result := new(UserInfo)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUser(params, result))

	assert.Nil(t, err)
	assert.Equal(t, &UserInfo{
		Subscribe:      1,
		OpenID:         "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
		NickName:       "Band",
		Sex:            1,
		Country:        "中国",
		City:           "广州",
		Province:       "广东",
		Language:       "zh_CN",
		HeadImgURL:     "http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
		SubscribeTime:  1382694957,
		UnionID:        "o6_bmasdasdsad6_2sgVt7hMZOPfL",
		Remark:         "",
		GroupID:        0,
		TagidList:      []int64{128, 2},
		SubscribeScene: AddSceneQRCode,
		QRScene:        98765,
		QRSceneStr:     "",
	}, result)
}

func TestBatchGetUserInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=ACCESS_TOKEN", []byte(`{"user_list":[{"lang":"zh_CN","openid":"otvxTs4dckWG7imySrJd6jSi0CWE"},{"lang":"zh_CN","openid":"otvxTs_JZ6SEiP0imdhpi50fuSZg"}]}`)).Return([]byte(`{
		"user_info_list": [
			{
				"subscribe": 1,
				"openid": "otvxTs4dckWG7imySrJd6jSi0CWE",
				"nickname": "iWithery",
				"sex": 1,
				"language": "zh_CN",
				"city": "揭阳",
				"province": "广东",
				"country": "中国",
				"headimgurl": "http://thirdwx.qlogo.cn/mmopen/xbIQx1GRqdvyqkMMhEaGOX802l1CyqMJNgUzKP8MeAeHFicRDSnZH7FY4XB7p8XHXIf6uJA2SCunTPicGKezDC4saKISzRj3nz/0",
			    "subscribe_time": 1434093047,
				"unionid": "oR5GjjgEhCMJFyzaVZdrxZ2zRRF4",
				"remark": "",
				"groupid": 0,
				"tagid_list": [128, 2],
				"subscribe_scene": "ADD_SCENE_QR_CODE",
				"qr_scene": 98765,
				"qr_scene_str": ""
		   },
			{
				"subscribe": 0,
				"openid": "otvxTs_JZ6SEiP0imdhpi50fuSZg"
			}
		]
	 }`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	params := &ParamsUserBatchGet{
		UserList: []*ParamsUserGet{
			{
				OpenID: "otvxTs4dckWG7imySrJd6jSi0CWE",
			},
			{
				OpenID: "otvxTs_JZ6SEiP0imdhpi50fuSZg",
			},
		},
	}
	result := new(ResultUserBatchGet)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BatchGetUser(params, result))

	assert.Nil(t, err)
	assert.Equal(t, []*UserInfo{
		{
			Subscribe:      1,
			OpenID:         "otvxTs4dckWG7imySrJd6jSi0CWE",
			NickName:       "iWithery",
			Sex:            1,
			Country:        "中国",
			City:           "揭阳",
			Province:       "广东",
			Language:       "zh_CN",
			HeadImgURL:     "http://thirdwx.qlogo.cn/mmopen/xbIQx1GRqdvyqkMMhEaGOX802l1CyqMJNgUzKP8MeAeHFicRDSnZH7FY4XB7p8XHXIf6uJA2SCunTPicGKezDC4saKISzRj3nz/0",
			SubscribeTime:  1434093047,
			UnionID:        "oR5GjjgEhCMJFyzaVZdrxZ2zRRF4",
			Remark:         "",
			GroupID:        0,
			TagidList:      []int64{128, 2},
			SubscribeScene: AddSceneQRCode,
			QRScene:        98765,
			QRSceneStr:     "",
		},
		{
			Subscribe: 0,
			OpenID:    "otvxTs_JZ6SEiP0imdhpi50fuSZg",
		},
	}, result)
}

func TestGetUserList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&next_openid=NEXT_OPENID").Return([]byte(`{
		"total": 2,
		"count": 2,
		"data": {
			"openid": ["OPENID1", "OPENID2"]
		},
		"next_openid": "NEXT_OPENID"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(ResultUserList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetUserList("NEXT_OPENID", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultUserList{
		Total: 2,
		Count: 2,
		Data: UserListData{
			OpenID: []string{"OPENID1", "OPENID2"},
		},
		NextOpenID: "NEXT_OPENID",
	}, result)
}

func TestGetBlackList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=ACCESS_TOKEN", []byte(`{"begin_openid":"OPENID1"}`)).Return([]byte(`{
		"total": 3,
		"count": 3,
		"data": {
			"openid": [
			  "OPENID1",
			  "OPENID2",
			  "OPENID10000"
		   	]
		},
		"next_openid": "OPENID10000"
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := new(ResultBlackList)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetBlackList("OPENID1", result))

	assert.Nil(t, err)
	assert.Equal(t, &ResultBlackList{
		Total: 3,
		Count: 3,
		Data: UserListData{
			OpenID: []string{"OPENID1", "OPENID2", "OPENID10000"},
		},
		NextOpenID: "OPENID10000",
	}, result)
}

func TestBlackUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=ACCESS_TOKEN", []byte(`{"openid_list":["OPENID1","OPENID2"]}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", BlackUsers("OPENID1", "OPENID2"))

	assert.Nil(t, err)
}

func TestUnBlackUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=ACCESS_TOKEN", []byte(`{"openid_list":["OPENID1","OPENID2"]}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UnBlackUsers("OPENID1", "OPENID2"))

	assert.Nil(t, err)
}

func TestSetUserRemark(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=ACCESS_TOKEN", []byte(`{"openid":"oDF3iY9ffA-hqb2vVvbr7qxf6A0Q","remark":"pangzi"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", SetUserRemark("oDF3iY9ffA-hqb2vVvbr7qxf6A0Q", "pangzi"))

	assert.Nil(t, err)
}