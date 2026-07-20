package mongocache

import (
	"context"
	"fmt"
	"github.com/xpwu/go-db-mongo/mongodb/filter"
	"github.com/xpwu/go-db-mongo/mongodb/index"
	"github.com/xpwu/go-db-mongo/mongodb/tagparser"
	"github.com/xpwu/go-db-mongo/mongodb/updater"
	"github.com/xpwu/go-log/log"
	"github.com/xpwu/go-x/exe"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/v2/bson/bsonrw"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"reflect"
	"sync"
	"time"
)

var (
	clients = sync.Map{}
)

type Client struct {
	*mongo.Client
	maxConn uint64
	id      string
}

func (c *Client) Disconnect(ctx context.Context) error {
	// nothing to do
	return nil
}

// 一个client可能在很多地方使用，调用此函数时，所有基于此client的请求都将失败
func (c *Client) DisconnectForce(ctx context.Context) error {
	clients.Delete(c.id)

	return c.Client.Disconnect(ctx)
}

func newClient(config *Config) (*Client, error) {
	client, err := mongo.NewClient(
		// 根据底层逻辑，顺序很重要，取的都是最后设置的值
		// 默认值先设置
		options.Client().
			SetAppName(exe.Name).
			SetConnectTimeout(5 * time.Second).
			SetServerSelectionTimeout(5 * time.Second).
			SetSocketTimeout(5 * time.Second).
			SetRetryWrites(true).
			SetMaxConnIdleTime(3 * time.Minute).
			SetReadPreference(readpref.SecondaryPreferred()).
			SetMinPoolSize(1).

			// 然后是URI 携带的参数，最后是配置中明确指明的设置
			ApplyURI(config.URI).
			SetMaxPoolSize(config.MaxConn).
			SetAuth(options.Credential{
				Username:    config.User,
				Password:    config.Password,
				PasswordSet: true,
			}).
			SetRegistry(defaultRegister),
	)

	if err != nil {
		return nil, err
	}

	return &Client{
		Client:  client,
		id:      config.id(),
		maxConn: config.MaxConn,
	}, nil
}

func MustGet(ctx context.Context, config Config) *Client {
	r, err := Get(ctx, config)
	if err != nil {
		panic(err)
	}

	return r
}

func Get(ctx context.Context, config Config) (ret *Client, err error) {
	_, logger := log.WithCtx(ctx)
	id := config.id()

	c, ok := clients.Load(id)
	if ok {
		ret = c.(*Client)
		if ret.maxConn < config.MaxConn {
			logger.Warning(fmt.Sprintf("return cache Client<%s>, whose maxCon(%d) is less than expected(%d)",
				config.URI, ret.maxConn, config.MaxConn))
		}
		return
	}

	nc, err := newClient(&config)
	if err != nil {
		return
	}

	c, ok = clients.LoadOrStore(id, nc)
	ret = c.(*Client)
	// new client
	if !ok {
		// 连接使用新的ctx
		err = ret.Connect(context.Background())
	}
	if err != nil {
		return
	}

	return
}

var structCodec, _ = bsoncodec.NewStructCodec(
	bsoncodec.StructTagParserFunc(tagparser.StructTagParser))

var defaultRegister = bson.NewRegistryBuilder().
	RegisterDefaultDecoder(reflect.Struct, structCodec).
	RegisterDefaultEncoder(reflect.Struct, structCodec).
	RegisterTypeEncoder(reflect.TypeOf((*updater.Updater)(nil)).Elem(),
		bsoncodec.ValueEncoderFunc(func(ec bsoncodec.EncodeContext,
			vw bsonrw.ValueWriter, val reflect.Value) error {

			v := val.Interface().(updater.Updater).ToBsonM()
			enc, err := ec.LookupEncoder(reflect.TypeOf(v))
			if err != nil {
				return err
			}

			return enc.EncodeValue(ec, vw, reflect.ValueOf(v))
		})).
	RegisterTypeEncoder(reflect.TypeOf((*filter.Filter)(nil)).Elem(),
		bsoncodec.ValueEncoderFunc(func(ec bsoncodec.EncodeContext,
			vw bsonrw.ValueWriter, val reflect.Value) error {

			v := val.Interface().(filter.Filter).ToBsonD()
			enc, err := ec.LookupEncoder(reflect.TypeOf(v))
			if err != nil {
				return err
			}

			return enc.EncodeValue(ec, vw, reflect.ValueOf(v))
		})).
	RegisterTypeEncoder(reflect.TypeOf((*index.Key)(nil)).Elem(),
		bsoncodec.ValueEncoderFunc(func(ec bsoncodec.EncodeContext,
			vw bsonrw.ValueWriter, val reflect.Value) error {

			v := val.Interface().(index.Key).ToBsonD()
			enc, err := ec.LookupEncoder(reflect.TypeOf(v))
			if err != nil {
				return err
			}

			return enc.EncodeValue(ec, vw, reflect.ValueOf(v))
		})).
	Build()
