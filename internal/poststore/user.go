package poststore

import (
	"GraphQL_api/graph/model"
)

// CreateUser creates new user in the store and returns id of new user.
func (ts *InMemoryStorage) CreateUser(firstName string, lastName string) (int, error) {
	ts.Lock()
	defer ts.Unlock()

	user := &model.User{
		ID:        ts.nextPostId,
		FirstName: firstName,
		LastName:  lastName,
	}

	ts.users[ts.nextUserId] = user
	ts.nextUserId++
	return user.ID, nil
}

// GetUser returns a user from the store, by id.
// If no such id exists, an error is returned.
func (ts *InMemoryStorage) GetUser(id int) (*model.User, error) {
	ts.Lock()
	defer ts.Unlock()

	user, err := ts.UserExists(id)
	return user, err
}

// DeleteUser deletes the user with the given id. If no such id exists, an error
// is returned.
func (ts *InMemoryStorage) DeleteUser(id int) (bool, error) {
	ts.Lock()
	defer ts.Unlock()

	_, err := ts.UserExists(id)
	if err != nil {
		return false, err
	}

	delete(ts.users, id)
	return true, nil
}

// GetAllUsers returns all users in the store.
func (ts *InMemoryStorage) GetAllUsers() ([]*model.User, error) {
	ts.Lock()
	defer ts.Unlock()

	allUsers := make([]*model.User, 0, len(ts.users))
	for _, user := range ts.users {
		allUsers = append(allUsers, user)
	}
	return allUsers, nil
}
