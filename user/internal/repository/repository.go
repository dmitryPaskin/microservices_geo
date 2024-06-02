package repository

import (
	proto "GEO_API/user/pkg/gRPC/api"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
)

type AuthRepository interface {
	SaveUser(user *proto.User) error
	CheckUser(user *proto.User) (bool, error)
	GetUser(user *proto.User) (*proto.User, error)
	GetListUsers() (*[]*proto.User, error)
}

type authRepository struct {
	db         *sql.DB
	sqlBuilder sq.StatementBuilderType
}

func New(database *sql.DB) AuthRepository {
	return &authRepository{
		db:         database,
		sqlBuilder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (ar *authRepository) SaveUser(user *proto.User) error {
	query := ar.sqlBuilder.Insert("users").
		Columns("login", "password").
		Values(user.Login, user.Password)

	if _, err := query.RunWith(ar.db).Exec(); err != nil {
		return err
	}

	return nil
}

func (ar *authRepository) CheckUser(user *proto.User) (bool, error) {
	query := ar.sqlBuilder.Select("COUNT(*)").
		From("users").
		Where(sq.Eq{"login": user.Login})

	row := query.RunWith(ar.db).QueryRow()
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (ar *authRepository) GetUser(user *proto.User) (*proto.User, error) {
	query := ar.sqlBuilder.Select("login", "password").
		From("users").Where(sq.Eq{"login": user.Login})

	row := query.RunWith(ar.db).QueryRow()
	newUser := proto.User{}
	if err := row.Scan(&newUser.Login, &newUser.Password); err != nil {
		return &proto.User{}, err
	}
	return &newUser, nil
}

func (ar *authRepository) GetListUsers() (*[]*proto.User, error) {
	query, args, err := ar.sqlBuilder.Select("Login", "Password").
		From("users").
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := ar.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*proto.User

	for rows.Next() {
		var user proto.User
		if err := rows.Scan(&user.Login, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &users, nil
}
