package paypalauth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

// Token は、アクセス トークン。
type Token struct {
	Client      *Client
	OAuth2Token *oauth2.Token
}

// UserInfo は、ユーザ情報を取得する。
func (t *Token) UserInfo(ctx context.Context, result interface{}) error {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/identity/oauth2/userinfo?schema=paypalv1.1", t.Client.Endpoint.RootURL), nil)
	if err != nil {
		return err
	}
	t.OAuth2Token.SetAuthHeader(req)

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, result); err != nil {
		return err
	}

	return nil
}
