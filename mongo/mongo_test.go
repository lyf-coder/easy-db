/*
 * Copyright © 2020 - present. liyongfei <liyongfei@walktotop.com>.
 *
 * Use of this source code is governed by an MIT-style
 * license that can be found in the LICENSE file.
 */

package mongo

import (
	"github.com/lyf-coder/easy-db/connect"
	"testing"
)

// getUri Test
func Test_getUri(t *testing.T) {
	// create Mongodb instance
	config := new(connect.Config)
	config.UserName = "userName"
	config.Password = "password"
	config.DatabaseName = "databaseName"
	config.Host = "127.0.0.1"
	config.Port = "27017"

	uri := getUri(config)
	expectVal := "mongodb://userName:password@127.0.0.1:27017/databaseName?"
	if uri != expectVal {
		t.Errorf("createUri error !")
		t.Errorf("expect value is \"%v\"", expectVal)
		t.Errorf("actual value is \"%v\"", uri)
	}
}
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
var mongodb = New(&config)

// New Test
func TestNew(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	err := mongodb.client.Ping(ctx, readpref.Primary())
	if err != nil{
		t.Error(err)
	}
}


func TestMongodb_Finds(t *testing.T) {
	filter := make(map[string]interface{})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	results,err := mongodb.Finds(ctx,"GO_TEST", filter, options.FindOpts{})
	if err!=nil{
		t.Error(err)
	}
	for _,val := range results{
		log.Println(val.Get("name"))
	}

}

func TestMongodb_Updates(t *testing.T) {
	filter := make(map[string]interface{})
	filter["name"] = "zhangsan"
	//update := bson.D{
	//	{"$set", bson.D{
	//		{"value", "777"},
	//	}},
	//}
	newUpdate := entity.Entity{}
	newUpdate.Set("$set:value","999")

	updateOpts := options.UpdateOpts{}
	updateOpts.SetUpsert(true)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	result,err := mongodb.UpdatesWithOptions(ctx,"GO_TEST", filter, newUpdate.GetData(), updateOpts)
	if err!=nil{
		t.Error(err)
		t.Error(result)
	}
	//if err!=nil{
	//	t.Error(err)
	//}
	//for _,val := range results{
	//	log.Println(val.Get("name"))
	//}
}

func TestMongodb_Deletes(t *testing.T) {
	defer func() {
		if err := mongodb.Close(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	filter := make(map[string]interface{})
	filter["name"] = "pi_test2"


	ctx, cancelFunc := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelFunc()
	//result,err := mongodb.Deletes("GO_TEST", filter)
	result,err := mongodb.Deletes(ctx,"GO_TEST", filter)
	if err!=nil{
		t.Error(err)
	}
	log.Println(result.DeletedCount)
	//if err!=nil{
	//	t.Error(err)
	//}
	//for _,val := range results{
	//	log.Println(val.Get("name"))
	//}
}


 */
