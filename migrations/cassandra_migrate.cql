-- Comment
USE novabankapp;
CREATE TABLE IF NOT EXISTS wallets (
                                    id uuid,
                                    user_id text,
                                    account_id text,
                                    email text,
                                    available_balance decimal,
                                    balance decimal,
                                    locked boolean,
                                    created_at timestamp,
                                    PRIMARY KEY (id)
                                                  );
CREATE TABLE IF NOT EXISTS wallet_links (
                                       id uuid,
                                       wallet_id text,
                                       account_id text,
                                       value text,
                                       link_date timestamp,
                                       PRIMARY KEY (id)
    );
CREATE TABLE IF NOT EXISTS wallet_states (
                                       id uuid,
                                       wallet_id text,
                                       is_locked boolean,
                                       is_deleted boolean,
                                       is_blacklisted boolean,
                                       PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS wallet_transactions (
                                             id uuid,
                                             credit_wallet_id text,
                                             debit_wallet_id text,
                                             amount decimal,
                                             description text,
                                             created_at timestamp,
                                             PRIMARY KEY (id)
    );
