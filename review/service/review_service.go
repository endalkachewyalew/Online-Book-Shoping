package services

import (
	
	"github.com/goEventListing/API/entity"
	"github.com/goEventListing/API/review"
)

// ReviweServiceImpl implements review.Reviewservice interface
type ReviweServiceImpl struct {
	ReviewRepo review.ReviewRepository
}

// NewReviewServiceImpl will create new CategoryService object
func NewReviewServiceImpl(RevRepo comment.ReviewRepository) *ReviweServiceImpl {
	return &ReviweServiceImpl{ReviewRepo: RevRepo}
}

// Reviews returns list of Reviews
func (rs *ReviweServiceImpl) Reviews() ([]entity.Comment, []error) {

	rvs, errs := rs.ReviewRepo.Reviews()

	if len(errs) > 0 {
		return nil, errs
	}
	return rvs, nil
}

// Review returns a Review with a given id
func (rs *ReviweServiceImpl) Review(id uint) (*entity.Comment, []error) {

	rvw, errs := rs.ReviewRepo.Comment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return rvw, nil

}

// UpdateReview updates a given review with a new data
func (rs *ReviweServiceImpl) UpdateReview(r *entity.Review) (*entity.Comment, []error) {

	revw, err := rs.ReviewRepo.UpdateReview(r)

	if len(err) > 0 {
		return nil, err
	}

	return revw, nil
}

// DeleteReview removes a review by its id
func (rs *ReviweServiceImpl) DeleteReview(id uint) (*entity.Comment, []error) {

	rvws, errs := rs.ReviewRepo.DeleteReview(id)

	if len(errs) > 0 {
		return nil, errs
	}

	return rvws, nil
}

//MakeReviewAndRating stores new review information
func (rs *ReviweServiceImpl) MakeReviewAndRating(r *entity.Comment) (*entity.Comment, []error) {

	rvws, errs := rs.ReviewRepo.MakeReviewAndRating(r)

	if len(errs) > 0 {
		return nil, errs
	}
	// er := rs.ReviewRepo.SetRating(r.EventID)

	// if er != nil {
	// 	return er
	// }

	return rvws, nil
}

// // GetMyReviews returns the reviews of a single user
// func (rs *ReviweServiceImpl) GetMyReviews(id uint) ([]entity.Review, []error) {
// 	fmt.Println("calling repo-----------")
// 	revi, errs := rs.ReviewRepo.GetMyReviews(id)
// 	fmt.Println("called repo-----------------", revi)

// 	if len(errs) > 0 {
// 		return nil, errs
// 	}
// 	fmt.Println("userRev-----------------", revi, errs)
// 	return revi, nil
// }

//EventReviews ...
func (rs *ReviweServiceImpl) EventReviews(id uint) ([]entity.Review, []error) {

	rvs, errs := rs.ReviewRepo.EventReviews(id)

	if len(errs) > 0 {
		return nil, errs
	}
	//fmt.Println("eventRev-----------------", rvs)
	return rvs, nil
}

