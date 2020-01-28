import (
	"database/sql"
	"errors"

)

// ReviewRepoImpl implements the review.ReviewRepository interface
type ReviewRepoImpl struct {
	conn *sql.DB
}

// NewReviewRepoImpl will create an object of PsqlReviewRepository
func NewReviewRepoImpl(Con *sql.DB) *ReviewRepoImpl {
	return &ReviewRepoImpl{conn: Con}
}
// comments returns all Reviews from the database
func (rri *ReviewRepoImpl) comments() ([]entity.Comment, error) {
	query := "SELECT * FROM comment"
	rows, err := rri.conn.Query(query)
	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	rvws := []entity.Comment{}

	for rows.Next() {
		review := entity.Comment{}
		err = rows.Scan(&comment.ID, &comment.Rating, &comment.UserID, &comment.Message, &comment.ReviewedAt)

		if err != nil {
			return nil, err
		}
		rvws = append(rvws, review)
	}

	return rvws, nil
}

// comment returns the reviews of a single book
func (rri *ReviewRepoImpl) comment(id int) (entity.Comment, error) {

	query := "SELECT * FROM comment WHERE id = $1"
	row := rri.conn.QueryRow(query, id)

	r := entity.Comment{}

	err := row.Scan(&r.ID, &r.UserName, &r.Massage, &r.User_id, &r.Book_id, &r.ReviewedAt)
	if err != nil {
		return r, err
	}

	return r, nil
}
// MakeReview stores new review information to database
func (rri *ReviewRepoImpl) MakeReview(r entity.Comment) error {

	query := "INSERT INTO comment (Username,ID,rating,Book_id,User_id,message,ReviewedAt) values($1, $2, $3, $4)"
	_, err := rri.conn.Exec(query,r.UserName, r.ID, r.rating, r.Book_id, r.User_id, r.Message,r.Email,r.ReviewedAt)
	if err != nil {
		return errors.New("reviewing has failed")
	}

	return nil
}

// GetMyReviews returns the reviews of a single user
func (rri *ReviewRepoImpl) GetMyReviews(id int) ([]entity.Comment, error) {

	query := "SELECT * FROM comment WHERE User_id = $1"
	rows, err := rri.conn.Query(query, id)

	if err != nil {
		return nil, errors.New("Could not query the database")
	}
	defer rows.Close()

	rvws := []entity.Comment{}

	for rows.Next() {
		review := entity.Comment{}
		err = rows.Scan(&review.ID, &review.Rating, &review.ReviewedAt, &review.User_id, &review.Book_id, &review.Message)
		if err != nil {
			return nil, err
		}

		rvws = append(rvws, review)
	}

	return rvws, nil
}

// UpdateReview updates a given object with a new data
func (rri *ReviewRepoImpl) UpdateReview(r entity.Comment) error {
	query := "UPDATE review SET rating=$1,event_id=$2, user_id=$3, message=$4 WHERE id=$5"
	_, err := rri.conn.Exec(query, r.Rating, r.Book_id, r.User_id, r.Message, r.ID,r.Username,r.ReviewedAt,r.Email)

	if err != nil {
		return errors.New("updating has failed")
	}

	return nil
}

// DeleteReview removes a review from a database by its id
func (rri *ReviewRepoImpl) DeleteReview(id int) error {
	query := "DELETE FROM comment WHERE id=$1"
	_, err := rri.conn.Exec(query, id)

	if err != nil {
		return errors.New("Delete has failed")
	}

	return nil
}

//SetRating sets the average rating of an event after every reviews
func (rri *ReviewRepoImpl) SetRating(Eid int) error {

	query := "SELECT AVG(rating) FROM comment WHERE Book_id = $1"
	row := rri.conn.QueryRow(query, Eid)

	var rating float32
	err := row.Scan(&rating)

	if err != nil {
		return errors.New("Could not make average in the database")
	}
	_, er := rri.conn.Exec("UPDATE events SET rating=$1 WHERE id=$1", rating, Eid)

	if er != nil {
		return errors.New("setting new rating has failed")
	}
	return nil
}



