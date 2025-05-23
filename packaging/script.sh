#!/bin/bash

set -e

# ==========================
# 🧩 Tham số & cấu hình
# ==========================

if [ -z "$1" ]; then
  echo "❌ Thiếu tên tổ chức. Cách dùng: ./package_chaincode.sh org1"
  exit 1
fi

ORG="$1"
CHAINCODE_NAME="token"
CHAINCODE_LABEL="${CHAINCODE_NAME}"
ADDRESS="${CHAINCODE_NAME}-service-${ORG}:7052"


echo "🔧 Đang đóng gói chaincode cho '${ORG}' với address '${ADDRESS}'..."

# ==========================
# 📁 Chuẩn bị thư mục
# ==========================

rm -rf code
mkdir -p code

# ==========================
# 📝 Ghi file connection.json
# ==========================

cat > code/connection.json <<EOF
{
  "address": "${ADDRESS}",
  "dial_timeout": "10s",
  "tls_required": false,
  "client_auth_required": false,
  "client_key": "-----BEGIN EC PRIVATE KEY----- ... -----END EC PRIVATE KEY-----",
  "client_cert": "-----BEGIN CERTIFICATE----- ... -----END CERTIFICATE-----",
  "root_cert": "-----BEGIN CERTIFICATE---- ... -----END CERTIFICATE-----"
}
EOF

# ==========================
# 📝 Ghi file metadata.json
# ==========================

cp metadata.json ./code/metadata.json

OUTPUT_PACKAGE="${CHAINCODE_LABEL}-${ORG}.tgz"

# # ==========================
# # 📦 Đóng gói .tar.gz đúng chuẩn
# # ==========================
cd ./code
# Tạo file nén chứa connection.json
tar cfz code.tar.gz connection.json

tar cfz ${OUTPUT_PACKAGE} code.tar.gz metadata.json


cp ${CHAINCODE_NAME}-${ORG}.tgz ../../

rm -rf code.tar.gz
cd ..

# rm -rf code
# ==========================
# 🧩 Đóng gói chaincode

echo "✅ Gói chaincode đã tạo: ${OUTPUT_PACKAGE}"



