package api

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Context struct {
	DBPool     *pgxpool.Pool
	AWSSession *session.Session
}

func GetAWSSession() (*session.Session, error) {
	return session.NewSession(
		&aws.Config{
			Region: aws.String(os.Getenv("AWS_REGION")),
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("AWS_ACCESS_KEY_ID"),
				os.Getenv("AWS_SECRET_ACCESS_KEY"),
				"",
			),
		},
	)
}

func GetDBConn(db_url string) (*pgxpool.Pool, error) {
	return pgxpool.Connect(context.Background(), db_url)
}
