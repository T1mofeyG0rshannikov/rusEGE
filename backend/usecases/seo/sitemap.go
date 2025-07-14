package seo

import (
	"encoding/xml"
	"fmt"
	"os"
	"rusEGE/repositories"
	"time"

	"github.com/joho/godotenv"
)

// Структура для URL в Sitemap
type SitemapURL struct {
	Loc        string  `xml:"loc"`
	Lastmod    string  `xml:"lastmod"`              // ISO 8601
	Changefreq string  `xml:"changefreq,omitempty"` // Optional
	Priority   float32 `xml:"priority,omitempty"`   // Optional (0.0 to 1.0)
}

// Структура для Sitemap
type Sitemap struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []SitemapURL `xml:"url"`
}

func GenerateSitemap(
	tr *repositories.GormTaskRepository,
) ([]byte, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	domain := os.Getenv("DOMAIN")

	var urls []SitemapURL
	
	urls = []SitemapURL{
		{
			Loc:        fmt.Sprintf("https://%s/", domain),
			Lastmod:    time.Now().Format(time.RFC3339),
			Changefreq: "monthly",
			Priority:   1.0,
		},
		{
			Loc:        fmt.Sprintf("https://%s/tasks", domain),
			Lastmod:    time.Now().Add(-time.Hour * 24 * 7).Format(time.RFC3339), // 7 days ago
			Changefreq: "weekly",
			Priority:   0.8,
		},
		{
			Loc:        fmt.Sprintf("https://%s/statistics", domain),
			Lastmod:    time.Now().Add(-time.Hour * 24 * 30).Format(time.RFC3339), // 30 days ago
			Changefreq: "monthly",
			Priority:   0.5,
		},
	}

	tasks, err := tr.All()
	if err != nil{
		return nil, err
	}

	for _, task := range(tasks){
		urls = append(urls, SitemapURL{
			Loc:        fmt.Sprintf("https://%s/task/%s", domain, task.Number),
			Lastmod:    time.Now().Add(-time.Hour * 24).Format(time.RFC3339), // 1 day ago
			Changefreq: "monthly",
			Priority:   0.5,
		})
	}

	sitemap := Sitemap{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs: urls,
	}

	output, err := xml.MarshalIndent(sitemap, "", "  ")
	if err != nil {
		return nil, err
	}

	xmlHeader := []byte(xml.Header)
	output = append(xmlHeader, output...)

	return output, nil
}