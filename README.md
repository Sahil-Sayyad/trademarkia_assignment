# Backend Development for Order and Inventory Management Platform using Golang
  
 <b> Develop a backend system in Golang for an order and inventory management
platform. This platform will feature dynamic pricing for products influenced by factors such as
demand and availability. <b> 

## Table of Contents
-  <b> [Getting Started](#getting-started)</b>
-  <b> [Tech Stack](#Tech-Stack) </b>
-  <b> [Project Demo](#Project-Demo) </b>
-  <b> [Author](#Author)</b>

## Getting Started
-  <b> 1. &nbsp; Clone Git Repo  </b>
    <br>----<i> https://github.com/Sahil-Sayyad/trademarkia_assignment.git</i><br><br>
-  <b> 2.  &nbsp;Set Up .env file  </b>
   <br>----<i> add postgres database configuration</i>
   <br>----<i> add jwt sceret</i> 
-  <b> 3. &nbsp; Then simply start your app </b>
   <br>----<i>go run main.go </i><br><br>


### Prerequisites
- <b> Go lang Must be Installed in your System.</b>

## Tech Stack

- <b> 1. Back-end </b>
   <p>Go lang: For server-side development.<br/>
      Go Fiber: As a framework to create the application's server-side routes and manage the HTTP requests and responses.
      PostgreSQL: As the RDBMS database to store and manage the  users, products, orders.</p>
-  <b> 2. Authentication </b>
    <p>JWT: For implementing the authentication system and managing user sign-up and sign-in.</p>

## Usage API Documentation POSTMAN 
<a href = "https://www.postman.com/research-specialist-63110380/workspace/trademarkia/collection/24358323-cfc4367c-4962-4059-8158-822d4b5ef3e7?action=share&creator=24358323"> <b>Link</b> </a>
```
API Structure : 

User-Side APIs:

- POST /api/users/signup          (Create a new user)
- POST /api/users/login           (User authentication)
- GET  /api/products              (Search for products)
- GET  /api/products/:id          (Get a specific product)
- POST /api/orders                (Place an order)
- GET  /api/users/dashboard       (View user's order history)

Admin-Side APIs:

- POST   /api/admin/sign-up        (Add a new admin)
- POST   /api/admin/login          (Admin authentication)
- POST   /api/admin/products       (Add a new product)
- PUT    /api/admin/products/:id   (Update a product)
- DELETE /api/admin/products/:id   (Remove a product)
- GET    /api/admin/orders         (Get all orders with filters/sorting)

        filters and Sorting :

                -Get all orders: /api/admin/orders
                -Filter by user ID: /api/admin/orders?user_id=1
                -Filter by product ID: /api/admin/orders?product_id=1
                -Sort by total price (ascending): /api/admin/orders?sort_by=total_price&order_by=asc

- GET    /api/admin/stats          (Get statistics on orders, inventory)

        Stats on Orders and Inventory :

                -Total Orders: Number of orders placed.
                -Total Revenue: Total revenue generated from orders.
                -Low Stock Products: Products with low inventory (quantity below a threshold).
                -Total Users: Number of registered users.
                -Recent Orders: A list of the most recent orders (you can adjust the Limit as needed).
                -Average Order Value: Calculated by dividing total revenue by the number of orders.

```
## Database Trigger:

Dynamic Pricing Trigger 

```
CREATE OR REPLACE FUNCTION update_dynamic_pricing()
RETURNS TRIGGER AS $$
DECLARE
    current_quantity INT;
    demand_factor    FLOAT; -- Calculate this based on recent sales.
    availability_factor FLOAT;
BEGIN

    -- 1. Fetch the current quantity
    SELECT quantity INTO current_quantity FROM products WHERE id = NEW.id; 

    -- 2. Calculate availability factor 
    availability_factor := 1.0 / (current_quantity + 1);

    -- 3. Calculate the new dynamic price
    NEW.price := NEW.price * (1 + demand_factor) * availability_factor; 

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER dynamic_pricing_trigger
BEFORE UPDATE OF quantity ON products
FOR EACH ROW
EXECUTE FUNCTION update_dynamic_pricing();
```
## Author


- Name: Sahil Sayyad
- GitHub:  <a href = "https://github.com/Sahil-Sayyad/trademarkia_assignment"> <b>Link</b> </a>
- Email: sahilsayyad.dev@gmail.com
