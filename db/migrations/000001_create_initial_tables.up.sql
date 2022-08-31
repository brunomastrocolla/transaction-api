CREATE TABLE IF NOT EXISTS operation_type(
    id INT PRIMARY KEY,
    description VARCHAR NOT NULL
);

INSERT INTO operation_type (id, description) VALUES
    ( 1, 'COMPRA A VISTA'),
    ( 2, 'COMPRA PARCELADA'),
    ( 3, 'SAQUE'),
    ( 4, 'PAGAMENTO') ON CONFLICT DO NOTHING;

CREATE TABLE IF NOT EXISTS accounts(
    id BIGSERIAL PRIMARY KEY,
    document_number VARCHAR NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS transactions(
    id BIGSERIAL PRIMARY KEY,
    account_id BIGINT NOT NULL,
    operation_type_id INT NOT NULL,
    amount FLOAT8 NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    FOREIGN KEY (account_id) REFERENCES accounts (id),
    FOREIGN KEY (operation_type_id) REFERENCES operation_type (id)
);
