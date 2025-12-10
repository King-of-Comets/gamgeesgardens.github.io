package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

type BlogPost struct {
	Filename string
	Title    string    `yaml:"title"`
	Date     time.Time `yaml:"date"`
	Keywords []string  `yaml:"keywords"`
	Body     string    `yaml:"body"`
	Images   []string  `yaml:"images"`
}

type PageData struct {
	Title     string
	BlogPosts []BlogPost
	Page      int
	TotalPages int
}

var templates *template.Template
var blogPosts []BlogPost
var blogPostsDir = "blog_posts"

func init() {
	templates = template.Must(template.New("").Funcs(template.FuncMap{
		"add":      func(a, b int) int { return a + b },
		"subtract": func(a, b int) int { return a - b },
	}).ParseGlob("templates/*.html"))
}

func main() {
	// Load blog posts on startup
	loadBlogPosts()

	// Watch for changes in blog_posts directory
	go watchBlogPosts()

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/blog_posts/", http.StripPrefix("/blog_posts/", http.FileServer(http.Dir("blog_posts"))))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/mission", missionHandler)
	http.HandleFunc("/services", servicesHandler)
	http.HandleFunc("/blog", blogHandler)

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadBlogPosts() {
	blogPosts = []BlogPost{}
	
	err := filepath.WalkDir(blogPostsDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(path, ".yaml") {
			data, err := os.ReadFile(path)
			if err != nil {
				log.Printf("Error reading %s: %v", path, err)
				return nil
			}

			var post BlogPost
			if err := yaml.Unmarshal(data, &post); err != nil {
				log.Printf("Error parsing %s: %v", path, err)
				return nil
			}
			post.Filename = filepath.Base(path)
			blogPosts = append(blogPosts, post)
		}
		return nil
	})

	if err != nil {
		log.Printf("Error walking blog_posts: %v", err)
	}

	// Sort by date descending
	sort.Slice(blogPosts, func(i, j int) bool {
		return blogPosts[i].Date.After(blogPosts[j].Date)
	})

	log.Printf("Loaded %d blog posts", len(blogPosts))
}

func watchBlogPosts() {
	lastModTime := time.Now()
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		info, err := os.Stat(blogPostsDir)
		if err != nil {
			continue
		}
		if info.ModTime().After(lastModTime) {
			log.Println("Blog posts directory changed, reloading...")
			loadBlogPosts()
			lastModTime = info.ModTime()
		}
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Home"}
	templates.ExecuteTemplate(w, "home.html", data)
}

func missionHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Mission"}
	templates.ExecuteTemplate(w, "mission.html", data)
}

func servicesHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Services"}
	templates.ExecuteTemplate(w, "services.html", data)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	page := 0
	if p := r.URL.Query().Get("page"); p != "" {
		if _, err := fmt.Sscanf(p, "%d", &page); err != nil {
			page = 0
		}
	}

	postsPerPage := 10
	totalPages := (len(blogPosts) + postsPerPage - 1) / postsPerPage
	
	if page < 0 {
		page = 0
	}
	if page >= totalPages && totalPages > 0 {
		page = totalPages - 1
	}

	start := page * postsPerPage
	end := start + postsPerPage
	if end > len(blogPosts) {
		end = len(blogPosts)
	}

	var pagePosts []BlogPost
	if start < len(blogPosts) {
		pagePosts = blogPosts[start:end]
	}

	data := PageData{
		Title:      "Blog",
		BlogPosts:  pagePosts,
		Page:       page,
		TotalPages: totalPages,
	}
	templates.ExecuteTemplate(w, "blog.html", data)
}