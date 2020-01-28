package review

import "github.com/goEventListing/API/entity"

// ReviewService specifies application review related services
type ReviewService interface {
	Reviews() ([]entity.Comment, []error)

	Review(id uint) (*entity.Comment, []error)
	MakeReviewAndRating(r *entity.Comment) (*entity.Comment, []error)

	UpdateReview(r *entity.Comment) (*entity.Comment, []error)
	DeleteReview(id uint) (*entity.Comment, []error)
	//EventReviews(id uint) ([]entity.Comment, []error)
	//GetMyReviews(id uint) ([]entity.Review, []error)
	// getMyRating(UID, EventID int) int
	// deleteComment(id int) error
	// getcomments(eventID int) []entity.Comment// reviews with no comments
	// justrate(EventID, UserID, rating int) error	Reviews() ([]entity.Review, []error)

}
