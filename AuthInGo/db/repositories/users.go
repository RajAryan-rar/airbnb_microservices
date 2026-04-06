package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() (*models.User, error)
	Create(username string, email string, hashedPassword string) error
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	// // prepare the query
	// query := "SELECT * FROM users";

	// //execute the query
	// rows, err := u.db.Query(query);

	// users := &[]models.User{}



	return nil,nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error {
	//prepare the query
	query := "DELETE FROM users WHERE id = ?";

	//execute the query
	result,err := u.db.Exec(query,1)

	if(err != nil) {
		fmt.Println("Error deleting user", err);
		return err;
	}

	rowAffected, rowErr := result.RowsAffected();

	if(rowErr != nil) {
		fmt.Println("Error getting rows affected!", rowErr);
		return rowErr;
	}

	if(rowAffected == 0) {
		fmt.Println("No rows were affected, user not deleted");
		return nil;
	}

	fmt.Println("user deleted successfully, rows affected : ", rowAffected);
	return nil
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db : _db,
	}
} 

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {
	query := "INSERT INTO users (username, email, password) VALUES (?,?,?)"

	result, err := u.db.Exec(query,username, email, hashedPassword)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err;
	}

	rowAffected, rowErr := result.RowsAffected();

	if rowErr != nil {
		fmt.Println("Error getting rows affected:", err);
		return rowErr
	}
	if rowAffected == 0 {
		fmt.Println("No rows were affected, user not created");
		return nil
	}

	fmt.Println("User created successfully, rows affected:", rowAffected)
	return nil

}

func (u *UserRepositoryImpl) GetById() (*models.User,error) {
	fmt.Println("fetching user in UserRepository")
	
	//Step 1: Prepare the query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"
	
	//Step 2: Execute the query
	row := u.db.QueryRow(query,1);

	//Step 3: Process the result
	user := &models.User{}
	
	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt) //this line copies data from row(result) to the user object
	
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return  nil, err;
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err;
		}
	}

	//Step 4: Print the user details 
	fmt.Println("User fetched successfully:", user);

	return user, nil
}