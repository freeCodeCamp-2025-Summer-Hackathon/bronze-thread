# API Endpoints for SwapCart Backend

This document outlines the core API endpoints for the SwapCart application.

---

### 1. User Authentication

#### A. Register User
* **Endpoint:** `POST /auth/register`
* **Description:** Creates a new user account.
* **Request Body:**
    ```json
    {
      "username": "john.doe",
      "email": "john.doe@example.com",
      "password": "secure_password_123"
    }
    ```
* **Response (Success - HTTP 201 Created):**
    ```json
    {
      "message": "User registered successfully",
      "user_id": "uuid-of-new-user"
    }
    ```

#### B. Login User
* **Endpoint:** `POST /auth/login`
* **Description:** Authenticates a user and provides an authentication token.
* **Request Body:**
    ```json
    {
      "email": "john.doe@example.com",
      "password": "secure_password_123"
    }
    ```
* **Response (Success - HTTP 200 OK):**
    ```json
    {
      "message": "Login successful",
      "token": "jwt_authentication_token_here",
      "user_id": "uuid-of-user"
    }
    ```

---

### 2. Item Management

#### A. List an Item
* **Endpoint:** `POST /items`
* **Description:** Allows a logged-in user to list an item for swapping.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**
    ```json
    {
      "name": "The Lord of the Rings Book Set",
      "description": "Hardcover trilogy, good condition.",
      "category": "books",
      "condition": "used - good - new",
      "photos": ["url_to_photo1.jpg", "url_to_photo2.jpg"],
      "location": {
        "latitude": 22.3072,
        "longitude": 73.1812
      },
    }
    ```
* **Response (Success - HTTP 201 Created):**
    ```json
    {
      "message": "Item listed successfully",
      "item_id": "uuid-of-new-item"
    }
    ```

#### B. Get All Items (Browse Nearby Listings)
* **Endpoint:** `GET /items`
* **Description:** Retrieves a list of available items, optionally filtered by location or category.
* **Authorization:** Optional (can be browsed by guests or logged-in users).
* **Request (Query Parameters):**
    * `?latitude=22.3072&longitude=73.1812&radius=10km` (for nearby listings)
    * `?category=gadgets`
    * `?search=smartphone`
* **Response (Success - HTTP 200 OK):**
    ```json
    {
      "items": [
        {
          "item_id": "uuid-item-1",
          "name": "Wireless Headphones",
          "description": "Noise-cancelling, used once.",
          "category": "gadgets",
          "condition": "like new",
          "owner_id": "uuid-user-A",
          "location": {
             "latitude": 22.3080,
             "longitude": 73.1820
          }
        },
        {
          "item_id": "uuid-item-2",
          "name": "Vintage Denim Jacket",
          "description": "Size M, good condition.",
          "category": "clothes",
          "condition": "used - good",
          "owner_id": "uuid-user-B",
          "location": {
             "latitude": 22.3050,
             "longitude": 73.1790
          }
        }
      ],
      "total_results": 2
    }
    ```

#### C. Get a Single Item
* **Endpoint:** `GET /items/{item_id}`
* **Description:** Retrieves details for a specific item listing.
* **Authorization:** Optional.
* **Response (Success - HTTP 200 OK):**
    ```json
    {
      "item_id": "uuid-item-1",
      "name": "Wireless Headphones",
      "description": "Noise-cancelling, used once.",
      "category": "gadgets",
      "condition": "like new",
      "photos": ["url_to_photo1.jpg", "url_to_photo2.jpg"],
      "owner_id": "uuid-user-A",
      "owner_username": "userA_username",
      "location": {
         "latitude": 22.3080,
         "longitude": 73.1820
      },
      "status": "available",
      "created_at": "2025-07-01T10:00:00Z"
    }
    ```

---

### 3. Swap Management

#### A. Initiate a Swap
* **Endpoint:** `POST /swaps`
* **Description:** Allows a user to propose a swap for another user's item.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**
    ```json
    {
      "offered_item_id": "uuid-user-offering-item",
      "requested_item_id": "uuid-user-requested-item"
    }
    ```
* **Response (Success - HTTP 201 Created):**
    ```json
    {
      "message": "Swap request initiated successfully",
      "swap_id": "uuid-of-new-swap",
      "status": "pending"
    }
    ```

#### B. Accept/Decline Swap Request
* **Endpoint:** `PUT /swaps/{swap_id}`
* **Description:** Allows the recipient of a swap request to accept or decline it.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**
    ```json
    {
      "action": "accept" // or "decline"
    }
    ```
* **Response (Success - HTTP 200 OK):**
    ```json
    {
      "message": "Swap request updated",
      "swap_id": "uuid-of-swap",
      "status": "accepted" // or "declined"
    }
    ```

---

### 4. Rating Management

#### A. Submit a Rating
* **Endpoint:** `POST /ratings`
* **Description:** Allows a user to rate their experience after a completed trade.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**
    ```json
    {
      "swap_id": "uuid-of-completed-swap",
      "rated_user_id": "uuid-of-user-being-rated",
      "rating": 5,
      "comment": "Smooth trade, item exactly as described!"
    }
    ```
* **Response (Success - HTTP 201 Created):**
    ```json
    {
      "message": "Rating submitted successfully",
      "rating_id": "uuid-of-new-rating"
    }
    ```