/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package options // import "github.com/lyf-coder/easy-db/options"
import (
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOneOpts
type InsertOneOpts struct {
	options.InsertOneOptions
}

// InsertOpts represents all possible options to the InsertWithOptions and InsertsWithOptions function.
type InsertOpts struct {
	options.InsertManyOptions
}

// FindOpts represent all possible options to the FindWithOptions and FindsWithOptions function.
type FindOpts struct {
	options.FindOptions
}

// UpdateOpts represents all possible options to the UpdateWithOptions() and UpdatesWithOptions() functions.
type UpdateOpts struct {
	options.UpdateOptions
}

// DeleteOpts represents all possible options to the DeleteWithOptions() and DeletesWithOptions() functions.
type DeleteOpts struct {
	options.DeleteOptions
}

// CountOpts represents all possible options to the CountWithOptions() function.
type CountOpts struct {
	options.CountOptions
}
