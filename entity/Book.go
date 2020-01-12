package entity

//Book -
type Book struct {
	ID          int
	Name        string
	Owner       int //AuthorID from the Author information
	PricePerDay float64
	OnDiscount  bool
	Discount    float32 //in percent
	OnSale      bool    // is the book open for sale
	ImagePath   string  // book image path
}
