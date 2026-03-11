package migrations

import (
	"github.com/flashbots/mev-boost-relay/database/vars"
	migrate "github.com/rubenv/sql-migrate"
)

var Migration012PayloadAddProposerIndex = &migrate.Migration{
	Id: "012-payload-add-proposer-index",
	Up: []string{`
		ALTER TABLE ` + vars.TableDeliveredPayload + ` ADD COLUMN proposer_index BIGINT NOT NULL DEFAULT 0;
	`},
	Down: []string{},

	DisableTransactionUp:   true,
	DisableTransactionDown: true,
}
