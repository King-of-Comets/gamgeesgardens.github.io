# Gamgee's Gardens Website

A stylish, environmentally-themed landscaping website with a Go backend that serves static HTML pages with cohesive CSS and dynamic blog post generation.

## Features

- **Home Page**: Hero section with splash image and feature highlights
- **Mission Statement**: Dedicated section with large image and mission content
- **Services Page**: List of landscaping services with contact information
- **Blog**: Dynamic blog posts loaded from YAML files with pagination (10 posts per page)
- **Auto-reload**: Automatically detects new blog posts without server restart

## Project Structure

```
gamgees-gardens/
├── main.go
├── go.mod
├── templates/
│   ├── base.html
│   ├── home.html
│   ├── mission.html
│   ├── services.html
│   └── blog.html
├── static/
│   ├── style.css
│   └── images/
│       ├── hero.jpg (optional)
│       └── mission.jpg (optional)
└── blog_posts/
    ├── spring-planting-tips.yaml
    ├── composting-basics.yaml
    ├── spring-garden.jpg (optional)
    ├── compost-bin.jpg (optional)
    └── rich-compost.jpg (optional)
```

## Setup Instructions

### 1. Initialize Project

```bash
# Create project directory
mkdir gamgees-gardens
cd gamgees-gardens

# Create subdirectories
mkdir templates static static/images blog_posts
```

### 2. Create Files

Copy the following files into your project:

- `main.go` - Main Go server file
- `go.mod` - Go module dependencies
- `templates/base.html` - Base HTML template
- `templates/home.html` - Home page template
- `templates/mission.html` - Mission page template
- `templates/services.html` - Services page template
- `templates/blog.html` - Blog page template
- `static/style.css` - Stylesheet
- `blog_posts/spring-planting-tips.yaml` - Example blog post 1
- `blog_posts/composting-basics.yaml` - Example blog post 2

### 3. Initialize Go Environment

```bash
# Initialize Go module (if not already done)
go mod init gamgees-gardens

# Download dependencies
go mod tidy
```

This will download the required `gopkg.in/yaml.v3` package.

### 4. Add Images (Optional)

For the best experience, add these images:

**Hero/Mission Images:**
- `static/images/hero.jpg` - Garden landscape for home page
- `static/images/mission.jpg` - Sustainable garden for mission page

**Blog Images:**
- `blog_posts/spring-garden.jpg`
- `blog_posts/compost-bin.jpg`
- `blog_posts/rich-compost.jpg`

The site will work without images but will look better with them.

### 5. Run the Server

```bash
# Run the server
go run main.go
```

You should see:
```
Loaded 2 blog posts
Server starting on http://localhost:8080
```

### 6. Access the Website

Open your browser and navigate to:
- **Home**: http://localhost:8080/
- **Mission**: http://localhost:8080/mission
- **Services**: http://localhost:8080/services
- **Blog**: http://localhost:8080/blog

## Adding New Blog Posts

To add a new blog post, create a new YAML file in the `blog_posts/` directory:

```yaml
title: "Your Blog Post Title"
date: 2024-12-07T10:00:00Z
keywords:
  - keyword1
  - keyword2
  - keyword3
images:
  - image1.jpg
  - image2.jpg
body: |
  Your blog post content goes here.
  
  You can use multiple paragraphs.
  
  The content will be displayed as-is.
```

**Important Notes:**
- Save the file with a `.yaml` extension
- Place any referenced images in the `blog_posts/` directory
- The server checks for changes every 5 seconds and auto-reloads
- Blog posts are automatically sorted by date (newest first)
- Pagination activates when there are more than 10 posts

## Building for Production

```bash
# Build the binary
go build -o gamgees-gardens main.go

# Run the binary
./gamgees-gardens
```

## Customization

### Changing Colors

Edit `static/style.css` and modify the CSS variables at the top:

```css
:root {
    --primary: #2d5016;    /* Dark green */
    --secondary: #6b8e23;  /* Olive green */
    --accent: #8fbc8f;     /* Light green */
    --light: #f4f7f0;      /* Off-white */
    --text: #2c3e20;       /* Dark text */
}
```

### Changing Port

Edit `main.go` and change the port in the last line:

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

### Modifying Content

- Edit HTML templates in `templates/` directory
- Update contact information in `templates/services.html`
- Modify mission text in `templates/mission.html`

## Troubleshooting

**Blog posts not loading:**
- Check that YAML files are properly formatted
- Ensure the `blog_posts/` directory exists
- Check server logs for parsing errors

**Images not displaying:**
- Verify image paths match the filenames in YAML
- Ensure images are in the `blog_posts/` directory
- Check file permissions

**Server won't start:**
- Ensure port 8080 is not in use
- Run `go mod tidy` to ensure dependencies are installed
- Check for syntax errors in templates

## License

This project is provided as-is for Gamgee's Gardens landscaping business.