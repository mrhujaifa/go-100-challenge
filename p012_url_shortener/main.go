package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

/*
===========================================
  Go 100 Challenge
  Problem: #012
  Level: 🟡 Medium
===========================================

Problem: URL Shortener

Topic:
- map
- string generation
- functions
- struct

Industry Use:
URL Shortening Services / Link Management /
Marketing Tools / Analytics Platforms

Rules (English):
- User can shorten a long URL
- System generates a random 6 character code
- User can:
    * Shorten a URL
    * Expand a short code back to original URL
    * View all shortened URLs
    * Exit
- Invalid URL (no http/https) => "Invalid URL! Must start with http:// or https://"
- Short code not found        => "Short code not found!"
- Same URL shortened twice    => return existing short code

Rules (বাংলা):
- User একটা long URL short করতে পারবে
- System random 6 character এর code generate করবে
- User করতে পারবে:
    * URL short করা
    * Short code দিয়ে original URL দেখা
    * সব shortened URL দেখা
    * Exit করা
- Invalid URL (http/https নেই) => "Invalid URL! Must start with http:// or https://"
- Short code না পেলে           => "Short code not found!"
- Same URL দুইবার দিলে         => existing short code return করবে

Example Run:
  === URL Shortener ===
  1. Shorten URL
  2. Expand URL
  3. View All
  4. Exit

  Choose: 1
  Enter URL: https://www.google.com
  Short code: go.ly/xK9mQz

  Choose: 1
  Enter URL: https://www.google.com
  Already shortened: go.ly/xK9mQz

  Choose: 1
  Enter URL: www.google.com
  Invalid URL! Must start with http:// or https://

  Choose: 2
  Enter short code: go.ly/xK9mQz
  Original URL: https://www.google.com

  Choose: 3
  go.ly/xK9mQz  =>  https://www.google.com

  Choose: 4
  Goodbye!
===========================================
*/

type Url struct {
	originalUrl string
	shortUrl    string
}

var reader = bufio.NewReader(os.Stdin)

