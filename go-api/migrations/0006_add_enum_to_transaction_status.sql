CREATE TYPE transaction_status AS ENUM (
    'PENDING',
    'APPROVED',
    'REJECTED',
    'ERROR'
);