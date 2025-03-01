-- Create the users table
CREATE TABLE users (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY, -- Unique user ID
    wallet_address VARCHAR(255) NOT NULL UNIQUE,        -- User's wallet address
    has_access BOOLEAN DEFAULT FALSE                    -- Whether the user has access
);

-- Create the matches table
CREATE TABLE matches (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY, -- Unique match ID
    created_at TIMESTAMP DEFAULT NOW(),                 -- Timestamp for match creation
    game_hash_id VARCHAR(255) UNIQUE NOT NULL,          -- Unique GameHashID for each session
    expiration_date TIMESTAMP NOT NULL,                 -- Expiration date for the session
    wallet_address VARCHAR(255) NOT NULL,               -- User's wallet address
    transaction_hash VARCHAR(88) NOT NULL,              -- Transaction hash from payment verification
    CONSTRAINT fk_matches_wallet_address FOREIGN KEY (wallet_address)
        REFERENCES users(wallet_address) ON DELETE CASCADE
);

-- Create the match_results table
CREATE TABLE match_results (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY, -- Unique result ID
    match_id INT NOT NULL,                               -- Associated match ID
    user_id INT NOT NULL,                                -- Associated user ID
    kills INT DEFAULT 0,                                 -- Number of kills the user achieved
    is_winner BOOLEAN DEFAULT FALSE,                     -- Whether the user won the match
    reward_amount NUMERIC(10, 2),                        -- Reward amount for the user
    FOREIGN KEY (match_id) REFERENCES matches (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Create payment_verifications table
CREATE TABLE payment_verifications (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    user_id INT NOT NULL,
    wallet_address VARCHAR(255) NOT NULL,
    amount NUMERIC(20, 8) NOT NULL,                     -- Stores SOL amount with 8 decimal places
    transaction_hash VARCHAR(88),                       -- Solana tx hash length
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'verified', 'failed')),
    access_granted BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);