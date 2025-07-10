# Database Schema Documentation

## Overview

This document describes the database schema for a item swapping platform. The schema supports user registration, item listing, swap requests, chat functionality, ratings, and notifications.

## Core Tables

# Database Entities

## USERS

## CATEGORIES

## ITEMS

## SWAP_REQUESTS

## SWAP_REQUEST_ITEMS

## SWAP_TRANSACTIONS

## SWAP_TRANSACTION_ITEMS

## CHAT_ROOMS

## CHAT_MESSAGES

## USER_RATINGS

## USER_SESSIONS

### 1.Users Table

Stores user account information and profile data.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique user identifier |
| email | string | UNIQUE, NOT NULL | User's email address |
| password_hash | string | NOT NULL | Hashed password (not returned in JSON) |
| name | string | NOT NULL | User's display name |
| phone | string | NULLABLE | Phone number |
| avatar_url | string | NULLABLE | Profile picture URL |
| bio | string | NULLABLE | User biography |
| location_lat | decimal(10,8) | NULLABLE | Latitude coordinate |
| location_lng | decimal(11,8) | NULLABLE | Longitude coordinate |
| city | string | NULLABLE | City name |
| state | string | NULLABLE | State/province |
| search_radius_km | int | DEFAULT 25 | Search radius in kilometers |
| rating_avg | decimal(3,2) | DEFAULT 0.0 | Average user rating |
| total_ratings | int | DEFAULT 0 | Total number of ratings received |
| total_swaps | int | DEFAULT 0 | Total completed swaps |
| created_at | timestamp | AUTO_CREATE | Account creation time |
| updated_at | timestamp | AUTO_UPDATE | Last update time |
| deleted_at |timestamp  | NULLABLE, INDEXED| Soft delete timestamp |

**Relationships:**

- Has many Items (as owner)
- Has many SwapRequests (as requester)
- Has many SwapRequests (as owner)
- Has many UserRatings (as rater)
- Has many UserRatings (as rated user)
- Has many ChatMessages (as sender)
- Has many ChatRooms (as user1)
- Has many ChatRooms (as user2)
- Has many SwapTransactions (as user1)
- Has many SwapTransactions (as user2)
- Has many SwapTransactionItems (as item owner)

### 2.Categories Table

Defines item categories for organization.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique category identifier |
| name | string | UNIQUE, NOT NULL | Category name |
| description | string | NULLABLE | Category description |
| created_at | timestamp | AUTO_CREATE | Creation time |

**Relationships:**

- Has many Items

### 3.Items Table

Stores items available for swapping.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique item identifier |
| owner_id | uint | NOT NULL, INDEXED, FK | Reference to Users table |
| category_id | uint | NOT NULL, INDEXED, FK | Reference to Categories table |
| name | string | NOT NULL | Item name |
| description | string | NOT NULL | Item description |
| condition | varchar(50) | CHECK constraint | Item condition (new, like_new, good, fair, poor) |
| estimated_value | decimal(10,2) | | Estimated item value |
| currency | varchar(3) | DEFAULT 'INR' | Currency code |
| image_urls | text[] | JSON serialized | Array of image URLs |
| status | varchar(20) | DEFAULT 'available', INDEXED | Item status (available, pending, swapped) |
| created_at | timestamp | AUTO_CREATE, INDEXED | Creation time |
| updated_at | timestamp | AUTO_UPDATE | Last update time |
| swapped_at | timestamp | NULLABLE | Swap completion time |

**Relationships:**

- Belongs to User (owner)
- Belongs to Category
- Has many SwapRequestItems
- Has many SwapRequests (as requested item)
- Has many ChatRooms (for discussions)
- Has many SwapTransactionItems (when traded)

### 3.Swap Requests Table

Manages swap requests between users.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique request identifier |
| requester_id | uint | NOT NULL, INDEXED, FK | User making the request |
| owner_id | uint | NOT NULL, INDEXED, FK | Owner of requested item |
| requested_item_id | uint | NOT NULL, INDEXED, FK | Item being requested |
| status | varchar(20) | DEFAULT 'pending', INDEXED | Request status (pending, accepted, rejected, cancelled, completed) |
| created_at | timestamp | AUTO_CREATE | Request creation time |
| updated_at | timestamp | AUTO_UPDATE | Last update time |
| expires_at | timestamp | INDEXED | Request expiration time |

**Relationships:**

- Belongs to User (requester)
- Belongs to User (owner)
- Belongs to Item (requested item)
- Has many SwapRequestItems
- Has many ChatRooms
- Has one SwapTransaction

### 4.Swap Request Items Table

Links swap requests to multiple items (many-to-many relationship).

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique identifier |
| swap_request_id | uint | NOT NULL, INDEXED, FK | Reference to SwapRequests |
| item_id | uint | NOT NULL, INDEXED, FK | Reference to Items |
| item_type | varchar(20) | CHECK constraint | Type of item (requested, offered) |
| created_at | timestamp | AUTO_CREATE | Creation time |

**Relationships:**

- Belongs to SwapRequest
- Belongs to Item

### 5.Swap Transactions Table

Tracks confirmed swaps and their completion status.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique transaction identifier |
| swap_request_id | uint | NOT NULL, UNIQUE, FK | Reference to SwapRequest |
| user1_id | uint | NOT NULL, INDEXED, FK | First user in swap |
| user2_id | uint | NOT NULL, INDEXED, FK | Second user in swap |
| created_at | timestamp | AUTO_CREATE | Transaction creation time |
| completed_at | timestamp | NULLABLE | Completion time |

