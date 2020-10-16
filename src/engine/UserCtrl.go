package engine

import (
	"context"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NgeKaworu/time-mgt-go/src/models"
	"github.com/NgeKaworu/time-mgt-go/src/parsup"
	"github.com/NgeKaworu/time-mgt-go/src/resultor"
	"github.com/NgeKaworu/time-mgt-go/src/utils"
	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson"
)

// Login 登录
func (d *DbEngine) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}
	if len(body) == 0 {
		resultor.RetFail(w, "not has body")
		return
	}

	p, err := parsup.ParSup().ConvJSON(body)
	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	err = utils.Required(p, map[string]string{
		"pwd":   "密码不能为空",
		"email": "邮箱不能为空",
	})

	t := d.GetColl(models.TUser)

	email := strings.ToLower(strings.Replace(p["email"].(string), " ", "", -1))
	res := t.FindOne(context.Background(), bson.M{
		"email": email,
	})

	if res.Err() != nil {
		resultor.RetFail(w, "没有此用户")
		return
	}

	var u models.User

	err = res.Decode(&u)
	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	dec, err := d.Auth.CFBDecrypter(*u.Pwd)
	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	if string(dec) != p["pwd"] {
		resultor.RetFail(w, "用户名密码不匹配，请注意大小写。")
		return
	}

	tk, err := d.Auth.GenJWT(u.ID.Hex())

	if err != nil {
		resultor.RetFail(w, err.Error())
		return
	}

	resultor.RetOk(w, tk)
	return
}
