package serve

import (
	"bytes"
	"fkpoe/utls"
	"fmt"
	http "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"os"
	"strings"
)

func forwardRequest(c *fiber.Ctx, url string) error {
	method := string(c.Request().Header.Method())
	ua := c.Request().Header.Peek("User-Agent")
	body := c.Request().Body()
	jar := tls_client.NewCookieJar()
	profile, _ := utls.GetProfileUa(string(ua))
	options := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(120),
		tls_client.WithClientProfile(profile),
		tls_client.WithNotFollowRedirects(),
		tls_client.WithCookieJar(jar),
	}
	client, err := tls_client.NewHttpClient(tls_client.NewNoopLogger(), options...)
	if err != nil {
		log.Printf("Failed to create HTTP client: %v. Retrying...\n", err)
		return err
	}
	contentType := string(c.Request().Header.ContentType())
	var req *http.Request
	if strings.HasPrefix(contentType, "application/octet-stream") ||
		strings.HasPrefix(contentType, "video/") ||
		strings.HasPrefix(contentType, "audio/") ||
		strings.HasPrefix(contentType, "text/event-stream") {
		req, err = http.NewRequest(method, url, bytes.NewReader(body))
	} else {
		req, err = http.NewRequest(method, url, io.NopCloser(bytes.NewReader(body)))
	}
	if err != nil {
		fmt.Printf("Failed to create request: %v\n\n", err)
		return err
	}
	req.Header = utls.GetBrowserFrom(c)
	//fmt.Print(os.Getenv("PB_COOKIE"))
	//if os.Getenv("PB_COOKIE") != "" {
	//	req.Header.Add("Cookie", "p-b:"+os.Getenv("PB_COOKIE"))
	//}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Failed to send request: %v\n", err)
		return err
	}
	defer res.Body.Close()

	cookies := make(map[string]*fiber.Cookie)
	for _, cookie := range res.Cookies() {
		cookies[cookie.Name] = &fiber.Cookie{
			Name:     cookie.Name,
			Value:    cookie.Value,
			Expires:  cookie.Expires,
			Domain:   cookie.Domain,
			Path:     cookie.Path,
			Secure:   cookie.Secure,
			HTTPOnly: cookie.HttpOnly,
		}
	}

	if res.StatusCode == http.StatusMovedPermanently ||
		res.StatusCode == http.StatusFound ||
		res.StatusCode == http.StatusTemporaryRedirect ||
		res.StatusCode == http.StatusPermanentRedirect {

		redirectURL, err := res.Location()
		if err != nil {
			fmt.Printf("Failed to get redirect location: %v\n", err)
			return err
		}
		localURL := utls.LocalBaseURL + redirectURL.Path
		if redirectURL.RawQuery != "" {
			localURL += "?" + redirectURL.RawQuery
		}
		c.Status(res.StatusCode)
		for _, cookie := range cookies {
			c.Cookie(cookie)
		}
		c.Set("Location", localURL)
		return nil
	} else {
		responseContentType := res.Header.Get("Content-Type")
		if strings.HasPrefix(responseContentType, "text/html") {
			tmpFile, err := os.CreateTemp("", "response-*.html")
			if err != nil {
				fmt.Printf("Failed to create temporary file: %v\n", err)
				return err
			}
			defer os.Remove(tmpFile.Name())
			_, err = io.Copy(tmpFile, res.Body)
			if err != nil {
				fmt.Printf("Failed to write response body to file: %v\n", err)
				return err
			}
			tmpFile.Close()
			c.Status(res.StatusCode)
			c.Set("Content-Type", responseContentType)
			for _, cookie := range cookies {
				c.Cookie(cookie)
			}
			return c.SendFile(tmpFile.Name())
		} else if strings.HasPrefix(responseContentType, "application/javascript") {
			tmpFile, err := os.CreateTemp("", "response-*.js")
			if err != nil {
				fmt.Printf("Failed to create temporary file: %v\n", err)
				return err
			}
			defer os.Remove(tmpFile.Name())
			_, err = io.Copy(tmpFile, res.Body)
			if err != nil {
				fmt.Printf("Failed to write response body to file: %v\n", err)
				return err
			}
			tmpFile.Close()

			c.Status(res.StatusCode)
			c.Set("Content-Type", responseContentType)
			for _, cookie := range cookies {
				c.Cookie(cookie)
			}
			return c.SendFile(tmpFile.Name())
		} else if strings.HasPrefix(responseContentType, "image/") ||
			strings.HasPrefix(responseContentType, "video/") ||
			strings.HasPrefix(responseContentType, "audio/") {
			c.Status(res.StatusCode)
			c.Set("Content-Type", responseContentType)
			for _, cookie := range cookies {
				c.Cookie(cookie)
			}
			_, err = io.Copy(c, res.Body)
			if err != nil {
				fmt.Printf("Failed to copy response body: %v\n", err)
				return err
			}
			return nil
		} else {
			c.Status(res.StatusCode)
			for key, values := range res.Header {
				for _, value := range values {
					c.Set(key, value)
				}
			}
			for _, cookie := range cookies {
				c.Cookie(cookie)
			}
			_, err := io.Copy(c, res.Body)
			if err != nil {
				fmt.Printf("Failed to copy response body: %v\n", err)
				return err
			}
			return nil
		}
	}
}

func APIHandler(app *fiber.App) {
	app.All("/*", func(c *fiber.Ctx) error {
		path := c.Path()
		url := "https://poe.com" + path
		return forwardRequest(c, url)
	})
}
