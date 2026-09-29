# ADR 0003: Selective home-authority federation

- **Status:** Accepted
- **Date:** 2026-09-28

## Context

Self-hosting gives deployment control but does not let independent organizations coordinate without one becoming the other's tenant. Whole-graph replication would expose private topology and multiply conflict, retention, and moderation problems. Global blockchain consensus is unrelated to most work authority.

## Decision

Each federated object has one home authority. Nodes exchange only addressed boundary contracts, proposals, accepted facts within sender authority, receipts, and selected evidence. Remote nodes apply their own policy and cannot directly overwrite a home-owned object.

Delivery is signed, at least once, idempotent, versioned, and tolerant of partitions. Disputes remain explicit. Home migration requires a signed transfer chain.

## Consequences

- teams keep private internal decomposition and local vocabulary;
- cross-organization commitments can be shared without a global database;
- home-node outage and migration need explicit recovery;
- cryptographic identity, moderation, replay defense, and disclosure inference are core federation work;
- federation follows a correct local kernel and two-node experiment.

## Reconsider if

- pilots require peer-to-peer co-ownership with no meaningful home authority;
- a mature external protocol fully covers Trail's commitment semantics and safety;
- the target users obtain no value from cross-organization coordination.

