//+build wireinject

package main

import "github.com/google/wire"

func GetPostService() *PostService {
	panic(wire.Build(
		NewPostService,
		NewPostUsecase,
		NewPostRepo,
	))
}