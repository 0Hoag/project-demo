package http

import (
	"time"

	"github.com/google/uuid"
	domain "github.com/zeross/project-demo/internal/models"
	"github.com/zeross/project-demo/internal/users"
)

type createReq struct {
	Username  string     `json:"username" binding:"required"`
	Phone     string     `json:"phone" binding:"required"`
	Password  string     `json:"password" binding:"required"`
	AvatarUrl string     `json:"avatar_url"`
	Bio       string     `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
}

func (r createReq) toInput() users.CreateInput {
	return users.CreateInput{
		Username:  r.Username,
		Phone:     r.Phone,
		Password:  r.Password,
		AvatarUrl: r.AvatarUrl,
		Bio:       r.Bio,
		Birthday:  r.Birthday,
	}
}

func (r createReq) validate() error {
	if r.Username == "" || r.Phone == "" || r.Password == "" {
		return errWrongBody
	}
	return nil
}

type updateUserReq struct {
	ID        string     `json:"id" binding:"required"`
	Password  *string    `json:"password"`
	AvatarUrl *string    `json:"avatar_url"`
	Bio       *string    `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
}

func (r updateUserReq) toInput() users.UpdateInput {
	return users.UpdateInput{
		ID:        r.ID,
		Password:  r.Password,
		AvatarUrl: r.AvatarUrl,
		Bio:       r.Bio,
		Birthday:  r.Birthday,
	}
}

func (r updateUserReq) validate() error {
	if uuid.Validate(r.ID) != nil {
		return errWrongQuery
	}

	if r.Password != nil && *r.Password == "" {
		return errWrongBody
	}
	return nil
}

type listUsersReq struct {
	ID       string `form:"id"`
	Username string `form:"username"`
	Phone    string `form:"phone"`
}

func (r listUsersReq) toInput() users.ListInput {
	return users.ListInput{
		Filter: users.Filter{
			ID:       r.ID,
			Username: r.Username,
			Phone:    r.Phone,
		},
	}
}

func (r listUsersReq) validate() error {
	if r.ID != "" && uuid.Validate(r.ID) != nil {
		return errWrongQuery
	}
	return nil
}

type usesResp struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Phone     string     `json:"phone"`
	AvatarUrl string     `json:"avatar_url"`
	Bio       string     `json:"bio"`
	Birthday  *time.Time `json:"birthday"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (h handler) newUserResp(u domain.User) usesResp {
	return usesResp{
		ID:        u.ID.String(),
		Username:  u.Username,
		Phone:     u.Phone,
		AvatarUrl: u.AvatarUrl,
		Bio:       u.Bio,
		Birthday:  u.Birthday,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (h handler) newListUserResp(us []domain.User) []usesResp {
	items := make([]usesResp, 0, len(us))

	for _, u := range us {
		item := usesResp{
			ID:        u.ID.String(),
			Username:  u.Username,
			Phone:     u.Phone,
			AvatarUrl: u.AvatarUrl,
			Bio:       u.Bio,
			Birthday:  u.Birthday,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		}

		items = append(items, item)
	}

	return items
}