func readInput() string {
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func main() {

	urls := &[]Url{}

	var Choice int

	for {
		fmt.Println(`
██╗   ██╗██████╗ ██╗         ███████╗██╗  ██╗ ██████╗ ██████╗ ████████╗
██║   ██║██╔══██╗██║         ██╔════╝██║  ██║██╔═══██╗██╔══██╗╚══██╔══╝
██║   ██║██████╔╝██║         ███████╗███████║██║   ██║██████╔╝   ██║   
██║   ██║██╔══██╗██║         ╚════██║██╔══██║██║   ██║██╔══██╗   ██║   
╚██████╔╝██║  ██║███████╗    ███████║██║  ██║╚██████╔╝██║  ██║   ██║   
 ╚═════╝ ╚═╝  ╚═╝╚══════╝    ╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝   ╚═╝   
`)

		fmt.Println("╔════════════════════════════════════════════════════╗")
		fmt.Println("║              🚀 URL SHORTENER MENU                 ║ ")
		fmt.Println("╠════════════════════════════════════════════════════╣")
		fmt.Println("║                                                    ║")
		fmt.Println("║   [1] Shorten Long URL                             ║")
		fmt.Println("║   [2] Expand Short URL                             ║")
		fmt.Println("║   [3] View All Saved URLs                          ║")
		fmt.Println("║   [4] Exit Application                             ║")
		fmt.Println("║                                                    ║")
		fmt.Println("╚════════════════════════════════════════════════════╝")

		fmt.Print("\n👉 Enter Your Choice (1-4): ")

		fmt.Scanln(&Choice)

		switch Choice {
		case 1:
			shortenLongUrl(urls)
		case 2:
			expandShortUrl(urls)
		case 3:
			displayAllUrls(urls)
		case 4:
			return
		}

	}

}

//  Choose: 1
//   Enter URL: https://www.google.com
//   Short code: go.ly/xK9mQz

func shortenLongUrl(urls *[]Url) {
	fmt.Println("Enter URL: ")
	urlInput := readInput()

	if !strings.HasPrefix(urlInput, "https://") && !strings.HasPrefix(urlInput, "http://") {

		fmt.Println("\n┌──────────────────────────────────────┐")
		fmt.Println("│ ❌ Invalid URL Format                │")
		fmt.Println("├──────────────────────────────────────┤")
		fmt.Println("│ URL must start with:                │")
		fmt.Println("│ • https://                          │")
		fmt.Println("│ • http://                           │")
		fmt.Println("└──────────────────────────────────────┘")

		return
	}

	// 	  Choose: 1
	//   Enter URL: https://www.google.com
	//   Already shortened: go.ly/xK9mQz

	for _, url := range *urls {
		if url.originalUrl == urlInput {
			fmt.Println("\n┌──────────────────────────────────────┐")
			fmt.Println("│ ✨ URL Already Shortened             │")
			fmt.Println("├──────────────────────────────────────┤")
			fmt.Printf("│ 🔗 Long URL:  %s\n", url.originalUrl)
			fmt.Printf("│ ✂️  Short URL: go.ly/%s\n", url.shortUrl)
			fmt.Println("└──────────────────────────────────────┘")
			return // Exit early since it already exists
		}
	}

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const codeLength = 6

	var sb strings.Builder

	sb.Grow(codeLength)

	for i := 0; i < codeLength; i++ {
		randomIndex := rand.IntN(len(charset))
		sb.WriteByte(charset[randomIndex])
	}

	shortCode := sb.String()

	newUrlMapping := Url{
		originalUrl: urlInput,
		shortUrl:    shortCode,
	}
	*urls = append(*urls, newUrlMapping)

	fmt.Println("\n┌──────────────────────────────────────┐")
	fmt.Println("│ 🎉 URL Shortened Successfully!       │")
	fmt.Println("├─┬──────────────────────────────────┤")
	fmt.Printf("│ 🔗 Long URL:  %s\n", urlInput)
	fmt.Printf("│ ✂️  Short URL: go.ly/%s\n", shortCode)
	fmt.Println("└──────────────────────────────────────┘")
}

func expandShortUrl(urls *[]Url) {
	fmt.Print("Enter Short URL or Code: ")
	input := readInput()

	var shortCode string
	if strings.Contains(input, "/") {
		parts := strings.Split(input, "/")
		shortCode = parts[len(parts)-1]
	} else {
		shortCode = input
	}

	// Trim whitespace just in case
	shortCode = strings.TrimSpace(shortCode)

	for _, u := range *urls {
		if u.shortUrl == shortCode {
			fmt.Println("\n┌──────────────────────────────────────┐")
			fmt.Println("│ 🔍 Short Code Found!                 │")
			fmt.Println("├──────────────────────────────────────┤")
			fmt.Printf("│ ✂️  Short URL: go.ly/%s\n", u.shortUrl)
			fmt.Printf("│ 🔗 Long URL:  %s\n", u.originalUrl)
			fmt.Println("└──────────────────────────────────────┘")
			return
		}
	}

	fmt.Println("\n┌──────────────────────────────────────┐")
	fmt.Println("│ ❌ Short Code Not Found              │")
	fmt.Println("├──────────────────────────────────────┤")
	fmt.Printf("│ The code '%s' does not exist      \n", shortCode)
	fmt.Println("│ in our registry data system.         │")
	fmt.Println("└──────────────────────────────────────┘")
}

func displayAllUrls(urls *[]Url) {

	if len(*urls) == 0 {
		fmt.Println("\n┌──────────────────────────────────────┐")
		fmt.Println("│ 📭 Registry Empty                    │")
		fmt.Println("├──────────────────────────────────────┤")
		fmt.Println("│ No URLs have been shortened yet.     │")
		fmt.Println("└──────────────────────────────────────┘")
		return
	}

	fmt.Println("\n┌──────────────────────────────────────┐")
	fmt.Println("│ 📋 Current Shortened URLs            │")
	fmt.Println("├──────────────────────────────────────┤")

	for _, u := range *urls {
		fmt.Printf("│  go.ly/%s  =>  %s\n", u.shortUrl, u.originalUrl)
	}

	fmt.Println("└──────────────────────────────────────┘")
}
