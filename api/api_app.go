// Copyright © 2016 Thomas Rabaix <thomas.rabaix@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package api

import (
	"github.com/rande/goapp"
	"github.com/rande/pkgmirror"
	"goji.io"
	"goji.io/pat"
)

func ConfigureApp(config *pkgmirror.Config, l *goapp.Lifecycle) {

	l.Prepare(func(app *goapp.App) error {
		mux := app.Get("mux").(*goji.Mux)
		mux.HandleFuncC(pat.Get("/api/mirrors"), Api_GET_MirrorServices(app))

		return nil
	})

}
