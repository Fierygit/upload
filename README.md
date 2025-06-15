
# Public File Upload Service

A lightweight, high-performance file hosting service that provides instant public URLs for uploaded files via simple `curl` commands.

## Features
- Zero-configuration file hosting
- HTTP PUT-based uploads
- Instant public URL generation

## Usage
```bash
# Upload any file type
curl upload.firego.cn -T [your-file]

# Examples:
curl xxx.firego.cn -T document.pdf
curl xxx.firego.cn -T image.jpg
curl xxx.firego.cn -T data.bin


Response

Successful uploads return a public URL in plain text format:
https://xxx.firego.cn/[random-filename].[ext]
```
