
BEGIN;

CREATE TABLE maritials (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE educations (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE genders (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE card_categories (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE income_categories (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE attrition_flags (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE clients (
    id INT GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    age INT NOT NULL,
    maritial_id INT REFERENCES maritials(id) NOT NULL,
    education_id INT REFERENCES educations(id) NOT NULL,
    gender_id INT REFERENCES genders(id) NOT NULL,
    card_category_id INT REFERENCES card_categories(id) NOT NULL,
    attrition_flag_id INT REFERENCES attrition_flags(id) NOT NULL,
    months_on_book INT NOT NULL,
    credit_limit REAL NOT NULL,
    total_relationship_count INT NOT NULL,
    month_inactivate_12_mon INT NOT NULL,
    contacts_count_12_mon INT NOT NULL,
    total_removing_bal INT NOT NULL,
    avg_open_to_buy REAL NOT NULL,
    total_amt_chng_q4_q1 REAL NOT NULL,
    total_trans_amt REAL NOT NULL,
    total_trans_ct INT NOT NULL,
    total_ct_chng_q4_q1 REAL NOT NULL, 
    avg_utilization_ratio REAL NOT NULL,
    naive_bayes_classifier_1 REAL NOT NULL,
    naive_bayes_classifier_2 REAL NOT NULL    
);

COMMIT;
