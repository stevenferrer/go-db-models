package model

//User represents an application user
type User struct {
	ID         int64
	Email      string
	Password   string
	FirstName  string
	MiddleName string
	LastName   string
}

//UserFilter is used for filtering results
type UserFilter func(*UserFilterConfig) error

//UserFilterConfig is the config for filtering results
type UserFilterConfig struct {
	Limit      int64
	ID         int64
	Email      string
	FirstName  string
	MiddleName string
	LastName   string
}

func UserLimitFilter(limit int64) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Limit = limit
		return nil
	}
}

func UserIDFilter(id int64) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.ID = id
		return nil
	}
}

func UserEmailFilter(email string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Email = email
		return nil
	}
}

func UserFirstNameFilter(firstName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.FirstName = firstName
		return nil
	}
}

func UserMiddleNameFilter(middleName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.MiddleName = middleName
		return nil
	}
}

func UserLastNameFilter(lastName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.LastName = lastName
		return nil
	}
}

//UserStorer represents a user store
//this can be used to store user on different
//data store
type UserStorer interface {
	List(...UserFilter) ([]User, error)
	Get(...UserFilter) (User, error)
	Create(User) error
	Update(User) error
	Delete(...UserFilter) error
}
