package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"login-demo/internal/dao"
	"login-demo/internal/model/do"
	"net/http"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type WxLoginLogic struct{}

var WxLogin WxLoginLogic

var WxCache map[string]WechatLoginResponse = make(map[string]WechatLoginResponse)

var WxValidcode map[string]string = map[string]string{}

// 微信相关操作
func (*WxLoginLogic) WechatLogin(ctx context.Context, code string) (*WechatLoginResponse, error) {
	value, ok := WxCache[code]
	if !ok {
		url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", g.Cfg().MustGet(ctx, "wechat.appId"), g.Cfg().MustGet(ctx, "wechat.appSecret"), code)
		client := &http.Client{}
		wxreq, err := http.NewRequest("GET", url, nil)
		if err != nil {
			glog.Error(ctx, err)
			return nil, gerror.NewCode(gcode.New(403, "微信登录失败", err.Error()))
		}
		wxreq.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(wxreq)
		if err != nil {
			glog.Error(ctx, err)
			return nil, gerror.NewCode(gcode.New(403, "微信登录失败", err.Error()))
		}
		defer resp.Body.Close()

		var data WechatLoginResponse
		err = json.NewDecoder(resp.Body).Decode(&data)
		if err != nil {
			return nil, gerror.NewCode(gcode.New(1, "微信登录失败", "解析微信登录信息失败"))
		}
		if data.ErrCode != 0 {
			glog.Info(ctx, data)
			glog.Error(ctx, data.ErrMsg)
			return nil, gerror.NewCode(gcode.New(1, "微信登录失败", err.Error()))
		}
		WxCache[code] = data
		return &data, nil
	} else {
		return &value, nil
	}

}

type WechatLoginResponse struct {
	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

type WechatPhoneNumber struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	PhoneInfo struct {
		PhoneNumber     string `json:"phoneNumber"`
		PurePhoneNumber string `json:"purePhoneNumber"`
		CountryCode     string `json:"countryCode"`
		Watermark       struct {
			Timestamp int64  `json:"timestamp"`
			AppID     string `json:"appid"`
		} `json:"watermark"`
	} `json:"phone_info"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"` // 7200 秒
}

func (*WxLoginLogic) WechatAccessToken(ctx context.Context) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=%s&appid=%s&secret=%s", "client_credential", g.Cfg().MustGet(ctx, "wechat.appId"), g.Cfg().MustGet(ctx, "wechat.appSecret"))
	// 创建 HTTP 客户端
	client := ghttp.NewClient()

	// 发送 GET 请求
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Close()

	var accessToken AccessToken
	err = json.NewDecoder(resp.Body).Decode(&accessToken)
	return accessToken.AccessToken, err

}

func (*WxLoginLogic) WechatPhone(ctx context.Context, code string, encryptedData string, iv string) (*WechatPhoneNumber, error) {
	accessToken, err := WxLogin.WechatAccessToken(ctx)
	if err != nil {
		return nil, gerror.NewCode(gcode.New(1, "获取 accessToken", err.Error()))
	}
	// value, ok := WxCache[code]
	// if !ok {
	// 	panic("code not found")
	// }

	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/business/getuserphonenumber?access_token=%s", accessToken)

	// 创建 HTTP 客户端
	client := ghttp.NewClient()
	data := map[string]string{}
	data["code"] = code
	// 发送 GET 请求
	resp, err := client.Post(url, data)
	if err != nil {
		panic(err)
	}
	defer resp.Close()
	var phoneNumber WechatPhoneNumber
	err = json.NewDecoder(resp.Body).Decode(&phoneNumber)
	if err != nil {
		return nil, gerror.NewCode(gcode.New(1, "解析手机号失败", err.Error()))
	}
	if phoneNumber.ErrCode != 0 {
		fmt.Println("=====", phoneNumber)
		return nil, gerror.NewCode(gcode.New(phoneNumber.ErrCode, "获取手机号失败", phoneNumber.ErrMsg))
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("=====phoneNumber", phoneNumber)
	return &phoneNumber, err
}

func (*WxLoginLogic) SendValidCode(ctx context.Context, phone string) error {
	WxValidcode[phone] = "666666"
	return nil
}

// RegisterService 是注册相关的 service
type RegisterService struct {
}

func (s *RegisterService) VerifyCode(ctx context.Context, phone, validcode string) (err error) {
	validCode, ok := WxValidcode[phone]
	if ok {
		if validCode == validcode {
			return
		}
	}
	return errors.New("验证码不匹配")
}

var Register RegisterService

// Register 用于注册新用户
func (s *RegisterService) Register(ctx context.Context, phone, nickname, avatarUrl, openid, gender string) (token string, err error) {
	// 添加新用户
	user := &do.User{
		Passport: phone,
		Nickname: nickname,
		Password: "123456",
	}
	err = User.Add(ctx, user)
	if err != nil {
		err = gerror.NewCode(gcode.New(1, "微信注册失败", err.Error()))
		return "", err
	}

	// 添加微信用户信息
	wxuser := do.WxUser{
		UserId:    user.Id,
		OpenId:    openid,
		PhoneNo:   phone,
		AvatarUrl: avatarUrl,
		Nickname:  nickname,
		Gender:    gender,
	}
	dao.WxUser.Ctx(ctx).Insert(wxuser)

	// 生成 JWT token
	token, err = MyJwt.GenerateToken(ctx, phone)
	if err != nil {
		return "", err
	}

	return token, nil
}
