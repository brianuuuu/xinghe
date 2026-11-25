// Package store owns the storage boundaries:
// - Git formal document repository (serialized commits)
// - SQLite authoritative runtime tables (backup required)
// - SQLite rebuildable projections (read-only mirror of Git)
// - Content-addressed artifact storage and Git/SQLite command reconciliation
package store
