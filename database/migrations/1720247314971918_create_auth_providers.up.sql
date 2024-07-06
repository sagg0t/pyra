CREATE TABLE IF NOT EXISTS auth_providers
(
    id          BIGSERIAL       PRIMARY KEY,

    user_id     BIGSERIAL       REFERENCES users(id),

    name        VARCHAR(64)     NOT NULL,
    uid         VARCHAR(64)     NOT NULL,

    created_at  TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMPTZ     NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_auth_providers_on_name_and_uid
ON auth_providers (name, uid);
