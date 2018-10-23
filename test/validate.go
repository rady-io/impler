package test

import "time"

// user
var (
	user *User
)

/*
@Validate
 */
type (
	User struct {
		/*
		@NotZero
		 */
		Name string

		/*
		@Email
		 */
		Email string

		/*
		@Length 11
		 */
		Phone string

		/*
		@After 2018-09-24 00:00:00
		@Before 2018-09-25 00:00:00
		 */
		CreatedAt *time.Time
	}
)

/*
@Before log.Debug(user.Name)
@After log.Debug(user.Name)
 */
func (user *User) SetName(name string) {
	user.Name = name
}