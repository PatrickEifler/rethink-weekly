package main

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"log"
	"os"
)

type notifier interface {
	NotifiySubscriber(*Subscriber) (bool, error)
	ApproveSubscriber(*Subscriber) (bool, error)
	UnSubscribe(*Subscriber) (bool, error)
	SendNewsletter() (int, error)
}

type mailer struct{}

func (*mailer) NotifiySubscriber(subscriber *Subscriber) (bool, error) {
	mg := mailgun.NewMailgun("mg.noty.im", os.Getenv("MAILGUN_API"), "")

	m := mg.NewMessage(
		fmt.Sprintf("Vinh <%s>", os.Getenv("MAIL_FROM")), // From
		"Verify subscriptions on RethinkDB Goodies",      // Subject
		fmt.Sprintf("Hi please follow this link http://127.0.0.1:3000/api/subscriptions/%s/confirm to verify. You can ignore if you don't want to subscribe.", subscriber.ConfirmToken), // Plain-text body
		fmt.Sprintf("%s <%s>", subscriber.FirstName, subscriber.Email),                                                                                                                  // Recipients (vararg list)
	)

	_, _, err := mg.Send(m)

	if err != nil {
		log.Panic(err)
		return false, err
	}
	return true, nil
}

func (*mailer) ApproveSubscriber(subscriber *Subscriber) (bool, error) {
	mg := mailgun.NewMailgun("mg.noty.im", os.Getenv("MAILGUN_API"), "")

	m := mg.NewMessage(
		fmt.Sprintf("Vinh <%s>", os.Getenv("MAIL_FROM")),                                 // From
		"Thankyou for subscriptions on RethinkDB Goodies",                                // Subject
		fmt.Sprintf("Cool %s, we will send you weekly email now.", subscriber.FirstName), // Plain-text body
		fmt.Sprintf("%s <%s>", subscriber.FirstName, subscriber.Email),                   // From
	)

	_, _, err := mg.Send(m)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (*mailer) UnSubscribe(subscriber *Subscriber) (bool, error) {
	mg := mailgun.NewMailgun("mg.noty.im", os.Getenv("MAILGUN_API"), "")

	m := mg.NewMessage(
		fmt.Sprintf("Vinh <%s>", os.Getenv("MAIL_FROM")), // From
		"Sorry seeing you go",                            // Subject
		fmt.Sprintf("This is the last email we send you to let you know that we removed you from the list, and you will be no longer receive any email from us."), // Plain-text body
		fmt.Sprintf("%s <%s>", subscriber.FirstName, subscriber.Email),                                                                                            // From
	)

	_, _, err := mg.Send(m)

	if err != nil {
		return false, err
	}
	return true, nil
}

// Send out news letter on due date
func (*mailer) SendNewsletter() (int, error) {
	return 0, nil
}
