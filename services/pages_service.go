package services

var PagesService pagesServiceInterface = &pagesService{}

type pagesServiceInterface interface{}

type pagesService struct{}
