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
      "name": "John Doe",
      "email": "john.doe@example.com",
      "password": "secure_password_123",
    }
    ```

* **Response (Success - HTTP 201 Created):**

    ```json
    {
      "message": "User registered successfully",
      "user_id": 12345
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
      "user": {
        "id": 12345,
        "name": "John Doe",
        "email": "john.doe@example.com",
        "bio" : "empty",
        "avatar_url": "https://example.com/avatar.jpg",
        "rating_avg": 4.5,
        "total_ratings": 23,
        "total_swaps": 15,
      }
    }
    ```

#### C. Logout User

* **Endpoint:** `POST /auth/logout`
* **Description:** Invalidates the user's session token.
* **Authorization:** Requires Authorization: Bearer <token> header.
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "message": "Logout successful"
    }
    ```

---

### 2. Categories Management

#### A. Get All Categories

* **Endpoint**: `GET /categories`
* **Description**: Retrieves all available item categories.
* **Response (Success - HTTP 200 OK)**:

  ```json
  {
    "categories": [
      {
        "id": 1,
        "name": "Electronics",
        "description": "Electronic devices and gadgets"
      },
      {
        "id": 2,
        "name": "Books",
        "description": "Books and educational materials"
      },
      {
        "id": 3,
        "name": "Clothing",
        "description": "Apparel and accessories"
      }
    ]
  }
  ```

---

### 3. Item Management

#### A. List an Item

* **Endpoint:** `POST /items`
* **Description:** Allows a logged-in user to list an item for swapping.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**

    ```json
    {
      "category_id": 2,
      "name": "The Lord of the Rings Book Set",
      "description": "Hardcover trilogy, good condition.",
      "condition": "used", // good - new
      "estimated_value": 2500.00,
      "currency": "INR",
      "image_urls": ["https://example.com/photo1.jpg", "https://example.com/photo2.jpg"],
    }
    ```

* **Response (Success - HTTP 201 Created):**

    ```json
    {
      "message": "Item listed successfully",
      "item_id": 12345
    }
    ```

#### B. Get All Items (Browse Nearby Listings)

* **Endpoint:** `GET /items`
* **Description:** Retrieves a list of available items, optionally filtered by location or category.
* **Authorization:** Optional (can be browsed by guests or logged-in users).
* **Request (Query Parameters):**
  * `?latitude=22.3072&longitude=73.1812&radius=10`  (radius in km)
  * `?category_id=1`
  * `?search=smartphone`
  * `?condition=new,like_new`
  * `?status=available`

* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "items": [
        {
          "id": 12345,
          "name": "Wireless Headphones",
          "description": "Noise-cancelling, used once.",
          "category": {
            "id": 1,
            "name": "Electronics"
          },
          "condition": "like new",
          "estimated_value": 5000.00,
          "currency": "INR",
          "image_urls": ["https://example.com/photo1.jpg"],
          "owner": {
            "id": 67890,
            "name": "Jane Smith",
            "avatar_url": "https://example.com/avatar.jpg",
            "rating_avg": 4.8,
            "total_ratings": 12,
          },
          "status": "available",
          "created_at": "2025-07-01T10:00:00Z",
        },
        {},{}
      ],
      "total_results": 3
    }
    ```

#### C. Get a Single Item

* **Endpoint:** `GET /items/{item_id}`
* **Description:** Retrieves details for a specific item listing.
* **Authorization:** Optional.
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "id": 12345,
      "name": "Wireless Headphones",
      "description": "Noise-cancelling, used once.",
      "category": {
        "id": 1,
        "name": "Electronics",
        "description": "Electronic devices and gadgets"
      },
      "condition": "like new",
      "estimated_value": 5000.00,
      "currency": "INR",
      "image_urls": ["https://example.com/photo1.jpg", "https://example.com/photo2.jpg"],
      "owner": {
        "id": 67890,
        "name": "Jane Smith",
        "avatar_url": "https://example.com/avatar.jpg",
        "rating_avg": 4.8,
        "total_ratings": 12,
        "total_swaps": 8,
        "city": "Vadodara",
        "state": "Gujarat"
      },
      "status": "available",
      "created_at": "2025-07-01T10:00:00Z",
      "updated_at": "2025-07-01T10:00:00Z"
    }
    ```

#### D. Update Item

