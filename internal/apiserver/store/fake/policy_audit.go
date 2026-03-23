package fake

import "context"

type policyAudit struct {
	ds *datastore
}

func newPolicyAudits(ds *datastore) *policyAudit {
	return &policyAudit{ds}
}

// ClearOutdated clear data older than a given days.
func (p *policyAudit) ClearOutdated(ctx context.Context, maxReserveDays int) (int64, error) {
	return 0, nil
}
