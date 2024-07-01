# FoodieBuddie - Food Delivery E-Commerce Platform

A comprehensive food delivery platform built with Go, featuring multi-user roles, secure authentication, and real-time order management.

## 🚀 Features

### User Features
- Authentication with JWT and OTP verification
- Profile and address management
- Shopping cart functionality
- Favorites management
- Order placement and tracking
- Payment integration with Razorpay
- Coupon redemption system
- Search functionality for dishes and sellers

### Seller Features
- Seller authentication and profile management
- Product/dish management (CRUD operations)
- Order management and status updates
- Sales reporting with time-based filters
- Offer management
- Image upload for dishes

### Admin Features
- Category management
- Seller verification and management
- User management (block/unblock)
- Coupon management
- Platform-wide monitoring

## 🛠 Technology Stack

- **Backend**: Go with Fiber framework
- **Database**: PostgreSQL with GORM
- **Authentication**: JWT
- **Documentation**: Swagger/OpenAPI
- **Payment Gateway**: Razorpay
- **SMS Service**: Twilio
- **Image Storage**: Cloudinary
- **Configuration**: Viper

## 📝 API Documentation

### Common Endpoints (No Authentication)
- **Categories**
  - GET `/categories` - List all categories
  - GET `/categories/{id}` - Get specific category

- **Dishes**
  - GET `/dishes` - List dishes with pagination and filters
    - Query params: `p` (page), `l` (limit), `category`, `seller`
  - GET `/dishes/{id}` - Get specific dish details

- **Sellers**
  - GET `/user/sellers` - List sellers with pagination
    - Query params: `p` (page), `l` (limit)
  - GET `/user/sellers/{id}` - Get specific seller details

- **Authentication**
  - POST `/signup` - User registration
    - Required: email, firstName, phone
  - POST `/verifyOtp` - Verify OTP
  - POST `/seller/login` - Seller login
  - POST `/seller/register` - Seller registration
  - POST `/admin/login` - Admin login

### Protected Endpoints

#### 🛒 User Routes (Bearer Token Required)
- **Cart Management**
  - GET `/cart` - View cart
  - POST `/addToCart/{id}` - Add item to cart
  - DELETE `/cart/{id}/deleteItem` - Remove item from cart
  - POST `/cart/checkout` - Place order

- **Orders**
  - GET `/orders` - List all orders
  - GET `/orders/{id}` - Get order details
  - POST `/orders/cancel/{id}` - Cancel order

- **Favorites**
  - GET `/favourites` - View favorite items
  - POST `/addToFavourite/{id}` - Add to favorites
  - DELETE `/favourites/{id}/delete` - Remove from favorites

#### 🏪 Seller Routes (Bearer Token Required)
- **Dish Management**
  - POST `/seller/addDish` - Add new dish (multipart/form-data)
  - GET `/seller/dishes` - List all dishes
    - Query params: `category`
  - PUT `/seller/dishes/{id}` - Update dish
  - DELETE `/seller/dishes/{id}` - Delete dish

- **Order Management**
  - GET `/seller/orders` - View all orders
  - GET `/seller/orders/{id}` - Get order details
  - PUT `/seller/orders/{id}/status` - Update order status
    - Status options: COOKING, FOOD READY, DELIVERED

- **Sales & Offers**
  - GET `/seller/sales` - Get sales report
    - Query params: `filter` (time intervals)
  - POST `/seller/offers/addOffer` - Create new offer
  - GET `/seller/offers` - List all offers

#### 👑 Admin Routes (Bearer Token Required)
- **Category Management**
  - POST `/admin/categories/addCategory` - Add category
  - PATCH `/admin/categories/{id}/edit` - Update category

- **User Management**
  - GET `/admin/users` - List all users
  - PATCH `/admin/users/{id}/block` - Block user
  - PATCH `/admin/users/{id}/unblock` - Unblock user

- **Seller Management**
  - GET `/admin/sellers` - List all sellers
  - PATCH `/admin/sellers/{id}/verify` - Verify seller

- **Coupon Management**
  - GET `/admin/coupons` - List all coupons
  - POST `/admin/coupons/add` - Create coupon
  - PATCH `/admin/coupons/{id}` - Update coupon status

## 🚀 Getting Started

1. Clone the repository
```bash
git clone https://github.com/abdullahnettoor/food-delivery-eCommerce.git
```

2. Install dependencies
```bash
make deps
```

3. Set up environment variables (.env)
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=foodiebuddie

CLOUDINARY_CLOUD_NAME=your_cloud_name
CLOUDINARY_API_KEY=your_api_key
CLOUDINARY_API_SECRET=your_api_secret

TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_SERVICE_SID=your_service_sid

RAZORPAY_KEY_ID=your_key_id
RAZORPAY_SECRET=your_secret
```

4. Run the application
```bash
make run
```

For development with hot reload:
```bash
make nodemon
```

## 📚 Documentation

- Generate Swagger documentation:
```bash
make swag
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.
