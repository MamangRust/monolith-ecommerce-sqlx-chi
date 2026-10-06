# Migrasi Apigateway: Echo → Chi

> **Lingkup:** `monolith-ecommerce-grpc/service/apigateway` (+ harness `tests/` dan `pkg/upload_image`)
> **Router:** `github.com/labstack/echo/v4` → `github.com/go-chi/chi/v5`
> **Status:** ✅ Selesai — build ✅ · vet ✅ · unit test ✅
> **Pola:** sama dengan migrasi apigateway `monolith-payment-gateway-grpc` (lihat `MIGRATION_ECHO_TO_CHI.md` di sana).

---

## 1. Ringkasan

| Item | Jumlah |
|---|---|
| File yang import echo (sebelum) | 76 |
| Paket handler domain yang dimigrasi | 21 (`auth`, `banner`, `cart`, `category`, `merchant`, `merchant_award`, `merchant_business`, `merchant_detail`, `merchant_document`, `merchant_policy`, `merchant_social_link`, `order`, `order_item`, `product`, `review`, `review_detail`, `role`, `shipping_address`, `slider`, `transaction`, `user`) |
| Rute terdaftar (terkonversi) | 251 (120 GET · 111 POST · 20 DELETE) |
| Middleware global | 11 (lebih banyak dari payment-gateway) |
| Test harness `tests/` yang dikonversi | ±39 file |
| Rute terdampak duplikasi mount | 1 (`/api/order-item`) — dimerge |

## 2. Perbedaan dari Payment-Gateway (yang perlu disesuaikan)

Pola dasar identik, tapi ada 6 varian yang ditangani:

1. **Error handler global yang lebih pintar** — ecommerce punya `RegisterErrorHandler` yang memetakan `*AppError` shared → status code asli (bukan 500) dengan `trace_id` dari header `X-Request-ID`, dan fallback ke format default echo (`{"message": "..."}`). Di chi, logika ini pindah ke adapter `httpx.Handler` (untuk rute unwrapped) dan `apierror.HandleApiError` (untuk rute wrapped, plus dukungan `*echo.HTTPError` → `httpx.HTTPError`). `middlewares/errorhandler.go` dihapus.
2. **Middleware ekstra** — `trace.go` (OTel span per request) dan `metrics.go` (HTTP metrics) di-rewrite ke chi dengan `wrapResponseWriter` untuk menangkap status code (pengganti `c.Response().Status`); `rate limit` khusus prefix `/api/auth/` tetap sebagai middleware kondisional.
3. **Multipart / file upload** — handler `product`, `category`, `slider`, `merchant_document`, `review_detail`, `transaction` memakai `c.FormValue`/`c.FormFile`. Konversi: `r.FormValue`/`r.FormFile` (perhatikan `r.FormFile` mengembalikan **3 nilai** vs 2 di echo). `pkg/upload_image.ProcessImageUpload` diubah dari `echo.Context` → `http.ResponseWriter` (helper `writeJSONError` internal) — satu-satunya konsumennya adalah apigateway.
4. **JWT menyimpan user_id sebagai `int`** — `JWTAuth()` custom mengekstrak claim `sub` via `extractUserIDFromClaims` dan menyimpan `user_id` int ke context (echo-jwt lama menyimpan via `SuccessHandler`).
5. **Role validator berbasis gRPC** — `RoleValidatorGRPC` (bukan Kafka) di-rewrite ke chi middleware.
6. **Duplikasi mount `/api/order-item`** — echo mengizinkan dua `Group` dengan prefix sama; chi panik saat `Route()` duplikat. Rute query kini registrasi langsung di parent router dengan path penuh, command tetap pakai `Route()`.

## 3. Dependensi

**Ditambah:** `go-chi/chi/v5 v5.3.2`, `go-chi/cors v1.2.2`, `swaggo/http-swagger v1.3.4`.
**Tidak lagi di-import langsung:** `labstack/echo/v4`, `echo-jwt/v4`, `echo-swagger` (masih *indirect* via `shared/errors` yang dipakai modul `tests/`).

