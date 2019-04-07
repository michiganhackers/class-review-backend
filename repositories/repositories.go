package repositories

type Repositories struct {
    ReviewRepository IReviewRepository
    CourseRepository ICourseRepository

    // TODO: put database engine here
}

func DefaultRepositories() *Repositories {
    return &Repositories{
        ReviewRepository: DefaultReviewRepository(),
        CourseRepository: DefaultCourseRepository(),
    }
}
