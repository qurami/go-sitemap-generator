package stm

const (
	MaxSitemapLinks = 50000
	MaxSitemapNews = 1000
	MaxSitemapFilesize = 50000000 // bytes
)

var (
	// IndexXMLHeader exists for create sitemap xml as a specific sitemap document.
	IndexXMLHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>
      <sitemapindex xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">`)
	// IndexXMLFooter and IndexXMLHeader will used from user together .
	IndexXMLFooter = []byte("</sitemapindex>")
)

var (
	// XMLHeader exists for create sitemap xml as a specific sitemap document.
	XMLHeader = []byte(`<?xml version="1.0" encoding="UTF-8"?>
      <urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9" xmlns:xhtml="http://www.w3.org/1999/xhtml">`)
	// XMLFooter and XMLHeader will used from user together .
	XMLFooter = []byte("</urlset>")
)