## 4. Struktur Baru

```
service/apigateway/
├── apierror/          # NEW — ApiHandler versi net/http (AppError + HTTPError aware)
├── httpx/             # NEW — JSON/Bind/RealIP, request-scoped values, error adapter
├── apps/client.go     # bootstrap chi (Trace → Metrics → Recover → ... → JWT)
├── handler/           # 43 file → chi.Router
├── middlewares/       # auth, role(+gRPC), requiredRole, ratelimit, metrics, trace, pyroscope
└── ...
pkg/upload_image/      # ProcessImageUpload(w http.ResponseWriter, ...)
```

## 5. Pemetaan Konversi (inti)

Sama dengan payment-gateway; ringkas:

| Echo | Chi / net/http |
|---|---|
| `func(c echo.Context) error` | `func(w http.ResponseWriter, r *http.Request) error` |
| `c.FormValue` / `c.FormFile` (2 nilai) | `r.FormValue` / `r.FormFile` (3 nilai: file, header, err) |
| `RegisterErrorHandler(e)` global | adapter `httpx.Handler` + `apierror.HandleApiError` |
| `echo.NewHTTPError` → handler default `{"message":...}` | `httpx.NewHTTPError` + `httpx.WriteHTTPError` (format sama) |
| `c.Response().Status` (metrics/trace) | `wrapResponseWriter.Status()` |
| `e.GET/POST/DELETE` , `:param`, `Group("")` | `r.Get/Post/Delete`, `{param}`, `Route("")` + `.Get("/")` |
| Chain `m1(m2(h.X))` | `r.With(m1, m2).Get("/", ...)` |

Urutan middleware global dipertahankan: `Trace → HTTPMetrics → Recoverer → RequestID → Logger(JSON) → Pyroscope → CORS → Compress → Secure → RateLimit(/api/auth/) → JWTAuth`.

## 6. Verifikasi

```bash
cd monolith-ecommerce-grpc/service/apigateway
go build ./...        # ✅
go vet ./...          # ✅
go test -count=1 ./...  # ✅ handler (gateway_test: inventory rute, swagger↔rute, smoke 401/404/400/503) + middlewares
```

- `handler/gateway_test.go` di-rewrite ke `chi.Walk` untuk inventory rute (chi `Routes()` tidak punya field method) + normalisasi trailing slash pada root subrouter.
- Harness `tests/` (39 file): `Deps{E: …}` → `{Router: …}`, bypass middleware `c.Set("user_id", …)` → `context.WithValue`, `MockImageUpload` ikut signature baru, `RegisterErrorHandler(s.echo)` (no-op reinstall) dihapus.
- `shared` dan seluruh service gRPC tidak tersentuh; `pkg/upload_image` diubah karena satu-satunya konsumen adalah apigateway.

## 7. Catatan Perilaku

1. **Respons error tidak berubah:** `*AppError` → status asli + shape `ErrorResponse` (trace_id dari X-Request-ID untuk rute unwrapped, dari span OTel untuk rute wrapped); middleware rejection & 404/405 tetap format compact `{"message": "..."}`; 400 dari body/param tidak valid tetap 400.
2. **Smoke test 401/404/400/503 lulus** — matriks status pada middleware chain asli terjaga.
3. **Echo masih indirect dependency** via `shared/errors` (echo-typed, dipakai `tests/`).

## 8. Follow-up (Opsional)

- [ ] Smoke test end-to-end dengan skenario `hurl/` atau full stack docker.
- [ ] `swag init` regenerasi docs bila perlu (anotasi `@Router` tidak berubah).
- [ ] Migrasi `tests/` + `shared/errors` penuh ke net/http → echo hilang total.
- [ ] Titik yang sama untuk apigateway `monolith-pointofsale-grpc` (16 file echo) — pola di dokumen ini + payment-gateway mencakup semua variannya.
