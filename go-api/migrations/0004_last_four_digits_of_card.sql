ALTER TABLE cards
ADD COLUMN last_four_digits VARCHAR(4);
ALTER TABLE cards
ADD CONSTRAINT unique_card_token UNIQUE (card_token);