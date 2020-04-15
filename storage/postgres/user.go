package postgres

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq" // for db driver

	"bitbucket.org/alien_soft/api_gateway/pkg/etc"
	"bitbucket.org/alien_soft/api_gateway/storage/repo"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db: db}
}

func (cm *userRepo) Create(cl *repo.User) error {
	var (
		mail sql.NullString = etc.NullString(cl.Mail)
		err  error
	)

	if cl.Mail.GetValue() == "" {
		insertNew :=
			`INSERT INTO
		auth_users
		(
			id,
			created_at,
			updated_at,
			mail,
			password,
			access_token,
			refresh_token,
			user_type_id
		)
			VALUES  
			($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL, $2, $3, $4, $5)`
		_, err = cm.db.Exec(
			insertNew,
			cl.ID,
			cl.Password,
			cl.AccessToken,
			cl.RefreshToken,
			cl.UserTypeID,
		)
	} else {
		insertNew :=
			`INSERT INTO
		auth_users
		(
			id,
			created_at,
			updated_at,
			mail,
			password,
			access_token,
			refresh_token,
			user_type_id
		)
			VALUES  
			($1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, $2, $3, $4, $5, $6)`
		_, err = cm.db.Exec(
			insertNew,
			cl.ID,
			mail,
			cl.Password,
			cl.AccessToken,
			cl.RefreshToken,
			cl.UserTypeID,
		)
	}

	if pgerr, ok := err.(*pq.Error); ok {
		if pgerr.Code == "23505" {
			return repo.ErrAlreadyExists
		}
		return err
	} else if err != nil {
		return err
	}

	return nil
}

func (cm *userRepo) Update(cl *repo.User) error {
	updateClient := `
	UPDATE auth_users 
	SET
		updated_at = current_timestamp,
		mail = $1,
		password = $2,
		access_token = $3,
		refresh_token = $4,
		user_type_id = $5,
		is_verified = $6
	WHERE id = $7
	`
	_, err := cm.db.Exec(updateClient,
		cl.Mail,
		cl.Password,
		cl.AccessToken,
		cl.RefreshToken,
		cl.UserTypeID,
		cl.IsVerified,
		cl.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cm *userRepo) GetUser(id string) (*repo.User, error) {
	cli := repo.User{}
	var (
		createdAt, updatedAt time.Time
		mail                 sql.NullString
	)

	row := cm.db.QueryRow("SELECT id, created_at, updated_at, mail, "+
		" password, access_token, refresh_token, user_type_id "+
		"FROM auth_users WHERE deleted_at IS NULL AND id = $1", id)
	err := row.Scan(
		&cli.ID,
		&createdAt,
		&updatedAt,
		&mail,
		&cli.Password,
		&cli.AccessToken,
		&cli.RefreshToken,
		&cli.UserTypeID,
	)
	if err != nil {
		return nil, err
	}

	cli.CreatedAt = createdAt.Format(time.RFC3339)
	cli.UpdatedAt = updatedAt.Format(time.RFC3339)
	cli.Mail = etc.StringValue(mail)

	return &cli, nil
}

func (cm *userRepo) GetAllUsers() ([]*repo.User, error) {
	var (
		clients []*repo.User
		mail    sql.NullString
	)

	rows, err := cm.db.Queryx("SELECT id, created_at, updated_at, mail, " +
		" password, access_token, refresh_token, user_type_id" +
		"FROM auth_users WHERE deleted_at IS NULL")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var cli repo.User
		var createdAt, updatedAt time.Time
		if err := rows.Scan(
			&cli.ID,
			&createdAt,
			&updatedAt,
			&mail,
			&cli.Password,
			&cli.AccessToken,
			&cli.RefreshToken,
			&cli.UserTypeID,
		); err != nil {
			return nil, err
		}

		cli.CreatedAt = createdAt.Format(time.RFC3339)
		cli.UpdatedAt = updatedAt.Format(time.RFC3339)
		cli.Mail = etc.StringValue(mail)

		clients = append(clients, &cli)
	}

	return clients, nil
}

