package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration013SubmissionAddProposerIndex = &migrate.Migration{
	Id: "013-submission-add-proposer-index",
	Up: []string{`
		ALTER TABLE ` + vars.TableBuilderBlockSubmission + ` ADD COLUMN proposer_index BIGINT NOT NULL DEFAULT 0;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
