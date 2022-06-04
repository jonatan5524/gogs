// Copyright 2014 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"net/http"

	api "github.com/gogs/go-gogs-client"

	"gogs.io/gogs/internal/context"
	"gogs.io/gogs/internal/db"
)

func CreateAccessToken(c *context.APIContext, form api.CreateAccessTokenOption) {
	t, err := db.AccessTokens.Create(c.User.ID, form.Name)
	if err != nil {
		if db.IsErrAccessTokenAlreadyExist(err) {
			c.ErrorStatus(http.StatusUnprocessableEntity, err)
		} else {
			c.Error(err, "new access token")
		}
		return
	}
	c.JSON(http.StatusCreated, &api.AccessToken{Name: t.Name, Sha1: t.Sha1})
}
