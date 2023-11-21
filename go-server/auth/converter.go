package auth

import (
	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/db/model"
	pb "github.com/RakshitKumar04/ChatMate/tree/master/go-server/pb"
)

func ConvertUserObjectToUser(model *model.UserModel) *pb.User {
	return &pb.User{
		Username: model.Username,
		Name:     model.Name,
		Id:       int32(model.ID.Timestamp().Day()),
	}
}
