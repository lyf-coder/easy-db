/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package connect // import "github.com/lyf-coder/easy-db/connect"

// MONGODB
const MONGODB = "MONGODB"

// Config is connect database config struct
type Config struct {
	DbType       string
	UserName     string
	Password     string
	DatabaseName string
	Host         string
	Port         string
	Options      map[string]string
}
