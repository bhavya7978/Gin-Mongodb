package service

import (
	"context"
	"errors"
	"project_mongodb-go/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserServices {
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}
func (u *UserServiceImpl) CreateUser(user *entity.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserServiceImpl) GetUser(name *string) (*entity.User, error) {
	var user *entity.User
	filter := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, filter).Decode(&user)
	return user, err
}

func (u *UserServiceImpl) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	cursor, err := u.usercollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(u.ctx) {
		var user entity.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)
	if len(users) == 0 {
		return nil, errors.New("No document in database")
	}
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *entity.User) error {
	filter := bson.D{bson.E{Key: "name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "name", Value: user.Name}, bson.E{Key: "age", Value: user.Age}, bson.E{Key: "address", Value: user.Address}}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("No matched document for the update.")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "name", Value: name}}
	result, _ := u.usercollection.DeleteOne(u.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("No matched document for the update.")
	}
	return nil
}