func (cm *userRepo) Delete(id string) error {
	_, err := cm.db.Exec(`
		UPDATE auth_users 
		SET deleted_at = current_timestamp
		WHERE id = $1 
		  AND deleted_at IS NULL`, id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (cm *userRepo) CheckMail(mail string) (bool, error) {
	var exists int
	row := cm.db.QueryRow("SELECT 1 FROM auth_users where mail = $1", mail)
	err := row.Scan(&exists)
	if err == sql.ErrNoRows {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if exists == 1 {
		return true, nil
	}
	return false, nil
}

func (cm *userRepo) GetUserByMail(mail string) (*repo.User, error) {
	user := repo.User{}
	var mailNull sql.NullString

	row := cm.db.QueryRow("SELECT id, mail, password, access_token FROM auth_users WHERE mail=$1", mail)

	err := row.Scan(&user.ID, &mailNull, &user.Password, &user.AccessToken)
	if err != nil {
		return nil, err
	}

	user.Mail = etc.StringValue(mailNull)

	return &user, nil
}

func (cm *userRepo) UpdateUserTokens(id, accessToken, refreshToken string) error {
	updateTokens := `
	UPDATE auth_users
	SET
		access_token = $1,
		refresh_token = $2
	WHERE id = $3
	`
	result, err := cm.db.Exec(updateTokens,
		accessToken,
		refreshToken,
		id,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (cm *userRepo) CreateConnection(userID, socialID, accessToken, connType string) error {
	var exists int

	row := cm.db.QueryRow("SELECT count(1) from auth_connections where "+
		"social_id=$1 and type=$2", socialID, connType)

	err := row.Scan(&exists)
	if err != nil {
		return err
	}

	if exists >= 1 {
		update := `UPDATE auth_connections
		SET
			access_token = $1
		WHERE social_id = $2 and type=$3
		`

		_, err = cm.db.Exec(
			update,
			accessToken,
			socialID,
			connType,
		)
		if err != nil {
			return err
		}
	} else {
		if err != nil {
			return err
		}

		insert := `INSERT INTO
			auth_connections
			(
				user_id, 
				social_id, 
				access_token, 
				type
			)
			VALUES
			($1, $2, $3, $4)
		`

		_, err = cm.db.Exec(
			insert,
			userID,
			socialID,
			accessToken,
			connType,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func (cm *userRepo) IsConnectionExists(socialID, connType string) (bool, error) {
	var exists int

	row := cm.db.QueryRow("SELECT count(1) from auth_connections where "+
		"social_id=$1 and type=$2", socialID, connType)

	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	if exists >= 1 {
		return true, nil
	}

	return false, nil
}

func (cm *userRepo) IsConnectionExistsFromUserID(userID, connType string) (bool, error) {
	var exists int

	row := cm.db.QueryRow("SELECT count(1) from auth_connections where "+
		"user_id=$1 and type=$2", userID, connType)

	err := row.Scan(&exists)
	if err != nil {
		return false, err
	}

	if exists >= 1 {
		return true, nil
	}

	return false, nil
}

func (cm *userRepo) GetUserIDFromExistingConnection(socialID, accessToken, connType string) (string, error) {
	var userID string

	row := cm.db.QueryRow("SELECT user_id from auth_connections where "+
		"social_id=$1 and type=$2", socialID, connType)

	err := row.Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (cm *userRepo) DeleteUserConnections(userID string) error {
	_, err := cm.db.Exec(`DELETE FROM auth_connections WHERE user_id=$1`, userID)
	if err != nil {
		return err
	}

	return nil
}

func (cm *userRepo) SetMail(userID, mail string) error {
	result, err := cm.db.Exec(`
		UPDATE auth_users
		SET mail = $1
		WHERE id = $2`, mail, userID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (cm *userRepo) SetPassword(id, password string) error {
	result, err := cm.db.Exec(`
		UPDATE auth_users
		SET password = $1
		WHERE id = $2
	`, password, id)
	if err != nil {
		return err
	}

	if i, _ := result.RowsAffected(); i == 0 {
		return sql.ErrNoRows
	}

	return nil
}


