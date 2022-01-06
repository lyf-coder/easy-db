/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package db

/** if you need test the following func, you need modify config value first,
and then delete this annotation block

var config = connect.Config{
	DbType:       "DbType",
	UserName:     "UserName",
	Password:     "Password",
	DatabaseName: "DatabaseName",
	Host:         "ip",
	Port:         "27017",
	Options:      nil,
}

var db = New(&config)

func Test_New(t *testing.T) {
	const COLLECTION = "GO_TEST"
	// test Insert
	doc := map[string]string{}
	doc["name"] = "jack"
	doc["say"] = "hi!rose!"
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelFunc()
	id, err := db.Insert(ctx, COLLECTION, doc, nil)
	if err != nil {
		t.Error(err)
		t.Error(id)
	}

	// test Finds
	filter := make(map[string]interface{})

	results, err := db.Finds(ctx, COLLECTION, filter, nil)
	if err != nil {
		t.Error(err)
	}
	for _, val := range results {
		log.Println(val.Get("name"))
	}

	deleteFilter := map[string]string{}
	deleteFilter["name"] = "pi_test"

	result, err := db.Delete(ctx, COLLECTION, deleteFilter, nil)
	if err != nil {
		t.Error(err)
		t.Error(result)
	}
}
*/
