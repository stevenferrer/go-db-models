package user

// represents an application user
type User struct {
	ID         int64
	Email      string
	Password   string
	FirstName  string
	MiddleName string
	LastName   string
}

//Filter is used for filtering results
type Filter func(*FilterConfig) error

//FilterConfig is the config for filtering results
type FilterConfig struct {
	Limit      int64
	ID         int64
	Email      string
	FirstName  string
	MiddleName string
	LastName   string
}

func LimitFilter(limit int64) Filter {
	return func(fc *FilterConfig) error {
		fc.Limit = limit
		return nil
	}
}

func IDFilter(id int64) Filter {
	return func(fc *FilterConfig) error {
		fc.ID = id
		return nil
	}
}

func EmailFilter(email string) Filter {
	return func(fc *FilterConfig) error {
		fc.Email = email
		return nil
	}
}

func FirstNameFilter(firstName string) Filter {
	return func(fc *FilterConfig) error {
		fc.FirstName = firstName
		return nil
	}
}

func MiddleNameFilter(middleName string) Filter {
	return func(fc *FilterConfig) error {
		fc.MiddleName = middleName
		return nil
	}
}

func LastNameFilter(lastName string) Filter {
	return func(fc *FilterConfig) error {
		fc.LastName = lastName
		return nil
	}
}

//Storer represents a user store
//this can be used to store user on different
//data store
type Storer interface {
	List(...Filter) ([]User, error)
	Get(...Filter) (User, error)
	Create(User) error
	Update(User) error
	Delete(...Filter) error
}
