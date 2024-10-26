CREATE TABLE IF NOT EXISTS incomes (
    id SERIAL PRIMARY KEY,
    amount NUMERIC(11,2) NOT NULL DEFAULT 0,
    description VARCHAR(25) NOT NULL DEFAULT '',
    budget_id INT NOT NULL,
    entry_class_id INT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_budget
      FOREIGN KEY (budget_id)
      REFERENCES budgets(id)
      ON DELETE CASCADE,
    CONSTRAINT fk_entry_class
      FOREIGN KEY (entry_class_id)
      REFERENCES entry_classes(id)
      ON DELETE CASCADE
);
