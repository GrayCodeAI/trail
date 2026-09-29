package domain

import (
	"fmt"
	"testing"
	"time"
)

func TestProtectedEffectAuthorityNeverReturnsAcrossInterleavings(t *testing.T) {
	operations := []string{"revoke", "release", "deadline", "complete"}
	sequences := cartesianSequences(operations, 4)
	for _, sequence := range sequences {
		name := fmt.Sprintf("%s-%s-%s-%s", sequence[0], sequence[1], sequence[2], sequence[3])
		t.Run(name, func(t *testing.T) {
			rig := newTestRig(t)
			rig.policy.Allow(workspace, workload, ActionLeaseRelease)
			outcome := rig.createOutcome("outcome:explore")
			commitment := rig.createActiveCommitment("commitment:explore", outcome.ID)
			work := rig.createWork("work:explore", outcome.ID, commitment.ID)
			grant := rig.issueGrant("grant:explore", work.ID, rig.clock.Now().Add(4*time.Hour))
			lease := rig.acquireLease("lease:explore", work.ID, rig.clock.Now().Add(4*time.Hour))
			requireAccepted(t, rig.startRun("run:explore", work, lease, grant))

			model := struct {
				runActive     bool
				grantActive   bool
				leaseActive   bool
				deadlineValid bool
			}{true, true, true, true}
			assertEffectMatchesModel(t, rig, lease, grant, model.runActive && model.grantActive && model.leaseActive && model.deadlineValid)
			authorityLost := false

			for _, operation := range sequence {
				switch operation {
				case "revoke":
					if model.grantActive {
						current, _ := rig.kernel.Grant(grant.ID)
						requireAccepted(t, rig.kernel.RevokeGrant(RevokeGrantCommand{
							Meta:   rig.meta(CommandRevokeGrant, grant.ID, current.Version, direct(steward)),
							Reason: "model exploration",
						}))
						model.grantActive = false
					}
				case "release":
					if model.leaseActive {
						current, _ := rig.kernel.Lease(lease.ID)
						requireAccepted(t, rig.kernel.ReleaseLease(ReleaseLeaseCommand{
							Meta:   rig.meta(CommandReleaseLease, lease.ID, current.Version, direct(workload)),
							Reason: "model exploration",
						}))
						model.leaseActive = false
					}
				case "deadline":
					if model.deadlineValid {
						rig.clock.Advance(2 * time.Hour)
						model.deadlineValid = false
					}
				case "complete":
					if model.runActive {
						current, _ := rig.kernel.Run("run:explore")
						requireAccepted(t, rig.kernel.CompleteRun(CompleteRunCommand{
							Meta:    rig.meta(CommandCompleteRun, current.ID, current.Version, delegated(operator, agent)),
							Summary: "model exploration completion",
						}))
						model.runActive = false
					}
				}

				expected := model.runActive && model.grantActive && model.leaseActive && model.deadlineValid
				if authorityLost && expected {
					t.Fatal("model restored protected-effect authority without a new authorization path")
				}
				assertEffectMatchesModel(t, rig, lease, grant, expected)
				if !expected {
					authorityLost = true
				}
			}
		})
	}
}

func assertEffectMatchesModel(t *testing.T, rig *testRig, lease ClaimLease, grant CapabilityGrant, expected bool) {
	t.Helper()
	check := rig.kernel.CheckSideEffect("run:explore", lease.FencingToken, grant.ID)
	if check.Allowed != expected {
		t.Fatalf("protected-effect check allowed=%v code=%s, want allowed=%v", check.Allowed, check.Code, expected)
	}
}

func cartesianSequences(values []string, length int) [][]string {
	if length == 0 {
		return [][]string{{}}
	}
	shorter := cartesianSequences(values, length-1)
	sequences := make([][]string, 0, len(values)*len(shorter))
	for _, value := range values {
		for _, suffix := range shorter {
			sequence := make([]string, 1, length)
			sequence[0] = value
			sequence = append(sequence, suffix...)
			sequences = append(sequences, sequence)
		}
	}
	return sequences
}