* **Endpoint:** `PUT /items/{item_id}`
* **Description:** Updates an existing item (only by owner).
* **Authorization:** Requires Authorization: Bearer <token> header.
* **Request body:**

    ```json
    {
      "name": "Updated Item Name",
      "description": "Updated description",
      "condition": "fair",
      "estimated_value": 2000.00,
      "status": "available"
    }
    ```

* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "message": "Item updated successfully",
      "item_id": 12345
    }
    ```

#### E. Delete Item

* **Endpoint:** `DELETE /items/{item_id}`
* **Description:** Deletes an item (only by owner).
* **Authorization:** Requires Authorization: Bearer <token> header.
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "message": "Item deleted successfully"
    }
    ```

---

### 4. Swap Management

#### A. Initiate a Swap Request

* **Endpoint:** `POST /swap-requests`
* **Description:** Allows a user to propose a swap for another user's item.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**

    ```json
    {
      "requested_item_id": 12345,
      "offered_items": [67890, 11111]
    }
    ```

* **Response (Success - HTTP 201 Created):**

    ```json
    {
      "message": "Swap request initiated successfully",
      "swap_request_id": 98765,
      "status": "pending",
      "expires_at": "2025-07-08T10:00:00Z"
    }
    ```

#### B. Get Swap Requests

* **Endpoint:** `GET /swap-requests`
* **Description:** Retrieves swap requests for the authenticated user.
* **Authorization:** Requires Authorization: Bearer <token> header.
* **Request (Query Parameters):**
  * `?type=sent` (requests sent by user)
  * `?type=received` (requests received by user)
  * `?status=pending,accepted`
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "swap_requests": [
        {
          "id": 98765,
          "requester": {
            "id": 11111,
            "name": "Alice Johnson",
            "avatar_url": "https://example.com/avatar.jpg",
            "rating_avg": 4.2
          },
          "owner": {
            "id": 67890,
            "name": "Jane Smith",
            "avatar_url": "https://example.com/avatar.jpg",
            "rating_avg": 4.8
          },
          "requested_item": {
            "id": 12345,
            "name": "Wireless Headphones",
            "image_urls": ["https://example.com/photo1.jpg"]
          },
          "offered_items": [
            {
              "id": 67890,
              "name": "Bluetooth Speaker",
              "image_urls": ["https://example.com/photo2.jpg"]
            }
          ],
          "status": "pending",
          "created_at": "2025-07-01T10:00:00Z",
          "expires_at": "2025-07-08T10:00:00Z"
        }
      ]
    }
    ```

#### C. Accept/Decline Swap Request

* **Endpoint:** `PUT /swap-requests/{swap_request_id}`
* **Description:** Allows the recipient of a swap request to accept, reject, or cancel it.
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
      "swap_request_id": 98765,
      "status": "accepted" // or "declined"
    }
    ```

#### D. Complete Swap Transaction

* **Endpoint:** `POST /swap-requests/{swap_request_id}/complete`
* **Description:** Marks a swap transaction as completed
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "message": "Swap transaction completed",
      "swap_transaction_id": 55555,
      "status": "completed" // or "declined"
    }
    ```

---

## 5. Chat Management

### A. Get Chat Rooms

* **Endpoint:** `GET /chat-rooms`
* **Description:** Retrieves all chat rooms for the authenticated user.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "chat_rooms": [
        {
          "id": 33333,
          "other_user": {
            "id": 67890,
            "name": "Jane Smith",
            "avatar_url": "https://example.com/avatar.jpg"
          },
          "item": {
            "id": 12345,
            "name": "Wireless Headphones",
            "image_urls": ["https://example.com/photo1.jpg"]
          },
          "swap_request_id": 98765,
          "last_message": {
            "message": "Is this still available?",
            "created_at": "2025-07-01T15:30:00Z"
          },
          "is_active": true,
          "last_message_at": "2025-07-01T15:30:00Z"
        }
      ]
    }
    ```

### B. Get Chat Messages

