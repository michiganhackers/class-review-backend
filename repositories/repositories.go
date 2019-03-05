package repositories

type Repositories struct {
	ReviewRepository IReviewRepository

	// TODO: put database engine here
}

func DefaultRepositories() *Repositories {
	return &Repositories{
		ReviewRepository: DefaultReviewRepository(),
	}
}
