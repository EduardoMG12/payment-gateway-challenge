# 🌐 PayGateway Frontend

<div align="center">

**Modern React application for PayGateway payment system**

*Interactive user interface built with React, TypeScript, and modern web technologies*

[![React](https://img.shields.io/badge/React-18+-61DAFB?style=flat-square&logo=react&logoColor=black)](https://reactjs.org/)
[![TypeScript](https://img.shields.io/badge/TypeScript-5+-3178C6?style=flat-square&logo=typescript&logoColor=white)](https://www.typescriptlang.org/)
[![Vite](https://img.shields.io/badge/Vite-4+-646CFF?style=flat-square&logo=vite&logoColor=white)](https://vitejs.dev/)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind-CSS-38B2AC?style=flat-square&logo=tailwind-css&logoColor=white)](https://tailwindcss.com/)

</div>

## 📋 Overview

This is the **frontend application** for the PayGateway system, providing an intuitive and responsive user interface for managing accounts, cards, and transactions. Built with modern React patterns and optimized for performance and user experience.

---

## 📋 Table of Contents

<details>
<summary><strong>🌟 Features & Pages</strong></summary>

### Core Pages

#### 🏠 **Homepage**
- Project overview and introduction
- Technology showcase
- Interactive API documentation
- Downloadable Postman collection
- System architecture visualization

#### 👤 **Account Management**
- **Create Account** - User registration with validation
- **Account Dashboard** - User profile and account details
- Real-time account balance display

#### 💳 **Card Operations**
- **Create Virtual Cards** - Generate secure payment tokens
- **Card Management** - View and manage multiple cards
- **Copy Token** - One-click token copying with feedback
- Card-specific transaction filtering

#### 💰 **Transaction Processing**
- **Process Transactions** - Support for:
  - 🛒 **Purchase** - Debit transactions
  - 💵 **Deposit** - Credit transactions
  - 🔄 **Refund** - Transaction reversals
- **Real-time Flow Visualization** - Interactive transaction flow modal
- **Form Validation** - Client-side validation with error handling

#### 📊 **Financial Reporting**
- **Transaction History** - Comprehensive transaction listing
- **Account Statements** - Detailed financial reports
- **Balance Tracking** - Real-time balance updates
- **Filter & Search** - Filter by card, date, type
- **Export Options** - Data export capabilities

### UI/UX Features

#### 🎨 **Design System**
- **Modern UI Components** - Built with shadcn/ui
- **Responsive Design** - Mobile-first approach
- **Dark/Light Theme** - Theme switching capability
- **Consistent Typography** - Professional design language

#### ⚡ **Interactive Elements**
- **Real-time Updates** - Live transaction status
- **Loading States** - Skeleton loaders and spinners
- **Success/Error Feedback** - Toast notifications
- **Smooth Animations** - Framer Motion animations

#### 🔒 **User Experience**
- **Form Validation** - Real-time input validation
- **Error Handling** - Comprehensive error messages
- **Navigation** - Intuitive routing and breadcrumbs
- **Accessibility** - WCAG compliant components

</details>

<details>
<summary><strong>🛠️ Technology Stack</strong></summary>

### Core Framework
- **React 18** - Modern UI library with concurrent features
- **TypeScript 5+** - Type-safe development
- **Vite 4+** - Fast build tool and dev server

### Styling & UI
- **Tailwind CSS** - Utility-first CSS framework
- **shadcn/ui** - High-quality component library
- **Radix UI** - Accessible primitive components
- **Lucide Icons** - Beautiful icon library
- **Framer Motion** - Animation library

### State Management & Data
- **Zustand** - Lightweight state management
- **React Query** - Server state management
- **React Hook Form** - Performant form handling
- **Zod** - Schema validation

### Routing & Navigation
- **React Router** - Declarative routing
- **React Helmet Async** - Document head management for SEO

### Development Tools
- **ESLint** - Code linting
- **Prettier** - Code formatting
- **PostCSS** - CSS processing
- **Hot Toast** - Toast notifications

### Build & Deployment
- **Vite Bundle Analyzer** - Bundle analysis
- **Docker Support** - Containerized deployment
- **Environment Variables** - Configuration management

</details>

<details>
<summary><strong>🚀 Getting Started</strong></summary>

### Prerequisites

- **Node.js 18+** 
- **npm, yarn, or pnpm**
- **PayGateway Backend Services** (Go API & Rust Processor)

### Quick Start

```bash
# Navigate to frontend directory
cd frontend

# Install dependencies (choose one)
npm install
# or
yarn install
# or
pnpm install

# Start development server
npm run dev
# or
yarn dev
# or  
pnpm dev
```

The application will be available at **http://localhost:8081**

### Environment Configuration

Create a `.env.local` file in the frontend directory:

```env
# API Configuration
VITE_API_BASE_URL=http://localhost:8080

# Development flags
VITE_APP_ENV=development
VITE_ENABLE_DEVTOOLS=true
```

### Docker Development

```bash
# Build and run with Docker
docker build -t paygateway-frontend .
docker run -p 8081:8081 paygateway-frontend

# Or use Docker Compose
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up frontend
```

</details>

<details>
<summary><strong>📁 Project Structure</strong></summary>

```
frontend/
├── 📁 public/                   # Static assets
│   ├── favicon.ico             # Application favicon
│   ├── robots.txt              # SEO robots file
│   └── *.postman_collection.json # API collection
│
├── 📁 src/                      # Source code
│   ├── 📁 components/          # Reusable components
│   │   ├── 📁 ui/             # Base UI components (shadcn/ui)
│   │   ├── 📁 layout/         # Layout components
│   │   │   ├── Navbar.tsx     # Navigation component
│   │   │   └── Footer.tsx     # Footer component
│   │   └── FlowModal.tsx      # Transaction flow modal
│   │
│   ├── 📁 pages/              # Page components
│   │   ├── HomePage.tsx       # Landing page
│   │   ├── CreateAccountPage.tsx # Account creation
│   │   ├── CreateCardPage.tsx # Card management
│   │   ├── TransactionsPage.tsx # Transaction processing
│   │   ├── StatementPage.tsx  # Financial statements
│   │   └── NotFound.tsx       # 404 error page
│   │
│   ├── 📁 hooks/              # Custom React hooks
│   │   ├── use-toast.ts       # Toast notifications
│   │   ├── use-mobile.tsx     # Mobile detection
│   │   └── useSEO.tsx         # SEO management
│   │
│   ├── 📁 services/           # API services
│   │   └── api.ts             # HTTP client & API calls
│   │
│   ├── 📁 store/              # State management
│   │   └── userStore.ts       # User state (Zustand)
│   │
│   ├── 📁 lib/                # Utility libraries
│   │   └── utils.ts           # Helper functions
│   │
│   ├── 📁 config/             # Configuration
│   │   └── seo.ts             # SEO configuration
│   │
│   ├── App.tsx                # Root component
│   ├── main.tsx               # Application entry
│   └── index.css              # Global styles
│
├── 📄 index.html               # HTML template
├── 📄 package.json             # Dependencies & scripts
├── 📄 vite.config.ts          # Vite configuration
├── 📄 tailwind.config.ts      # Tailwind configuration
├── 📄 tsconfig.json           # TypeScript configuration
└── 📄 Dockerfile              # Docker configuration
```

</details>

<details>
<summary><strong>🔧 Development</strong></summary>

### Available Scripts

```bash
# Development
npm run dev          # Start development server
npm run build        # Build for production
npm run preview      # Preview production build
npm run lint         # Run ESLint
npm run type-check   # TypeScript type checking

# Testing (if configured)
npm run test         # Run test suite
npm run test:watch   # Run tests in watch mode
npm run test:coverage # Run tests with coverage
```

### Code Quality

#### ESLint Configuration
- React hooks rules
- TypeScript strict rules
- Import/export linting
- Accessibility rules

#### TypeScript Configuration
- Strict mode enabled
- Path mapping configured
- React JSX support
- Modern ES target

### Performance Optimizations

#### Bundle Optimization
- **Code Splitting** - Route-based splitting
- **Tree Shaking** - Unused code elimination
- **Asset Optimization** - Image and font optimization
- **Lazy Loading** - Component lazy loading

#### Runtime Performance
- **Memoization** - React.memo and useMemo
- **Virtual Scrolling** - For large lists
- **Debounced Inputs** - Search and form inputs
- **Optimized Re-renders** - Zustand optimizations

### SEO Implementation

#### Meta Tags Management
- Dynamic page titles
- Meta descriptions
- Open Graph tags
- Twitter Cards
- Canonical URLs

#### Structured Data
- Schema.org markup
- JSON-LD implementation
- Breadcrumb navigation
- Rich snippets support

</details>

<details>
<summary><strong>🔗 API Integration</strong></summary>

### HTTP Client Configuration

The frontend uses a centralized API client with:

- **Axios** for HTTP requests
- **Request/Response Interceptors**
- **Error handling**
- **Loading state management**
- **Retry logic**

### API Endpoints Used

| Service | Endpoint | Purpose |
|---------|----------|---------|
| **Accounts** | `POST /accounts` | Create new account |
| **Accounts** | `GET /accounts/{id}/balance` | Get account balance |
| **Cards** | `POST /cards` | Create new card |
| **Cards** | `GET /cards/{accountId}` | List user cards |
| **Transactions** | `POST /transactions` | Process transaction |
| **Transactions** | `GET /transactions/{accountId}` | Get transaction history |

### Error Handling

- **Network Errors** - Connection timeout handling
- **HTTP Errors** - Status code error mapping
- **Validation Errors** - Form field error display
- **User Feedback** - Toast notifications for all states

</details>

<details>
<summary><strong>🎨 Design System</strong></summary>

### Component Library

Built on **shadcn/ui** with customizations:

#### Base Components
- **Button** - Multiple variants and sizes
- **Input** - Form inputs with validation
- **Card** - Content containers
- **Badge** - Status indicators
- **Dialog** - Modal dialogs
- **Select** - Dropdown selections
- **Table** - Data tables
- **Tabs** - Tab navigation

#### Custom Components
- **FlowModal** - Transaction flow visualization
- **Navbar** - Application navigation
- **Footer** - Site footer
- **Loading States** - Skeleton loaders

### Color System

```css
/* CSS Custom Properties */
--primary: 222.2 84% 4.9%;
--secondary: 210 40% 96%;
--accent: 210 40% 93%;
--muted: 210 40% 96%;
--success: 142 76% 36%;
--warning: 38 92% 50%;
--destructive: 0 84.2% 60.2%;
```

### Typography Scale

- **Headings**: h1-h6 with responsive sizing
- **Body Text**: Optimized line heights
- **Code**: Monospace with syntax highlighting
- **Captions**: Muted text for metadata

</details>

---

## 🚀 Deployment

### Production Build

```bash
# Create production build
npm run build

# Preview production build locally
npm run preview
```

### Docker Deployment

```bash
# Build Docker image
docker build -t paygateway-frontend .

# Run container
docker run -p 8081:8081 paygateway-frontend
```

### Environment Variables

For production deployment, configure:

```env
VITE_API_BASE_URL=https://api.your-domain.com
VITE_APP_ENV=production
VITE_ENABLE_DEVTOOLS=false
```

## 🤝 Contributing

When contributing to the frontend:

1. Follow the established code style
2. Add proper TypeScript types
3. Include responsive design considerations
4. Test on multiple browsers
5. Update documentation as needed

## 📄 License

This frontend application is part of the PayGateway project and follows the same [MIT License](../LICENSE).

---

<div align="center">

**Part of the PayGateway ecosystem**

[🏠 Main Project](../README.md) • [🚀 Go API](../go-api/README.md) • [⚡ Rust Processor](../rust-processor/README.md)

</div>
- **TypeScript**
- **Vite**
- **Tailwind CSS**
- **Zustand** (State Management)
- **React Router**
- **Axios**