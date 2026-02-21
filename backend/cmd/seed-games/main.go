// seed-games uploads game cover images from the backend/static/games directory
// to the MinIO public bucket (gm-public) and ensures the bucket has the correct
// public-read policy.
//
// Run from the backend directory:
//
//	go run ./cmd/seed-games/
//
// Override the games directory path:
//
//	go run ./cmd/seed-games/ --games-dir /custom/path/to/games
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/adan-ayaz-stan/gamebuddies/backend/internal/storage"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	// Load .env (best-effort; env vars take precedence if .env is absent)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, falling back to environment variables")
	}

	// Default path: run script from the backend/ directory, images live in
	// the local static/games folder.
	defaultGamesDir := filepath.Join("static", "games")
	gamesDir := flag.String("games-dir", defaultGamesDir, "Path to the directory containing game cover images")
	flag.Parse()

	// MinIO connection settings
	endpoint := os.Getenv("MINIO_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:9000"
	}
	accessKey := os.Getenv("MINIO_ACCESS_KEY")
	secretKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := os.Getenv("MINIO_USE_SSL") == "true"

	mc, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalf("Failed to create MinIO client: %v", err)
	}

	ctx := context.Background()

	// ── Ensure the public games bucket exists ────────────────────────────────
	exists, err := mc.BucketExists(ctx, storage.GamesBucket)
	if err != nil {
		log.Fatalf("Failed to check whether bucket %q exists: %v", storage.GamesBucket, err)
	}
	if !exists {
		if err := mc.MakeBucket(ctx, storage.GamesBucket, minio.MakeBucketOptions{}); err != nil {
			log.Fatalf("Failed to create bucket %q: %v", storage.GamesBucket, err)
		}
		log.Printf("Bucket %q created.", storage.GamesBucket)
	} else {
		log.Printf("Bucket %q already exists.", storage.GamesBucket)
	}

	// Apply (or re-apply) public-read policy so images are served without auth
	policy := fmt.Sprintf(
		`{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"AWS":["*"]},"Action":["s3:GetObject"],"Resource":["arn:aws:s3:::%s/*"]}]}`,
		storage.GamesBucket,
	)
	if err := mc.SetBucketPolicy(ctx, storage.GamesBucket, policy); err != nil {
		log.Fatalf("Failed to set public-read policy on bucket %q: %v", storage.GamesBucket, err)
	}
	log.Printf("Public-read policy applied to bucket %q.", storage.GamesBucket)

	// ── Upload images ────────────────────────────────────────────────────────
	absDir, err := filepath.Abs(*gamesDir)
	if err != nil {
		log.Fatalf("Failed to resolve games directory path: %v", err)
	}
	log.Printf("Reading game images from: %s", absDir)

	entries, err := os.ReadDir(absDir)
	if err != nil {
		log.Fatalf("Failed to read games directory %q: %v", absDir, err)
	}

	scheme := "http"
	if useSSL {
		scheme = "https"
	}

	uploaded := 0
	skipped := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		ext := strings.ToLower(filepath.Ext(name))

		// Only process recognised image formats
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" && ext != ".gif" && ext != ".avif" {
			log.Printf("Skipping non-image file: %s", name)
			skipped++
			continue
		}

		contentType := mime.TypeByExtension(ext)
		if contentType == "" {
			contentType = "application/octet-stream"
		}

		imgPath := filepath.Join(absDir, name)
		f, err := os.Open(imgPath)
		if err != nil {
			log.Printf("WARNING: Cannot open %q — skipping: %v", imgPath, err)
			skipped++
			continue
		}

		info, err := f.Stat()
		if err != nil {
			f.Close()
			log.Printf("WARNING: Cannot stat %q — skipping: %v", imgPath, err)
			skipped++
			continue
		}

		_, err = mc.PutObject(ctx, storage.GamesBucket, name, f, info.Size(), minio.PutObjectOptions{
			ContentType: contentType,
		})
		f.Close()

		if err != nil {
			log.Printf("WARNING: Failed to upload %q: %v", name, err)
			skipped++
			continue
		}

		publicURL := fmt.Sprintf("%s://%s/%s/%s", scheme, endpoint, storage.GamesBucket, name)
		log.Printf("  ✓  %-30s  →  %s", name, publicURL)
		uploaded++
	}

	fmt.Println()
	log.Printf("Done. %d image(s) uploaded, %d skipped.", uploaded, skipped)
	log.Printf("Public URL pattern: %s://%s/%s/<filename>", scheme, endpoint, storage.GamesBucket)
}
