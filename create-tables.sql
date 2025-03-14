DROP TABLE IF EXISTS user_all_cards;

CREATE TABLE user_all_cards(
  id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
  card_name VARCHAR(128) NOT NULL,
  mana_cost VARCHAR(255) NOT NULL,
  mana_cost_colors CHAR[],
  types VARCHAR[] 
);
