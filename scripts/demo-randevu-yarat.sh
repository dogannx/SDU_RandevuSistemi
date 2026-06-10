#!/usr/bin/env bash
# Lokal backend'e demo amaçlı randevu yaratır — RabbitMQ'ya event publish tetiklemek için.
# Kullanım: bash scripts/demo-randevu-yarat.sh
set -e

API="http://localhost:8080/api/v1"
EMAIL="demo@local.com"
PASSWORD="demo1234"

echo "→ Login..."
TOKEN=$(curl -sS -X POST "$API/auth/login" \
  -H "Content-Type: application/json" \
  -d "{\"email\":\"$EMAIL\",\"password\":\"$PASSWORD\"}" \
  | python3 -c "import sys,json;print(json.load(sys.stdin)['data']['token'])")

echo "→ Öğretmen ID alınıyor..."
TEACHER_ID=$(curl -sS "$API/teachers" \
  | python3 -c "import sys,json;print(json.load(sys.stdin)['data'][0]['id'])")

DATE="2026-06-$((10 + RANDOM % 20))"
TIME="$((10 + RANDOM % 8)):00"

echo "→ Randevu yaratılıyor: $DATE $TIME"
curl -sS -X POST "$API/appointments" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d "{\"teacherId\":\"$TEACHER_ID\",\"date\":\"$DATE\",\"time\":\"$TIME\"}"
echo
echo "✓ Tamam. RabbitMQ UI'ye bak — appointment.created kuyruğunda anlık mesaj görünmeli."
