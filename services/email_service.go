package services

var EmailService emailServiceInterface = &emailService{}

type emailServiceInterface interface{}

type emailService struct{}
