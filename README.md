collection represents API endpoints for a restaurant management system. The operations include user registration, login, profile management, menu creation, reservations, and order management. Here’s a breakdown of the key endpoints:

User Registration (POST /api/users/register):

Registers a new user with fields like userid, username, email, role, password, firstName, and lastName.
User Login (POST /api/users/login):

Allows a user to log in using email and password.
Get Profile (GET /api/users/profile):

Retrieves user profile details, using a Bearer token for authentication.
Update Profile (PUT /api/users/profile):

Updates user profile details like username, email, and password.
Create Menu (POST /api/menu):

Creates a new menu item with fields such as menuid, name, description, price, and category. An image file can also be uploaded.
Create Reservation (POST /api/reservations):

Creates a reservation for a user with fields like userId, dateTime, numberOfPeople, and specialRequests.
Create Orders (POST /api/orders):

Creates a new order for a user by providing the userId, an array of items (menuItemId and quantity), and totalPrice.
Update Orders (PUT /api/orders/1):

Updates an existing order with modified items and status.
Update Menu (PUT /api/menu/updateMenu):

Updates an existing menu item by menuid with fields like name, description, price, and an image file.
Delete Menu (DELETE /api/menu/deleteMenu?id=chickenCurry2):

Deletes a menu item by its id.
Fetch All Menus (GET /api/menu/fetchAllMenu):

Retrieves all menu items.
Fetch All Orders (GET /api/ordersAll):

Retrieves all orders.
Fetch All Reservations (GET /api/getAllReservations):

Retrieves all reservations.
Get Menu by ID (GET /api/menu/fetchMenu?id=chickenCurry):

Retrieves a menu item by its id.
Get Reservations by ID (GET /api/reservationsByID?userID=11):

Retrieves reservations for a specific user by userID.
Update Reservations by ID (PUT /api/reservationsByID?userID=11):

Updates a reservation for a specific user by userID.
Get Orders by User ID (GET /api/orderByUser?userID=2):

Retrieves orders for a specific user by userID.
Each endpoint uses either Bearer token authentication or requires certain query parameters to function. The raw body for POST/PUT requests is typically in JSON format.




