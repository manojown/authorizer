package couchbase

import (
	"context"
	"time"

	"github.com/authorizerdev/authorizer/server/db/models"
	"github.com/google/uuid"
)

// AddSession to save session information in database
func (p *provider) AddSession(ctx context.Context, session models.Session) error {
	if session.ID == "" {
		session.ID = uuid.New().String()
	}

	session.CreatedAt = time.Now().Unix()
	session.UpdatedAt = time.Now().Unix()
	return nil
}

// DeleteSession to delete session information from database
func (p *provider) DeleteSession(ctx context.Context, userId string) error {
	return nil
}
