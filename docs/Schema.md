# Database Schema Documentation

## Overview

This document describes the database schema for a item swapping platform. The schema supports user registration, item listing, swap requests, chat functionality, ratings, and notifications.

## Core Tables

### Users Table

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
- Has many Notifications
- Has many UserRatings (as rater)
- Has many UserRatings (as rated user)

### Categories Table

Defines item categories for organization.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique category identifier |
| name | string | UNIQUE, NOT NULL | Category name |
| description | string | NULLABLE | Category description |
| created_at | timestamp | AUTO_CREATE | Creation time |

**Relationships:**

- Has many Items

### Items Table

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
| location_lat | decimal(10,8) | NULLABLE, INDEXED | Latitude coordinate |
| location_lng | decimal(11,8) | NULLABLE, INDEXED | Longitude coordinate |
| created_at | timestamp | AUTO_CREATE, INDEXED | Creation time |
| updated_at | timestamp | AUTO_UPDATE | Last update time |
| swapped_at | timestamp | NULLABLE | Swap completion time |

**Relationships:**

- Belongs to User (owner)
- Belongs to Category
- Has many SwapRequests (as requested item)
- Has many SwapRequests (as offered item)

### Swap Requests Table

Manages swap requests between users.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique request identifier |
| requester_id | uint | NOT NULL, INDEXED, FK | User making the request |
| owner_id | uint | NOT NULL, INDEXED, FK | Owner of requested item |
| requested_item_id | uint | NOT NULL, INDEXED, FK | Item being requested |
| offered_item_id | uint | NOT NULL, INDEXED, FK | Item being offered |
| status | varchar(20) | DEFAULT 'pending', INDEXED | Request status (pending, accepted, rejected, cancelled, completed)
| created_at | timestamp | AUTO_CREATE | Request creation time |
| updated_at | timestamp | AUTO_UPDATE | Last update time |
| expires_at | timestamp | INDEXED | Request expiration time |

**Relationships:**

- Belongs to User (requester)
- Belongs to User (owner)
- Belongs to Item (requested item)
- Belongs to Item (offered item)
- Has one ChatRoom
- Has one SwapTransaction

### Swap Transactions Table

Tracks confirmed swaps and their completion status.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique transaction identifier |
| swap_request_id | uint | NOT NULL, UNIQUE, FK | Reference to SwapRequest |
| user1_id | uint | NOT NULL, INDEXED, FK | First user in swap |
| user2_id | uint | NOT NULL, INDEXED, FK | Second user in swap |
| item1_id | uint | NOT NULL, INDEXED, FK | First item in swap |
| item2_id | uint | NOT NULL, INDEXED, FK | Second item in swap |
| created_at | timestamp | AUTO_CREATE | Transaction creation time |
| completed_at | timestamp | NULLABLE | Completion time |

**Relationships:**

- Belongs to SwapRequest
- Belongs to User (user1)
- Belongs to User (user2)
- Belongs to Item (item1)
- Belongs to Item (item2)
- Has many UserRatings

### Chat Rooms Table

Manages chat sessions between users for swap requests.

| Field | Type | Constraints | Description |
|-------|------|-------------|-------------|
| id | uint | PRIMARY KEY, AUTO_INCREMENT | Unique chat room identifier |
| swap_request_id | uint | NOT NULL, UNIQUE, FK | Reference to SwapRequest |
| user1_id | uint | NOT NULL, INDEXED, FK | First user in chat |
| user2_id | uint | NOT NULL, INDEXED, FK | Second user in chat |
| is_active | boolean | DEFAULT true | Chat room status |
| created_at | timestamp | AUTO_CREATE | Creation time |
| last_message_at | timestamp | DEFAULT CURRENT_TIMESTAMP | Last message time |

**Relationships:**

- Belongs to SwapRequest
- Belongs to User (user1)
- Belongs to User (user2)
- Has many ChatMessages

### Chat Messages Table

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

### User Ratings Table

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

### User Sessions Table

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

## Request/Response Models

### Authentication Models

- `LoginRequest`: Email and password for user login
- `RegisterRequest`: Email, password, and name

### Item Models

- `ItemCreateRequest`: All necessary fields for creating a new item listing

### Response Models

- `APIResponse`: Standard API response wrapper
- `WSMessage`: WebSocket message structure

## Notification Types

- `swap_request`: New swap request received
- `swap_accepted`: Swap request accepted
- `swap_rejected`: Swap request rejected
- `swap_completed`: Swap transaction completed

## Database Constraints and Business Rules

### User Constraints

- Email must be unique
- Password must be at least 6 characters
- Search radius defaults to 25km
- Ratings are calculated as averages

### Item Constraints

- Items have 5 condition levels: new, like_new, good, fair, poor
- Items have 3 status levels: available, pending, swapped
- Location coordinates use decimal precision for accuracy

### Swap Request Constraints

- Requests expire after 7 days by default
- Status progression: pending → accepted/rejected → completed
- Users cannot request their own items

### Rating Constraints

- Ratings must be between 1-5 stars
- Ratings are tied to completed swap transactions
- Users can only rate each other once per transaction

### Chat Constraints

- One chat room per swap request
- Messages support text only
- Chat rooms remain active until manually closed

## Indexes for Performance

- User email (unique)
- Item owner_id, category_id, status, location coordinates
- Swap request status, expiration time
- Chat message timestamps
- Various foreign key relationships

## Soft Delete Support

- Users table includes soft delete functionality via `deleted_at` timestamp
- Allows for data retention while hiding deleted accounts
