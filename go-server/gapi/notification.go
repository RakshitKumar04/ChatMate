package gapi

import (
	"fmt"

	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/pb"
	"github.com/RakshitKumar04/ChatMate/tree/master/go-server/services"
)

// GetNotifications is a gRPC service method that streams notifications to the client.
func (server *Server) GetNotifications(req *pb.EmptyRequest, stream pb.GrpcServerService_GetNotificationsServer) error {
	fmt.Println("Notification service started")

	// Create a channel for receiving notifications
	notificationCh := make(chan *pb.NotificationMessage)

	// Start the background service to listen for new user additions
	go services.NotificationNewUser(server.dbCollection.Users, notificationCh)

	fmt.Println("Notification service created")

	// Continuously listen for events
	for {
		select {
		case <-stream.Context().Done():
			// Client disconnected, close the channel and exit
			close(notificationCh)
			return nil
		case notification, ok := <-notificationCh:
			if !ok {
				// The channel has been closed, exit the loop
				return nil
			}
			// Send the received notification to the client
			if err := stream.Send(&pb.NotificationMessage{
				Title:       notification.Title,
				Id:          int32(notification.Id),
				Description: notification.Description,
			}); err != nil {
				return err
			}
		}
	}
}