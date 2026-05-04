// Model for Products in the database
// this serves as a template for the product data structure and how it will be represented in JSON format when sent to the client.

package model

type Product struct {
	ID    int     `json:"id_product"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// `json:"xxx"` is a struct tag that specifies how the field should be encoded/decoded when converting to/from JSON.

/* 
For example, when a Product struct is converted to JSON, the ID field will be represented as "id_product", 
the Name field as "name", and the Price field as "price".
This allows for more control over the JSON representation of the data when it is sent to the client.
 */