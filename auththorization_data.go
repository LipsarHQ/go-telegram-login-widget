package telegramloginwidget

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/url"
	"strconv"
	"strings"
)

type AuthorizationData struct {
	FirstName string `json:"first_name,omitempty"`
	Hash      string `json:"hash,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	PhotoURL  string `json:"photo_url,omitempty"`
	Username  string `json:"username,omitempty"`
	AuthDate  int64  `json:"auth_date,omitempty"`
	ID        int64  `json:"id,omitempty"`
}

func (x *AuthorizationData) Check(token string) (err error) {
	h := x.Sum(token)
	if subtle.ConstantTimeCompare([]byte(x.Hash), []byte(h)) != 1 {
		return ErrHashInvalid
	}

	return nil
}

func (x *AuthorizationData) Sum(token string) string {
	return hex.EncodeToString(hashHMAC([]byte(x.String()), hashSHA256([]byte(token)), sha256.New))
}

func (x *AuthorizationData) String() string {
	vs := make([]string, 0, totalFields)
	if x.AuthDate != 0 {
		vs = append(vs, "auth_date="+strconv.FormatInt(x.AuthDate, 10))
	}

	if len(x.FirstName) != 0 {
		vs = append(vs, "first_name="+x.FirstName)
	}

	if x.ID != 0 {
		vs = append(vs, "id="+strconv.FormatInt(x.ID, 10))
	}

	if len(x.LastName) != 0 {
		vs = append(vs, "last_name="+x.LastName)
	}

	if len(x.PhotoURL) != 0 {
		vs = append(vs, "photo_url="+x.PhotoURL)
	}

	if len(x.Username) != 0 {
		vs = append(vs, "username="+x.Username)
	}

	return strings.Join(vs, "\n")
}

func NewFromQuery(values url.Values) (modelAuthorizationData *AuthorizationData, err error) {
	modelAuthorizationData = &AuthorizationData{
		FirstName: values.Get("first_name"),
		Hash:      values.Get("hash"),
		LastName:  values.Get("last_name"),
		PhotoURL:  values.Get("photo_url"),
		Username:  values.Get("username"),
		AuthDate:  0,
		ID:        0,
	}
	if len(modelAuthorizationData.Hash) == 0 {
		return nil, ErrHashInvalid
	}

	//nolint:errcheck // No need to check error.
	modelAuthorizationData.AuthDate, _ = strconv.ParseInt(values.Get("auth_date"), 10, 64)

	//nolint:errcheck // No need to check error.
	modelAuthorizationData.ID, _ = strconv.ParseInt(values.Get("id"), 10, 64)

	return modelAuthorizationData, nil
}

func NewFromURI(uri string) (modelAuthorizationData *AuthorizationData, err error) {
	var u *url.URL

	if u, err = url.ParseRequestURI(uri); err != nil {
		return nil, err
	}

	return NewFromQuery(u.Query())
}
