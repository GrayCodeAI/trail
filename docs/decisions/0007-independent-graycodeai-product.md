# ADR 0007: Trail is an independent GraycodeAI product

- **Status:** Accepted
- **Date:** 2026-09-28

## Context

Trail belongs to the GraycodeAI product family and should feel coherent with the wider ecosystem. It must also remain self-hostable, independently operable, and usable without GraycodeAI cloud identity. Treating Trail as an internal feature of Graycode Platform would couple coordination authority, storage, releases, and availability to a separate product boundary.

The user clarified on 2026-09-28 that Trail is an independent product under `graycodeai.com`.

## Decision

Trail is an independent product by GraycodeAI. Its hosted surface is intended for `trail.graycodeai.com`.

Trail owns its:

- source repository and release lifecycle;
- API and protocol compatibility policy;
- deployables, persistence, queues, keys, backups, and operations;
- local and federated identity configuration;
- product UI, CLI, agent edge, and self-hosted distribution;
- workspace authority, policies, events, projections, and audit state.

Graycode Platform may provide optional cross-product account, organization, entitlement, and billing services through versioned authenticated HTTP APIs. Trail must continue to operate in self-hosted profiles without those services. The products do not share a database, private package, queue, deployment unit, signing authority, or implicit administrator role.

Visual design may reuse published GraycodeAI design tokens or public packages with explicit versions. Trail must not import private application code from Graycode Platform.

## Consequences

- Trail failures and upgrades can be isolated from the company website and other hosted products;
- self-hosted users can configure local or external identity without a GraycodeAI account;
- cross-product navigation and billing require explicit contracts and failure handling;
- hosted single sign-on must exchange bounded identity assertions rather than database identifiers;
- product analytics, retention, support access, and incident response need Trail-specific policies;
- Trail can federate with non-GraycodeAI nodes as a peer protocol implementation.

## Evidence gate

Before hosted launch, test Platform unavailability, entitlement staleness, account unlinking, export, and self-hosted installation. Trail's core coordination path must remain available when optional Platform services are unreachable, subject to locally cached entitlement policy for hosted commercial features.

## Reconsider if

- a future legal or operational structure requires a different public domain;
- self-hosted distribution is deliberately discontinued through a separate product decision;
- a versioned public Platform capability replaces a Trail-owned concern without sharing authority state.