* **Endpoint:** `GET /chat-rooms/{chat_room_id}/messages`
* **Description:** Retrieves messages from a specific chat room.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Query Parameters:**
  * `?page=1&limit=50`
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "messages": [
        {
          "id": 77777,
          "sender": {
            "id": 11111,
            "name": "Alice Johnson",
            "avatar_url": "https://example.com/avatar.jpg"
          },
          "message": "Is this still available?",
          "created_at": "2025-07-01T15:30:00Z"
        }
      ],
      "page": 1,
      "limit": 50,
      "total_pages": 1
    }
    ```

### C. Send Message

* **Endpoint:** `POST /chat-rooms/{chat_room_id}/messages`
* **Description:** Sends a message in a chat room.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**

    ```json
    {
      "message": "Yes, it's still available. Would you like to make an offer?"
    }
    ```

* **Response (Success - HTTP 201 Created):**

    ```json
    {
      "message": "Message sent successfully",
      "message_id": 88888
    }
    ```

### D. Create Chat Room

* **Endpoint:** `POST /chat-rooms`
* **Description:** Creates a new chat room for discussing an item.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**

    ```json
    {
      "other_user_id": 67890,
      "item_id": 12345
    }
    ```

* **Response (Success - HTTP 201 Created):**

    ```json
    {
      "message": "Chat room created successfully",
      "chat_room_id": 33333
    }
    ```

---

## 6. Rating Management

### A. Submit a Rating

* **Endpoint:** `POST /ratings`
* **Description:** Allows a user to rate their experience after a completed swap.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**

    ```json
    {
      "swap_transaction_id": 55555,
      "rated_user_id": 67890,
      "rating": 5,
      "review": "Smooth trade, item exactly as described!"
    }
    ```

* **Response (Success - HTTP 201 Created):**

    ```json
    {
      "message": "Rating submitted successfully",
      "rating_id": 99999
    }
    ```

### B. Get User Ratings

* **Endpoint:** `GET /users/{user_id}/ratings`
* **Description:** Retrieves ratings for a specific user.
* **Authorization:** Optional.
* **Query Parameters:**
  * `?page=1&limit=20`
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "ratings": [
        {
          "id": 99999,
          "rater": {
            "id": 11111,
            "name": "Alice Johnson",
            "avatar_url": "https://example.com/avatar.jpg"
          },
          "rating": 5,
          "review": "Smooth trade, item exactly as described!",
          "created_at": "2025-07-01T16:00:00Z"
        }
      ],
      "average_rating": 4.8,
      "total_ratings": 12,
      "page": 1,
      "limit": 20,
      "total_pages": 1
    }
    ```

---

## 7. User Profile Management

### A. Get User Profile

* **Endpoint:** `GET /users/{user_id}`
* **Description:** Retrieves public profile information for a user.
* **Authorization:** Optional.
* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "id": 67890,
      "name": "Jane Smith",
      "avatar_url": "https://example.com/avatar.jpg",
      "bio": "Love trading books and electronics!",
      "city": "Vadodara",
      "state": "Gujarat",
      "rating_avg": 4.8,
      "total_ratings": 12,
      "total_swaps": 8,
      "created_at": "2025-06-01T10:00:00Z"
    }
    ```

### B. Update User Profile

* **Endpoint:** `PUT /users/profile`
* **Description:** Updates the authenticated user's profile.
* **Authorization:** Requires `Authorization: Bearer <token>` header.
* **Request Body:**

    ```json
    {
      "name": "Jane Smith Updated",
      "bio": "Updated bio",
      "phone": "+91-9876543210",
      "avatar_url": "https://example.com/new-avatar.jpg",
      "city": "Ahmedabad",
      "state": "Gujarat",
      "location_lat": 23.0225,
      "location_lng": 72.5714,
      "search_radius_km": 50
    }
    ```

* **Response (Success - HTTP 200 OK):**

    ```json
    {
      "message": "Profile updated successfully"
    }
    ```

---

## 8. Error Responses

All endpoints return consistent error responses:

### Validation Error (HTTP 400 Bad Request)

```json
{
  "error": "Validation failed",
  "details": {
    "email": ["Email is required"],
    "password": ["Password must be at least 6 characters"]
  }
}
```

### Unauthorized (HTTP 401 Unauthorized)

```json
{
  "error": "Unauthorized",
  "message": "Invalid or expired token"
}
```

### Forbidden (HTTP 403 Forbidden)

```json
{
  "error": "Forbidden",
  "message": "You don't have permission to access this resource"
}
```

### Not Found (HTTP 404 Not Found)

```json
{
  "error": "Not found",
  "message": "The requested resource was not found"
}
```

### Internal Server Error (HTTP 500 Internal Server Error)

```json
{
  "error": "Internal server error",
  "message": "An unexpected error occurred"
}
```
