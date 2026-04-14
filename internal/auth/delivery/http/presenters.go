package http

import (
	"regexp"

	"github.com/zeross/project-demo/internal/auth"
)

// Khớp với user admin trong migration 00002 (đăng nhập dev).
const seedAdminLoginPhone = "__seed_admin__"

type loginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (r loginReq) toInput() auth.LoginInput {
	return auth.LoginInput{
		Phone:    r.Phone,
		Password: r.Password,
	}
}

func (r loginReq) validate() error {
	if r.Phone != seedAdminLoginPhone {
		if matched, _ := regexp.MatchString(`^\d{9,15}$`, r.Phone); !matched {
			return errWrongBody
		}
	}

	if len(r.Password) < 6 {
		return errWrongBody
	}

	return nil
}

type authResp struct {
	Token string `json:"token"`
}

func (h handler) newAuthResp(out auth.LoginResponse) authResp {
	return authResp{Token: out.Token}
}
