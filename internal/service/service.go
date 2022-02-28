package service

import (
	"context"

	"github.com/quizlyfun/quizly-backend/internal/domain"
)

type (
	Account interface {
		// Create creates new account.
		Create(ctx context.Context, acc domain.Account) (domain.Account, error)

		// GetByID account.
		GetByID(ctx context.Context, aid string) (domain.Account, error)

		// GetByEmail account.
		GetByEmail(ctx context.Context, email string) (domain.Account, error)

		// Delete sets account IsArchive state to true.
		Delete(ctx context.Context, aid, sid string) error

		// Verify verifies account using provided token.
		Verify(ctx context.Context, aid, code string) error
	}

	AccountRepo interface {
		// Create account with given credentials, returns id of created account.
		Create(ctx context.Context, acc domain.Account) (domain.Account, error)

		// FindByID account in DB.
		FindByID(ctx context.Context, aid string) (domain.Account, error)

		// FindByEmail account in DB.
		FindByEmail(ctx context.Context, email string) (domain.Account, error)

		// Archive sets entity.Account.IsArchive state to provided value.
		Archive(ctx context.Context, aid string, archive bool) error
	}

	Auth interface {
		// EmailLogin creates new session using provided account email and password.
		EmailLogin(ctx context.Context, email, password string, d domain.Device) (domain.Session, error)

		// Logout logs out session by id.
		Logout(ctx context.Context, sid string) error

		// NewAccessToken creates new token with subject (accountID) and audience claims.
		NewAccessToken(ctx context.Context, aid, password, audience string) (string, error)

		// ParseAccessToken parses token and returns subject from claims.
		ParseAccessToken(ctx context.Context, token, audience string) (string, error)
	}

	Session interface {
		// Create new session for account with id and device of given provider.
		Create(ctx context.Context, aid string, d domain.Device) (domain.Session, error)

		// GetByID session.
		GetByID(ctx context.Context, sid string) (domain.Session, error)

		// GetAll account sessions using provided account id.
		GetAll(ctx context.Context, aid string) ([]domain.Session, error)

		// Terminate session by id excluding current session with id.
		Terminate(ctx context.Context, sid, currSid string) error

		// TerminateAll account sessions excluding current session with id.
		TerminateAll(ctx context.Context, aid, sid string) error
	}

	Email interface {
		// SendAccountVerificationEmail sends email with verification link to given email.
		SendAccountVerificationEmail(ctx context.Context, to string) error
	}

	SessionRepo interface {
		// Create new session in DB.
		Create(ctx context.Context, s domain.Session) error

		// FindByID session.
		FindByID(ctx context.Context, sid string) (domain.Session, error)

		// FindAll accounts sessions by provided account id.
		FindAll(ctx context.Context, aid string) ([]domain.Session, error)

		// Delete session by id.
		Delete(ctx context.Context, sid string) error

		// DeleteAll account sessions by provided account id excluding current session.
		DeleteAll(ctx context.Context, aid, sid string) error
	}
)
