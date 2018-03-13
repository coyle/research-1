package boltdb

// User defines the schema
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	UUID     string `json:"uuid"`
}

// CreateUser calls the underlying bolt database instance to create a user
func (bdb *Client) CreateUser(key, value []byte) {

	bdb.UsersBucket.Put(key, value)

}
