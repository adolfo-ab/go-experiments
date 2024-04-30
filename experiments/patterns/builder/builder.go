package main

import "fmt"

type NotificationBuilder struct {
	Title    string
	Subtitle string
	Message  string
	Image    string
	Icon     string
	Priority int
	NotType  string
}

func newNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{}
}

func (nb *NotificationBuilder) SetTitle(title string) {
	nb.Title = title
}

func (nb *NotificationBuilder) SetSubtitle(subtitle string) {
	nb.Subtitle = subtitle
}

func (nb *NotificationBuilder) SetMessage(message string) {
	nb.Message = message
}

func (nb *NotificationBuilder) SetImage(image string) {
	nb.Image = image
}

func (nb *NotificationBuilder) SetIcon(icon string) {
	nb.Icon = icon
}

func (nb *NotificationBuilder) SetPriority(priority int) {
	nb.Priority = priority
}

func (nb *NotificationBuilder) SetNotType(notType string) {
	nb.NotType = notType
}

func (nb *NotificationBuilder) Build() (*Notification, error) {
	if nb.Title == "" || nb.Message == "" {
		return nil, fmt.Errorf("notifications must have a title and a message")
	}

	if nb.Priority < 0 || nb.Priority > 99 {
		return nil, fmt.Errorf("priority %d is out of range, must be in the range 0-99", nb.Priority)
	}

	return &Notification{
		title:    nb.Title,
		subtitle: nb.Subtitle,
		message:  nb.Message,
		image:    nb.Image,
		icon:     nb.Icon,
		priority: nb.Priority,
		notType:  nb.NotType,
	}, nil
}
