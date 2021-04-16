package openwechat

import (
	"errors"
	"regexp"
)

var (
	uuidRegexp        = regexp.MustCompile(`uuid = "(.*?)";`)
	statusCodeRegexp  = regexp.MustCompile(`window.code=(\d+);`)
	syncCheckRegexp   = regexp.MustCompile(`window.synccheck=\{retcode:"(\d+)",selector:"(\d+)"\}`)
	redirectUriRegexp = regexp.MustCompile(`window.redirect_uri="(.*?)"`)
)

const (
	appId = "wx782c26e4c19acffb"

	baseUrl                 = "https://wx.qq.com"
	jsLoginUrl              = "https://login.wx.qq.com/jslogin"
	webWxNewLoginPage       = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxnewloginpage"
	qrcodeUrl               = "https://login.weixin.qq.com/qrcode/"
	loginUrl                = "https://login.wx.qq.com/cgi-bin/mmwebwx-bin/login"
	webWxInitUrl            = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxinit"
	webWxStatusNotifyUrl    = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxstatusnotify"
	webWxSyncUrl            = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsync"
	webWxSendMsgUrl         = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsendmsg"
	webWxGetContactUrl      = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetcontact"
	webWxSendMsgImgUrl      = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsendmsgimg"
	webWxSendAppMsgUrl      = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxsendappmsg"
	webWxBatchGetContactUrl = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxbatchgetcontact"
	webWxOplogUrl           = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxoplog"
	webWxVerifyUserUrl      = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxverifyuser"
	syncCheckUrl            = "https://webpush.wx.qq.com/cgi-bin/mmwebwx-bin/synccheck"
	webWxUpLoadMediaUrl     = "https://file.wx.qq.com/cgi-bin/mmwebwx-bin/webwxuploadmedia"
	webWxGetMsgImgUrl       = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetmsgimg"
	webWxGetVoiceUrl        = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetvoice"
	webWxGetVideoUrl        = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetvideo"
	webWxLogoutUrl          = "https://wx.qq.com/cgi-bin/mmwebwx-bin/webwxlogout"
	webWxGetMediaUrl        = "https://file.wx.qq.com/cgi-bin/mmwebwx-bin/webwxgetmedia"

	jsonContentType = "application/json; charset=utf-8"
)

// 消息类型
const (
	TextMessage  = 1
	ImageMessage = 3
	AppMessage   = 6
)

// 登录状态
const (
	statusSuccess = "200"
	statusScanned = "201"
	statusTimeout = "400"
	statusWait    = "408"
)

// errors
var (
	noSuchUserFoundError = errors.New("no such user found")
)

const ALL = 0

// sex
const (
	MALE   = 1
	FEMALE = 2
)
