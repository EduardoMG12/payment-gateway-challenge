-- Add migration script here
ALTER TABLE transactions
ADD COLUMN refund_transaction_id UUID REFERENCES transactions(id);