package main
type IPostRepo interface {}

func NewPostRepo() IPostRepo  {
	return new(IPostRepo)
}

type IPostUsecase interface {}
type postUsecase struct {
	repo IPostRepo
}

func NewPostUsecase(repo IPostRepo) IPostUsecase {
	return postUsecase{repo: repo}
}

type PostService struct {
	usecase IPostUsecase
}

func NewPostService(u IPostUsecase) *PostService  {
	return &PostService{usecase: u}
}