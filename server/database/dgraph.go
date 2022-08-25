package database

import (
	"context"
	"log"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
)

type DgraphDB struct {
	dGraphClient *dgo.Dgraph
}

var DgraphClient DgraphDB

func NewDatabaseClient(conn *grpc.ClientConn) *DgraphDB {
	db := new(DgraphDB)
	db.dGraphClient = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return db
}

const query = `{
	programs(func: has(code))
	{
		uid,
		name,
		code,
	}
}`
const getById = `query getById($uid: string) {
	getById(func: uid($uid)) {
	  uid,
	  name,
	  code
	}
  }`

func (resp DgraphDB) GetPrograms(ctx context.Context) (*api.Response, error) {
	txn := resp.dGraphClient.NewTxn()
	defer txn.Commit(ctx)
	res, err := txn.Query(ctx, query)
	if err != nil {
		log.Println(err)
	}
	return res, err
}

func (resp DgraphDB) GetWithId(uid string, ctx context.Context) (*api.Response, error) {
	txn := resp.dGraphClient.NewTxn()
	defer txn.Commit(ctx)
	res, err := txn.QueryWithVars(ctx, getById, map[string]string{"$uid": uid})
	if err != nil {
		log.Println(err)
	}
	return res, err
}

func (resp DgraphDB) Add(json []byte, ctx context.Context) (*api.Response, error) {
	txn := resp.dGraphClient.NewTxn()
	defer txn.Commit(ctx)
	mutation := &api.Mutation{
		SetJson: json,
	}
	res, err := txn.Mutate(ctx, mutation)
	return res, err

}

func (resp DgraphDB) Delete(json []byte, ctx context.Context) (*api.Response, error) {
	txn := resp.dGraphClient.NewTxn()
	defer txn.Commit(ctx)
	mutation := &api.Mutation{
		DeleteJson: json,
	}
	res, err := txn.Mutate(ctx, mutation)
	return res, err
}
