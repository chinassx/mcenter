package impl

import (
	"context"
	"fmt"

	"github.com/infraboard/mcenter/apps/token"
	"github.com/infraboard/mcube/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) save(ctx context.Context, tk *token.Token) error {
	if _, err := s.col.InsertOne(ctx, tk); err != nil {
		return exception.NewInternalServerError("inserted token(%s) document error, %s",
			tk.AccessToken, err)
	}

	return nil
}

func (s *service) get(ctx context.Context, id string) (*token.Token, error) {
	filter := bson.M{"_id": id}

	ins := token.NewToken(token.NewIssueTokenRequest())
	if err := s.col.FindOne(ctx, filter).Decode(ins); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, exception.NewNotFound("token %s not found", id)
		}

		return nil, exception.NewInternalServerError("find token %s error, %s", id, err)
	}

	return ins, nil
}

func (s *service) update(tk *token.Token) error {
	_, err := s.col.UpdateOne(context.TODO(), bson.M{"_id": tk.AccessToken}, bson.M{"$set": tk})
	if err != nil {
		return exception.NewInternalServerError("update token(%s) error, %s", tk.AccessToken, err)
	}

	return nil
}

func (s *service) delete(ctx context.Context, ins *token.Token) error {
	if ins == nil || ins.AccessToken == "" {
		return fmt.Errorf("access tpken is nil")
	}

	result, err := s.col.DeleteOne(ctx, bson.M{"_id": ins.AccessToken})
	if err != nil {
		return exception.NewInternalServerError("delete token(%s) error, %s", ins.AccessToken, err)
	}

	if result.DeletedCount == 0 {
		return exception.NewNotFound("book %s not found", ins.AccessToken)
	}

	return nil
}

func newQueryRequest(req *token.QueryTokenRequest) *queryRequest {
	return &queryRequest{req}
}

type queryRequest struct {
	*token.QueryTokenRequest
}

func (r *queryRequest) FindOptions() *options.FindOptions {
	pageSize := int64(r.Page.PageSize)
	skip := int64(r.Page.PageSize) * int64(r.Page.PageNumber-1)

	opt := &options.FindOptions{
		Sort:  bson.D{{Key: "create_at", Value: -1}},
		Limit: &pageSize,
		Skip:  &skip,
	}

	return opt
}

func (r *queryRequest) FindFilter() bson.M {
	filter := bson.M{}
	if r.GrantType != nil {
		filter["grant_type"] = r.GrantType
	}
	if r.Username != "" {
		filter["username"] = r.Username
	}
	return filter
}
