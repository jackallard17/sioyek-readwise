package main

var (
	AuthEndpoint       = "https://readwise.io/api/v2/auth/"
	BooksEndpoint      = "https://readwise.io/api/v2/books?page_size=1000&category=books&source=OctoberForKobo"
	CoverEndpoint      = "https://readwise.io/api/v2/books/%d"
	HighlightsEndpoint = "https://readwise.io/api/v2/highlights/"
	MaxHighlightLen    = 8096 // It's actually 8191 but we'll go under the limit anyway
	SourceCategory     = "books"
	SourceType         = "sioyek-readwise"
	UserAgent          = "sioyek-readwise <https://github.com/jackallard17/sioyek-readwise>"
)