**Relationships:**

- Belongs to SwapRequest (one-to-one)
- Belongs to User (user1)
- Belongs to User (user2)
- Has many SwapTransactionItems
- Has many UserRatings

### 6.Swap Transaction Items Table

Links swap transactions to multiple items involved in the trade.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique identifier |
| swap_transaction_id | uint | NOT NULL, INDEXED, FK | Reference to SwapTransactions |
| item_id | uint | NOT NULL, INDEXED, FK | Reference to Items |
| user_id | uint | NOT NULL, INDEXED, FK | Owner of the item |
| created_at | timestamp | AUTO_CREATE | Creation time |

**Relationships:**

- Belongs to SwapTransaction
- Belongs to Item
- Belongs to User (item owner)

### 7.Chat Rooms Table

Manages chat sessions between users for swap requests or general item inquiries.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique chat room identifier |
| swap_request_id | uint | NULLABLE, INDEXED, FK | Reference to SwapRequest (optional) |
| item_id | uint | NULLABLE, INDEXED, FK | Reference to Item (optional) |
| user1_id | uint | NOT NULL, INDEXED, FK | First user in chat |
| user2_id | uint | NOT NULL, INDEXED, FK | Second user in chat |
| is_active | boolean | DEFAULT true | Chat room status |
| created_at | timestamp | AUTO_CREATE | Creation time |
| last_message_at | timestamp | DEFAULT CURRENT_TIMESTAMP | Last message time |

**Relationships:**

- Belongs to SwapRequest (optional)
- Belongs to Item (optional)
- Belongs to User (user1)
- Belongs to User (user2)
- Has many ChatMessages

### 8.Chat Messages Table

Stores individual messages within chat rooms.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique message identifier |
| chat_room_id | uint | NOT NULL, INDEXED, FK | Reference to ChatRoom |
| sender_id | uint | NOT NULL, INDEXED, FK | Message sender |
| message | string | NOT NULL | Message content |
| created_at | timestamp | AUTO_CREATE, INDEXED | Message creation time |

**Relationships:**

- Belongs to ChatRoom
- Belongs to User (sender)

### 9.User Ratings Table

Stores ratings and reviews between users after completed swaps.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique rating identifier |
| swap_transaction_id | uint | NOT NULL, INDEXED, FK | Reference to SwapTransaction |
| rater_id | uint | NOT NULL, INDEXED, FK | User giving the rating |
| rated_user_id | uint | NOT NULL, INDEXED, FK | User receiving the rating |
| rating | int | CHECK (1-5) | Numeric rating (1-5 stars) |
| review | string | NULLABLE | Optional text review |
| created_at | timestamp | AUTO_CREATE | Rating creation time |

**Relationships:**

- Belongs to SwapTransaction
- Belongs to User (rater)
- Belongs to User (rated user)

### 10.User Sessions Table

Manages user authentication sessions.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique session identifier |
| user_id | uint | NOT NULL, INDEXED, FK | Reference to Users table |
| token_hash | string | NOT NULL | Hashed session token |
| expires_at | timestamp | NOT NULL, INDEXED | Session expiration time |
| created_at | timestamp | AUTO_CREATE | Session creation time |

**Relationships:**

- Belongs to User

-------------------------------------------------------------------------------------------------------------

## Request/Response Models

### Authentication Models

- `LoginRequest`: Email and password for user login
- `RegisterRequest`: Email, password, and name

### Item Models

- `ItemCreateRequest`: All necessary fields for creating a new item listing
- `ItemUpdateRequest`: Fields for updating existing items
- `ItemResponse`: Item data with owner information

### Swap Models

- `SwapRequestCreateRequest`: Create new swap request
- `SwapRequestUpdateRequest`: Update swap request status
- `SwapRequestResponse`: Swap request with item details

### Chat Models

- `ChatMessageRequest`: Send new message
- `ChatMessageResponse`: Message with sender info

### Rating Models

- `UserRatingRequest`: Submit rating and review
- `UserRatingResponse`: Rating with user details

### Response Models

- `APIResponse`: Standard API response wrapper
- `WSMessage`: WebSocket message structure

## Database Constraints and Business Rules

### User Constraints

- Email must be unique
- Password must be at least 6 characters
- Search radius defaults to 25km
- Ratings are calculated as averages

### Item Constraints

- Items have 5 condition levels: new, like_new, good, fair, poor
- Items have 3 status levels: available, pending, swapped
- Items inherit location from their owner (Users table)

### Swap Request Constraints

- Requests expire after 7 days by default
- Status progression: pending → accepted/rejected → completed
- Users cannot request their own items
- Multiple items can be requested/offered per swap request

### Rating Constraints

- Ratings must be between 1-5 stars
- Ratings are tied to completed swap transactions
- Users can only rate each other once per transaction

### Chat Constraints

- Chat rooms can exist independently or be linked to swap requests/items
- Messages support text only
- Chat rooms remain active until manually closed
- Users can chat about items before making swap requests

## Indexes for Performance

- User email (unique)
- Item owner_id, category_id, status
- Swap request status, expiration time
- Chat message timestamps
- SwapRequestItems: swap_request_id, item_id
- SwapTransactionItems: swap_transaction_id
- ChatRooms: swap_request_id, item_id
- Various foreign key relationships

## Soft Delete Support

- Users table includes soft delete functionality via `deleted_at` timestamp
- Allows for data retention while hiding deleted accounts
