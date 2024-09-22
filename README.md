1. User Management Endpoints
   
1.1. Register a New User
Endpoint: POST /api/users/register
Description: Registers a new user.
Request Body:
userid (Integer)
username (String)
email (String)
role (String)
password (String)
firstName (String)
lastName (String)

1.2. User Login
Endpoint: POST /api/users/login
Description: Allows a user to log in using their email and password.
Request Body:
email (String)
password (String)

1.3. Get User Profile
Endpoint: GET /api/users/profile
Description: Retrieves user profile details using a Bearer token for authentication.

1.4. Update User Profile
Endpoint: PUT /api/users/profile
Description: Updates user profile details (e.g., username, email, password).
Request Body:
username (String)
email (String)
password (String)

2. Menu Management Endpoints
   
2.1. Create a New Menu Item
Endpoint: POST /api/menu
Description: Creates a new menu item.
Request Body:
menuid (String)
name (String)
description (String)
price (Integer)
category (String)
Optionally upload an image file.

2.2. Update an Existing Menu Item
Endpoint: PUT /api/menu/updateMenu
Description: Updates an existing menu item by menuid.
Request Body:
menuid (String)
name (String)
description (String)
price (Integer)
Optionally upload an image file.

2.3. Delete a Menu Item
Endpoint: DELETE /api/menu/deleteMenu
Description: Deletes a menu item by its ID.
Query Parameter: id (String)

2.4. Fetch All Menu Items
Endpoint: GET /api/menu/fetchAllMenu
Description: Retrieves all menu items.

2.5. Fetch a Menu Item by ID
Endpoint: GET /api/menu/fetchMenu
Description: Retrieves a specific menu item by menuid.
Query Parameter: id (String)

3. Reservation Management Endpoints

3.1. Create a New Reservation
Endpoint: POST /api/reservations
Description: Creates a reservation for a user.
Request Body:
userId (Integer)
dateTime (String, ISO format)
numberOfPeople (Integer)
specialRequests (String, optional)

3.2. Fetch All Reservations
Endpoint: GET /api/getAllReservations
Description: Retrieves all reservations.

3.3. Fetch Reservations by User ID
Endpoint: GET /api/reservationsByID
Description: Retrieves reservations for a specific user.
Query Parameter: userID (Integer)

3.4. Update Reservation by User ID
Endpoint: PUT /api/reservationsByID
Description: Updates a reservation for a specific user.
Query Parameter: userID (Integer)
Request Body:
dateTime (String, ISO format)
numberOfPeople (Integer)
specialRequests (String, optional)

4. Order Management Endpoints

4.1. Create a New Order
Endpoint: POST /api/orders
Description: Creates a new order for a user.
Request Body:
userId (Integer)
items: Array of objects containing:
menuItemId (String)
quantity (Integer)
totalPrice (Integer)
status (String, e.g., Pending)

4.2. Update an Order
Endpoint: PUT /api/orders/{orderid}
Description: Updates an existing order with modified items and status.
Request Body:
items: Array of objects containing:
menuItemId (String)
quantity (Integer)
status (String, e.g., Completed)

4.3. Fetch All Orders
Endpoint: GET /api/ordersAll
Description: Retrieves all orders.

4.4. Fetch Orders by User ID
Endpoint: GET /api/orderByUser
Description: Retrieves orders for a specific user by their userID.
Query Parameter: userID (Integer)
