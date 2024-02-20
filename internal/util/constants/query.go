package constants

// product handler queries
const (
	FETCH_ALL_PRODUCTS  = "SELECT p.id, p.title, p.short_description, p.description, p.price, p.quantity, p.image_url, c.id, c.title FROM products p, categories c where p.category = c.id"
	FETCH_PRODUCT_BY_ID = "SELECT p.id, p.title, p.short_description, p.description, p.price, p.quantity, p.image_url, c.id, c.title FROM products p, categories c where p.category = c.id and p.id = $1"
)

// auth handler queries
const (
	CREATE_USER     = "INSERT INTO users (id, first_name, last_name, email, password) values ($1, $2, $3, $4, $5)"
	FIND_USER_EMAIL = "SELECT id, password FROM users where email = $1"
)

// cart handler queries
const (
	FETCH_CART       = "SELECT c.id, c.quantity, c.total_price, p.id, p.title, p.image_url FROM cart_items c, products p WHERE c.product_id = p.id AND user_id = $1"
	FIND_CART_ITEM   = "SELECT quantity, total_price FROM cart_items where product_id = $1"
	ADD_TO_CART      = "INSERT INTO cart_items (id, user_id, product_id, quantity, total_price) VALUES ($1, $2, $3, $4, $5)"
	UPDATE_CART      = "UPDATE cart_items SET quantity = $1, total_price = $2 where product_id = $3 RETURNING id"
	DELETE_CART_ITEM = "DELETE FROM cart_items where id = $1"
	CLEAR_CART       = "DELETE FROM cart_items where user_id = $1"
)
