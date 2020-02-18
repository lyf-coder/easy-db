/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package result // import "github.com/lyf-coder/easy-db/result"
import mongoDriver "go.mongodb.org/mongo-driver/mongo"

// UpdateResult Db Update result struct
type UpdateResult struct {
	mongoDriver.UpdateResult
}

// DeleteResult Db Delete result struct
type DeleteResult struct {
	mongoDriver.DeleteResult
}
